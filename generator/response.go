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
}

func NewResponse(handlerName OperationName, status string, response *specification.Response, components Components, cfg Config) (*Response, Imports, error) {
	r := &Response{Response: response}
	r.Name = string(handlerName) + "Response" + strings.Title(status)
	if response.Content.Has("application/json") {
		r.Name += "JSON"
	}

	var imports Imports

	for _, c := range response.Content.List {
		switch c.Name {
		case "application/json":
			s, ims, err := NewSchema(c.V.Schema, components)
			if err != nil {
				return nil, nil, fmt.Errorf("schema for %q content response: %w", c.Name, err)
			}
			imports = append(imports, ims...)

			r.ContentJSON = Just(ResponseContentSchema{
				Spec: c.V.Schema,
				Type: s,
			})
		default:
		}
	}

	headers := make([]ResponseHeader, len(response.Headers.List))
	headersMap := make(map[*specification.Object[string, specification.Ref[specification.Header]]]*ResponseHeader, len(response.Headers.List))
	for i, c := range response.Headers.List {
		headersMap[c] = &headers[i]
	}

	for _, header := range response.Headers.List {
		h, ims, err := NewResponseHeader(header.Name, header.V, components, cfg)
		if err != nil {
			return nil, nil, fmt.Errorf("new response %q: %w", header.Name, err)
		}
		imports = append(imports, ims...)
		*headersMap[header] = h
	}

	r.Headers = headers

	return r, imports, nil
}

type ResponseHeader struct {
	Spec *specification.Header

	FieldName string
	Key       string
	Required  bool
	Schema    interface {
		Render
		Parser
	}
	IsCustomType bool
}

func NewResponseHeader(name string, ref specification.Ref[specification.Header], components Components, cfg Config) (zero ResponseHeader, _ Imports, _ error) {
	s, ims, err := NewSchema(ref.Value().Schema, components)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}

	if !ref.Value().Required {
		s = NewOptionalType(s, cfg)
	}

	h := ResponseHeader{
		Spec: ref.Value(),

		FieldName: PublicFieldName(name),
		Key:       name,
		Required:  ref.Value().Required,
		Schema:    s,
	}
	if _, ok := s.(CustomType); ok {
		h.IsCustomType = true
	}
	return h, ims, nil
}

type ResponseContentSchema struct {
	Spec specification.Ref[specification.Schema]
	Type SchemaType
}
