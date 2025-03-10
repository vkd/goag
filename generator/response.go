package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type Response struct {
	*specification.Response

	Name string

	Headers []ResponseHeader

	// Content specification.Map[specification.Ref[SchemaType]]
	ContentJSON Maybe[ResponseContentSchema]
	ContentBody Maybe[string]
}

func NewResponse(handlerName OperationName, status string, response *specification.Response, components Componenter, cfg Config) (*Response, Imports, error) {
	r := &Response{Response: response}
	r.Name = string(handlerName) + "Response" + strings.Title(status)

	var imports Imports

	for _, c := range response.Content.List {
		switch c.Name {
		case "application/json":
			r.Name += "JSON"

			s, ims, err := NewSchema(c.V.Schema, NamedComponenter{Componenter: components, Name: string(handlerName) + "Response" + status + "JSONBody"}, cfg)
			if err != nil {
				return nil, nil, fmt.Errorf("schema for %q content response: %w", c.Name, err)
			}
			imports = append(imports, ims...)

			r.ContentJSON = Just(ResponseContentSchema{
				Type: s,
			})
		default:
			r.ContentBody = Just(c.Name)
		}
	}

	r.Headers = make([]ResponseHeader, 0, len(response.Headers.List))

	for _, header := range response.Headers.List {
		h, ims, err := NewResponseHeader(header.Name, header.V, components, cfg)
		if err != nil {
			return nil, nil, fmt.Errorf("new response %q: %w", header.Name, err)
		}
		r.Headers = append(r.Headers, h)
		imports = append(imports, ims...)
	}

	return r, imports, nil
}

type ResponseHeader struct {
	FieldName string
	Key       string
	Required  bool
	Schema    interface {
		GoTypeRender
		Parser
	}
	IsCustomType bool
}

func NewResponseHeader(name string, ref specification.Ref[specification.Header], components Componenter, cfg Config) (zero ResponseHeader, _ Imports, _ error) {
	schema, ims, err := NewSchema(ref.Value().Schema, components, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}
	var s interface {
		GoTypeRender
		Parser
	} = schema
	if !ref.Value().Required {
		s = NewOptionalType(schema, cfg)
	}

	h := ResponseHeader{
		FieldName: PublicFieldName(name),
		Key:       name,
		Required:  ref.Value().Required,
		Schema:    s,
	}
	h.IsCustomType = schema.IsCustom()
	return h, ims, nil
}

type ResponseContentSchema struct {
	Type Schema
}
