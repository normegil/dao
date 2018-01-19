package dao

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/gchaincl/dotsql"
	"github.com/pkg/errors"
)

type queryKey string

const (
	getAllEntities = queryKey("getAllEntities")
	getAllIDs      = queryKey("getAllIds")
	size           = queryKey("size")
	get            = queryKey("get")
	insert         = queryKey("insert")
	update         = queryKey("update")
	remove         = queryKey("delete")
)

type ParametrizedQueries struct {
	parameters map[string]string
	queries    map[string]string
}

//go:generate go-bindata -pkg $GOPACKAGE -o sql.generated.go sql/...

func NewGenericQueries(table string, columns []string) (*ParametrizedQueries, error) {
	values := make([]string, 0)
	sets := make([]string, 0)
	for index, columnName := range columns {
		indexStr := strconv.Itoa(index + 1)
		values = append(values, "$"+indexStr)
		sets = append(sets, columnName+"=$"+indexStr)
	}
	return NewParametrizedQueries(map[string]string{
		"table":   table,
		"columns": strings.Join(columns, ","),
		"values":  strings.Join(values, ","),
		"sets":    "SET " + strings.Join(sets, ","),
	})
}

func NewParametrizedQueries(parameters map[string]string) (*ParametrizedQueries, error) {
	const genericSQLQueriesPath = "sql/generic.sql"
	genericSQLBytes, err := Asset(genericSQLQueriesPath)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not load %s", genericSQLQueriesPath)
	}
	dot, err := dotsql.Load(bytes.NewReader(genericSQLBytes))
	if err != nil {
		return nil, errors.Wrapf(err, "Could not parse %s", genericSQLQueriesPath)
	}
	return &ParametrizedQueries{
		parameters: parameters,
		queries:    dot.QueryMap(),
	}, nil
}

func (p ParametrizedQueries) GetAllEntities() string {
	return p.replaceWithParameters(p.queries["GetAll"], p.parameters)
}
func (p ParametrizedQueries) GetAllIDs() string {
	return p.replaceWithParameters(p.queries["GetIDs"], p.parameters)
}
func (p ParametrizedQueries) TotalNumberOfEntities() string {
	return p.replaceWithParameters(p.queries["TotalNumberOfEntities"], p.parameters)
}
func (p ParametrizedQueries) Get() string {
	return p.replaceWithParameters(p.queries["GetByID"], p.parameters)
}
func (p ParametrizedQueries) Insert() string {
	return p.replaceWithParameters(p.queries["Insert"], p.parameters)
}
func (p ParametrizedQueries) Update() string {
	return p.replaceWithParameters(p.queries["Update"], p.parameters)
}
func (p ParametrizedQueries) Delete() string {
	return p.replaceWithParameters(p.queries["Delete"], p.parameters)
}
func (p ParametrizedQueries) replaceWithParameters(query string, parameters map[string]string) string {
	finalQuery := query
	for toReplace, newValue := range p.parameters {
		finalQuery = strings.Replace(finalQuery, "%{"+toReplace+"}", newValue, -1)
	}
	return finalQuery
}
