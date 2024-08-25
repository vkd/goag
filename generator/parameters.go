package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type QueryParameter struct {
	Name        string
	Description string
	FieldName   string
	Required    bool
	Type        Schema
}

func NewQueryParameter(refP specification.Ref[specification.QueryParameter], components Componenter, cfg Config) (zero *QueryParameter, _ Imports, _ error) {
	s := refP.Value()
	out := QueryParameter{
		Description: s.Description,
	}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	out.Required = s.Required
	var err error
	var ims Imports
	out.Type, ims, err = NewSchema(s.Schema, components, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	return &out, ims, nil
}

type PathParameters []*PathParameter

func (s PathParameters) Get(name string) (zero *PathParameter, _ error) {
	for _, p := range s {
		if p.s.Name == name {
			return p, nil
		}
	}
	return zero, fmt.Errorf("path parameter %q: not found", name)
}

type PathParameter struct {
	s *specification.PathParameter

	Name          string
	FieldName     string
	FieldTypeName string
	Type          Schema
	Description   string
}

func NewPathParameter(rs specification.Ref[specification.PathParameter], components Componenter, cfg Config) (zero *PathParameter, _ Imports, _ error) {
	s := rs.Value()
	out := PathParameter{Description: s.Description}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)

	var ims Imports
	var err error
	out.Type, ims, err = NewSchema(s.Schema, components, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("schema: %w", err)
	}
	return &out, ims, nil
}

type HeaderParameter struct {
	Name          string
	Description   string
	FieldName     string
	FieldTypeName string
	Type          Schema
	Schema        Schema
	Required      bool
}

func NewHeaderParameter(sr specification.Ref[specification.HeaderParameter], components Componenter, cfg Config) (zero *HeaderParameter, _ Imports, _ error) {
	s := sr.Value()
	out := HeaderParameter{
		Description: s.Description,
	}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)

	schema, ims, err := NewSchema(s.Schema, components, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	out.Type = schema
	out.Schema = schema
	out.Required = s.Required
	return &out, ims, nil
}

type CookieParameter struct {
	specification.CookieParameter
}
