package dao_test

import (
	"testing"
	"github.com/normegil/dao"
	"strconv"
	"fmt"
)

func TestComplyToInterface(t *testing.T) {
	var _ dao.DAO = &dao.Memory{}
}

func TestDelete(t *testing.T) {
	testcases := []struct{
		Name string
		Content []dao.IdentifiableEntity
		IDToDelete dao.Identifier
	}{
		{
			Name: "Single element",
			Content: []dao.IdentifiableEntity{
				TestEntity{
					Identifier: TestIdentifier(1),
				},
			},
			IDToDelete: TestIdentifier(1),
		},
		{
			Name: "First element",
			Content: []dao.IdentifiableEntity{
				TestEntity{
					Identifier: TestIdentifier(1),
				},
				TestEntity{
					Identifier: TestIdentifier(2),
				},
				TestEntity{
					Identifier: TestIdentifier(3),
				},
			},
			IDToDelete: TestIdentifier(1),
		},
		{
			Name: "Last element",
			Content: []dao.IdentifiableEntity{
				TestEntity{
					Identifier: TestIdentifier(1),
				},
				TestEntity{
					Identifier: TestIdentifier(2),
				},
				TestEntity{
					Identifier: TestIdentifier(3),
				},
			},
			IDToDelete: TestIdentifier(3),
		},
	}
	for _, testdata := range testcases {
		t.Run(testdata.Name, func(t *testing.T) {
			memoryDAO := dao.NewMemory(nil)
			for _, entity := range testdata.Content {
				if _, err := memoryDAO.Set(entity); nil != err {
					t.Fatalf("saving %+v: %s", testdata.Content, entity)
				}
			}

			if err := memoryDAO.Delete(testdata.IDToDelete); nil != err {
				t.Fatalf("deleting '%s': %s", testdata.IDToDelete, err.Error())
			}

			for _, content := range testdata.Content {
				entity, err := memoryDAO.Get(content.ID())
				if err != nil {
					t.Fatalf("loading '%s': %s", testdata.IDToDelete, err.Error())
				}

				if nil != entity && content.ID() == testdata.IDToDelete {
					t.Errorf("entity '%s' was not removed", testdata.IDToDelete)
				} else if nil == entity && content.ID() != testdata.IDToDelete{
					t.Errorf("entity '%s' removed when it shouldn't (ID '%s' should be removed)", content.ID(), testdata.IDToDelete)
				}
			}
		})
	}
}

type TestEntity struct {
	Identifier TestIdentifier
}

func (t TestEntity) ID() dao.Identifier {
	return t.Identifier
}

func (t TestEntity) WithID(id dao.Identifier) (dao.IdentifiableEntity, error) {
	identifier, ok := id.(TestIdentifier)
	if !ok {
		return nil, fmt.Errorf("identifier shoud be a TestIdentifier")
	}
	return TestEntity{
		Identifier: TestIdentifier(identifier),
	}, nil
}

type TestIdentifier int

func (ti TestIdentifier) String() string {
	return strconv.Itoa(int(ti))
}