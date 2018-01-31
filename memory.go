package dao

import (
	"github.com/pkg/errors"
	"fmt"
)

type Memory struct {
	store []Entity
	generator IdentifierGenerator
}

func NewMemory(generator IdentifierGenerator) *Memory {
	if nil == generator {
		generator = UUIDIdentifierGenerator{}
	}
	return &Memory{
		generator: generator,
		store: make([]Entity, 0),
	}
}

func (d Memory) GetAllEntities(pagination Pagination) ([]Entity, error) {
	start := pagination.offset
	end := pagination.offset + pagination.limit
	if int64(len(d.store)) < end {
		end = int64(len(d.store))
	}
	return d.store[start:end], nil
}
func (d Memory) GetAllIDs(pagination Pagination) ([]Identifier, error) {
	entities, err := d.GetAllEntities(pagination)
	if err != nil {
		return nil, errors.Wrapf(err, "loading all identifiers")
	}

	var identifiers []Identifier
	for _, entity := range entities {
		iEntity, ok := entity.(IdentifiableEntity)
		if !ok {
			return nil, fmt.Errorf("GetIDs is only supported by IdentifiableEntity")
		}
		identifiers = append(identifiers, iEntity.ID())
	}
	return identifiers, nil
}

func (d Memory) TotalNumberOfEntities() (int64, error)             {
	return int64(len(d.store)), nil
}

func (d Memory) Get(id Identifier) (Entity, error)                 {
	_, entity, err := d.search(id)
	return entity, err
}

func (d *Memory) Set(entity IdentifiableEntity) (Identifier, error) {
	var index int
		var foundEntity Entity
	if nil == entity.ID() {
		identifier, err := d.generator.Generate(entity)
		if err != nil {
			return nil, errors.Wrapf(err, "generating id for %+v", entity)
		}
		newEntity, err := entity.WithID(identifier)
		if err != nil {
			return nil, errors.Wrapf(err, "replacing id of '%+v' with '%+v'", entity, identifier)
		}
		entity = newEntity
	} else {
		var err error
		index, foundEntity, err = d.search(entity.ID())
		if err != nil {
			return nil, errors.Wrapf(err, "searching for entity with id '%s'", entity.ID())
		}
	}

	if nil != foundEntity {
		d.store[index] = entity
	} else {
		d.store = append(d.store, entity)
	}
	return entity.ID(), nil
}

func (d Memory) search(id Identifier) (int, Entity, error) {
	for i, entity := range d.store {
		iEntity, ok := entity.(IdentifiableEntity)
		if !ok {
			return 0, nil, fmt.Errorf("get by ID is only supported by IdentifiableEntity")
		}
		if iEntity.ID() == id {
			return i, entity, nil
		}
	}
	return 0, nil, nil
}

func (d *Memory) Delete(id Identifier) error                        {
	index, _, err := d.search(id)
	if err != nil {
		return errors.Wrapf(err, "deleting '%+v'", id)
	}

	if index+1 < len(d.store) {
		d.store = append(d.store[0:index], d.store[index+1:]...)
	} else {
		d.store = d.store[0:index]
	}

	return nil
}