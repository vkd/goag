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
	Type        SchemaType
}

func NewQueryParameter(refP specification.Ref[specification.QueryParameter], components Components) (zero *QueryParameter, _ Imports, _ error) {
	s := refP.Value()
	out := QueryParameter{
		Description: s.Description,
	}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	out.Required = s.Required
	var err error
	var ims Imports
	out.Type, ims, err = NewSchema(s.Schema, components)
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
	Type          SchemaType
	Description   string
}

func NewPathParameter(rs specification.Ref[specification.PathParameter], componenets Components) (zero *PathParameter, _ Imports, _ error) {
	s := rs.Value()
	out := PathParameter{Description: s.Description}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)

	var ims Imports
	var err error
	out.Type, ims, err = NewSchema(s.Schema, componenets)
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
	Type          SchemaType
	Required      bool
}

func NewHeaderParameter(sr specification.Ref[specification.HeaderParameter], components Components) (zero *HeaderParameter, _ Imports, _ error) {
	s := sr.Value()
	out := HeaderParameter{
		Description: s.Description,
	}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)

	var ims Imports
	var err error
	out.Type, ims, err = NewSchema(s.Schema, components)
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	out.Required = s.Required
	return &out, ims, nil
}

type CookieParameter struct {
	specification.CookieParameter
}
