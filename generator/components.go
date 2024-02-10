package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Components struct {
	Schemas []SchemaComponent
}

func NewComponents(spec specification.Components) (zero Components, _ error) {
	var cs Components

	cs.Schemas = make([]SchemaComponent, 0, len(spec.Schemas))
	for _, c := range spec.Schemas {
		s, err := NewSchema(c.Schema)
		if err != nil {
			return zero, fmt.Errorf("parse schema for %q type: %w", c.Name, err)
		}
		cs.Schemas = append(cs.Schemas, SchemaComponent{Name: c.Name, Type: s})
	}
	return cs, nil
}

type SchemaComponent struct {
	Name string
	Type Render
}
