package dao

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type UUIDMapper struct{}

func (m UUIDMapper) ToIdentifiers(rows *sql.Rows) ([]Identifier, error) {
	identifiers := make([]Identifier, 0)
	for rows.Next() {
		var id uuid.UUID
		err := rows.Scan(&id)
		if err != nil {
			return nil, errors.Wrapf(err, "scan id into uuid")
		}
		identifiers = append(identifiers, id)
	}
	return identifiers, nil
}
