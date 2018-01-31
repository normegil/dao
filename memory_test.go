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

func TestGetAllEntities(t *testing.T) {
	testcases := []struct{
		Name string
		Pagination dao.Pagination
		Content []dao.IdentifiableEntity
		Expected []dao.IdentifiableEntity
	}{
		{
			Name:"Empty DAO",
			Pagination: dao.Pagination{},
			Content: []dao.IdentifiableEntity{},
			Expected: []dao.IdentifiableEntity{},
		},
		{
			Name:"Default Pagination",
			Pagination: dao.Pagination{},
			Content: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(1)},
				TestEntity{TestIdentifier(2)},
				TestEntity{TestIdentifier(3)},
			},
			Expected: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(1)},
				TestEntity{TestIdentifier(2)},
				TestEntity{TestIdentifier(3)},
			},
		},
		{
			Name:"With Offset",
			Pagination: dao.Pagination{}.WithOffset(1),
			Content: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(1)},
				TestEntity{TestIdentifier(2)},
				TestEntity{TestIdentifier(3)},
			},
			Expected: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(2)},
				TestEntity{TestIdentifier(3)},
			},
		},
		{
			Name:"With Limit",
			Pagination: dao.Pagination{}.WithLimit(2),
			Content: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(1)},
				TestEntity{TestIdentifier(2)},
				TestEntity{TestIdentifier(3)},
			},
			Expected: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(1)},
				TestEntity{TestIdentifier(2)},
			},
		},
		{
			Name:"With Offset & Limit",
			Pagination: dao.Pagination{}.WithOffset(1).WithLimit(1),
			Content: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(1)},
				TestEntity{TestIdentifier(2)},
				TestEntity{TestIdentifier(3)},
			},
			Expected: []dao.IdentifiableEntity{
				TestEntity{TestIdentifier(2)},
			},
		},
	}
	for _, testdata := range testcases {
		t.Run(testdata.Name, func(t *testing.T) {
			memoryDAO := dao.NewMemory(nil)
			for _, entity := range testdata.Content {
				if _, err := memoryDAO.Set(entity); nil != err {
					t.Fatalf("inserting %+v: %s", entity, err.Error())
				}
			}

			entities, err := memoryDAO.GetAllEntities(testdata.Pagination)
			if err != nil {
				t.Fatalf("loading entities: %s", err.Error())
			}

			if len(testdata.Expected) != len(entities) {
				t.Fatalf("Expected (%d) & Loaded (%d) entities doesn't have the same number of entities", len(testdata.Expected), len(entities))
			}

			for i, expected := range testdata.Expected {
				if expected != entities[i] {
					t.Errorf("Expected (%+v) & Loaded (%+v) entities doesn't match (Index: %d)", expected, entities[i], i)
				}
			}
		})
	}
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