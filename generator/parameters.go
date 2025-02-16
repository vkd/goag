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

func NewQueryParameter(refP specification.Ref[specification.QueryParameter], components Componenter, cfg Config) (zero *QueryParameter, _ Imports, _ error) {
	s := refP.Value()
	out := QueryParameter{
		Description: s.Description,
	}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	out.Required = s.Required
	schema, ims, err := NewSchema(s.Schema, components, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	var st SchemaType = schema
	if !s.Required {
		st = NewOptionalType(schema, cfg)
	}
	out.Type = st
	return &out, ims, nil
}

type PathParameters []*PathParameter

func (s PathParameters) Get(name string) (zero *PathParameter, _ error) {
	for _, p := range s {
		if p.Name == name {
			return p, nil
		}
	}
	return zero, fmt.Errorf("path parameter %q: not found", name)
}

type PathParameter struct {
	Name        string
	FieldName   string
	Type        Schema
	Description string
}

func NewPathParameter(rs specification.Ref[specification.PathParameter], components Componenter, cfg Config) (zero *PathParameter, _ Imports, _ error) {
	s := rs.Value()

	tp, ims, err := NewSchema(s.Schema, components, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("schema: %w", err)
	}
	if tp.Kind() != SchemaKindPrimitive {
		return nil, nil, fmt.Errorf("path parameter could only be a primitive type, found %q", tp.Kind())
	}
	return &PathParameter{
		Name:        s.Name,
		FieldName:   PublicFieldName(s.Name),
		Type:        tp,
		Description: s.Description,
	}, ims, nil
}

func (p PathParameter) RenderToStringInline(from string) (string, error) {
	return p.Type.RenderFormat(from + "." + p.FieldName)
}

type HeaderParameter struct {
	Name          string
	Description   string
	FieldName     string
	FieldTypeName string
	Type          SchemaType
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
	var st SchemaType = schema
	if !s.Required {
		st = NewOptionalType(st, cfg)
	}
	out.Type = st
	out.Required = s.Required
	return &out, ims, nil
}

type CookieParameter struct {
	specification.CookieParameter
}
