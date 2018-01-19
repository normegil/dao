package dao

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type Database struct {
	idGenerator IdentifierGenerator
	mapper      Mapper
	queries     map[queryKey]*sql.Stmt
}

func newDatabaseFromBuilder(b *DatabaseBuilder) (*Database, error) {
	db := b.db
	queries := b.queries

	var err error
	preparedQueries := make(map[queryKey]*sql.Stmt)
	preparedQueries[getAllEntities], err = db.Prepare(queries.GetAllEntities())
	if err != nil {
		return nil, errors.Wrapf(err, "preparing %s", queries.GetAllEntities())
	}
	preparedQueries[getAllIDs], err = db.Prepare(queries.GetAllIDs())
	if err != nil {
		return nil, errors.Wrapf(err, "preparing %s", queries.GetAllIDs())
	}
	preparedQueries[size], err = db.Prepare(queries.TotalNumberOfEntities())
	if err != nil {
		return nil, errors.Wrapf(err, "preparing %s", queries.TotalNumberOfEntities())
	}
	preparedQueries[get], err = db.Prepare(queries.Get())
	if err != nil {
		return nil, errors.Wrapf(err, "preparing %s", queries.Get())
	}
	preparedQueries[insert], err = db.Prepare(queries.Insert())
	if err != nil {
		return nil, errors.Wrapf(err, "preparing %s", queries.Insert())
	}
	preparedQueries[update], err = db.Prepare(queries.Update())
	if err != nil {
		return nil, errors.Wrapf(err, "preparing %s", queries.Update())
	}
	preparedQueries[remove], err = db.Prepare(queries.Delete())
	if err != nil {
		return nil, errors.Wrapf(err, "preparing %s", queries.Delete())
	}

	return &Database{
		mapper:      b.mapper,
		queries:     preparedQueries,
		idGenerator: b.generator,
	}, nil
}

func (d *Database) Close() {
	for _, query := range d.queries {
		query.Close()
	}
}

func (d *Database) GetAllEntities(p Pagination) ([]Entity, error) {
	rows, err := d.queries[getAllEntities].Query(p.Offset(), p.Limit())
	if err != nil {
		return nil, errors.Wrapf(err, "retrieving entities from database")
	}
	defer rows.Close()
	entities, err := d.mapper.ToEntities(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "map result set to entity")
	}
	if err = rows.Err(); nil != err {
		return nil, errors.Wrapf(err, "looping through entity rows")
	}
	return entities, nil
}

func (d *Database) GetAllIDs(p Pagination) ([]Identifier, error) {
	queries := d.queries
	getAllQuery := queries[getAllIDs]
	rows, err := getAllQuery.Query(p.Offset(), p.Limit())
	if err != nil {
		return nil, errors.Wrapf(err, "retrieving entities from database")
	}
	defer rows.Close()
	identifiers, err := d.mapper.ToIdentifiers(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "map result set to identifiers")
	}
	if err = rows.Err(); nil != err {
		return nil, errors.Wrapf(err, "looping through identifiers rows")
	}
	return identifiers, nil
}

func (d *Database) TotalNumberOfEntities() (int64, error) {
	row := d.queries[size].QueryRow()
	var nbItems int64
	err := row.Scan(&nbItems)
	if err != nil {
		return 0, errors.Wrapf(err, "counting number of entities in database")
	}
	return nbItems, nil
}

func (d *Database) Get(id Identifier) (Entity, error) {
	rows, err := d.queries[get].Query(id)
	if err != nil {
		return nil, errors.Wrapf(err, "retrieving entities from database")
	}
	defer rows.Close()
	entities, err := d.mapper.ToEntities(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "map result set to entity")
	}
	if err = rows.Err(); nil != err {
		return nil, errors.Wrapf(err, "looping through entity rows")
	}
	nbEntities := len(entities)
	if nbEntities > 1 {
		return nil, fmt.Errorf("expected only one entity identified by '%s' but got %d", id, nbEntities)
	}
	if nbEntities == 0 {
		return nil, nil
	}
	return entities[0], nil
}

func (d *Database) Set(entity IdentifiableEntity) (Identifier, error) {
	id := entity.ID()
	shouldInsert := false
	if nil == id {
		generator := d.idGenerator
		id, err := generator.Generate(entity)
		if err != nil {
			return nil, errors.Wrapf(err, "generating Identifier")
		}
		entity, err = entity.WithID(id)
		if err != nil {
			return nil, errors.Wrapf(err, "setting ID")
		}
		shouldInsert = true
	} else {
		found, err := d.Get(id)
		if err != nil {
			return nil, errors.Wrapf(err, "checking if entity exist")
		}
		shouldInsert = nil == found
	}

	s, err := d.mapper.ToSlice(entity)
	if err != nil {
		return nil, errors.Wrapf(err, "turn an entity into a slice of fields")
	}
	if shouldInsert {
		_, err := d.queries[insert].Exec(s...)
		if err != nil {
			return nil, errors.Wrapf(err, "inserting '%+v'", entity)
		}
	} else {
		_, err := d.queries[update].Exec(s...)
		if err != nil {
			return nil, errors.Wrapf(err, "updating '%+v'", entity)
		}
	}

	return entity.ID(), nil
}

func (d *Database) Delete(id Identifier) error {
	_, err := d.queries[remove].Exec(id)
	if err != nil {
		return errors.Wrapf(err, "deleting '%s'", id.String())
	}
	return nil
}

type DatabaseBuilder struct {
	db        *sql.DB
	mapper    Mapper
	queries   Queries
	generator IdentifierGenerator
}

func NewDatabaseBuilder(db *sql.DB) *DatabaseBuilder {
	return &DatabaseBuilder{
		db:        db,
		generator: UUIDIdentifierGenerator{},
	}
}

func (b *DatabaseBuilder) Mapper(mapper Mapper) *DatabaseBuilder {
	b.mapper = mapper
	return b
}

func (b *DatabaseBuilder) Queries(queries Queries) *DatabaseBuilder {
	b.queries = queries
	return b
}

func (b *DatabaseBuilder) IdentifierGenerator(generator IdentifierGenerator) *DatabaseBuilder {
	b.generator = generator
	return b
}

func (b *DatabaseBuilder) build() (*Database, error) {
	return newDatabaseFromBuilder(b)
}
