package specification

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

type Components struct {
	Schemas []ComponentSchema
}

func NewComponents(spec openapi3.Components) Components {
	var cs Components

	cs.Schemas = make([]ComponentSchema, 0, len(spec.Schemas))
	for name, c := range spec.Schemas {
		cs.Schemas = append(cs.Schemas, ComponentSchema{Name: name, Schema: NewSchema(c)})
	}
	sort.Slice(cs.Schemas, func(i, j int) bool { return cs.Schemas[i].Name < cs.Schemas[j].Name })

	return cs
}

type ComponentSchema struct {
	Name string
	Schema
}
