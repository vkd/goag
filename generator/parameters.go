package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type QueryParameter struct {
	s *specification.QueryParameter

	Name      string
	FieldName string
	Required  bool
	Type      Formatter
}

func NewQueryParameter(s *specification.QueryParameter) (zero *QueryParameter, _ Imports, _ error) {
	out := QueryParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	out.Required = s.Required
	var err error
	var ims Imports
	out.Type, ims, err = NewParameterSchema(s.Schema.Value())
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	return &out, ims, nil
}

func (p QueryParameter) ExecuteFormat(to, from string) (string, error) {
	if sliceType, ok := p.Type.(SliceType); ok {
		switch sliceType.Items.(type) {
		case StringType:
			return Assign(to, from, false), nil
		default:
			return ExecuteTemplate("ClientQueryParameterFormatToSliceStrings", TData{
				"From":  from,
				"Items": sliceType.Items,
				"To":    to,
			})
		}
	}

	if !p.Required {
		pointer := "*"
		if _, ok := p.Type.(CustomType); ok {
			pointer = ""
		}
		return ExecuteTemplate("ClientQueryParameterFormatOptional", TData{
			"From":    from,
			"FromPtr": pointer + from,
			"To":      to,

			"Formatter": p.Type,
		})
	}

	return ExecuteTemplate("ClientQueryParameterFormatRequired", TData{
		"From":      from,
		"To":        to,
		"Formatter": p.Type,
	})
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
}

func NewPathParameter(rs specification.Ref[specification.PathParameter]) (zero *PathParameter, _ Imports, _ error) {
	s := rs.Value()
	out := PathParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if rs.Ref() != nil {
		out.FieldTypeName = rs.Ref().Name
	}
	var ims Imports
	var err error
	out.Type, ims, err = NewParameterSchema(s.Schema.Value())
	if err != nil {
		return nil, nil, fmt.Errorf("schema: %w", err)
	}
	return &out, ims, nil
}

type HeaderParameter struct {
	s *specification.HeaderParameter

	Name          string
	FieldName     string
	FieldTypeName string
	Type          Formatter
	Required      bool
}

func NewHeaderParameter(sr specification.Ref[specification.HeaderParameter]) (zero *HeaderParameter, _ Imports, _ error) {
	s := sr.Value()
	out := HeaderParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if sr.Ref() != nil {
		out.FieldTypeName = sr.Ref().Name
	}
	var ims Imports
	var err error
	out.Type, ims, err = NewParameterSchema(s.Schema.Value())
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	out.Required = s.Required
	return &out, ims, nil
}

type CookieParameter struct {
	specification.CookieParameter
}
