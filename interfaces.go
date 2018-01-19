package dao

import (
	"database/sql"
	"fmt"
)

type Identifier interface {
	fmt.Stringer
}

type IdentifierGenerator interface {
	Generate(entity Entity) (Identifier, error)
}

type Entity interface{}

type IdentifiableEntity interface {
	Entity
	ID() Identifier
	WithID(Identifier) (IdentifiableEntity, error)
}

type DAO interface {
	GetAllEntities(Pagination) ([]Entity, error)
	GetAllIDs(Pagination) ([]Identifier, error)
	TotalNumberOfEntities() (int64, error)
	Get(Identifier) (Entity, error)
	Set(entity IdentifiableEntity) (Identifier, error)
	Delete(Identifier) error
}

type Mapper interface {
	ToEntities(*sql.Rows) ([]Entity, error)
	ToIdentifiers(*sql.Rows) ([]Identifier, error)
	ToSlice(Entity) ([]interface{}, error)
}

type Queries interface {
	GetAllEntities() string
	GetAllIDs() string
	TotalNumberOfEntities() string
	Get() string
	Insert() string
	Update() string
	Delete() string
}
