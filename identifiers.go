package dao

import "github.com/satori/go.uuid"

type UUIDIdentifierGenerator struct {
}

func (g UUIDIdentifierGenerator) Generate(_ Entity) (Identifier, error) {
	return uuid.NewV4(), nil
}

