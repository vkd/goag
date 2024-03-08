package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Components struct {
	Imports Imports

	Schemas []SchemaComponent
}

func NewComponents(spec specification.Components) (zero Components, _ error) {
	var cs Components

	cs.Schemas = make([]SchemaComponent, 0, len(spec.Schemas.List))
	for _, c := range spec.Schemas.List {
		s, ims, err := NewSchema(c.V.Value())
		if err != nil {
			return zero, fmt.Errorf("parse schema for %q type: %w", c.Name, err)
		}
		cs.Schemas = append(cs.Schemas, SchemaComponent{Name: c.Name, Type: s})
		cs.Imports = append(cs.Imports, ims...)
	}
	return cs, nil
}

func (c Components) Render() (string, error) {
	return ExecuteTemplate("Components", c)
}

type SchemaComponent struct {
	Name string
	Type Render
}

func (s SchemaComponent) Render() (string, error) {
	return ExecuteTemplate("SchemaComponent", s)
}
