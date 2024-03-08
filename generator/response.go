package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type Response struct {
	*specification.Response

	Name       string
	StatusCode string

	Headers []ResponseHeader

	// Content specification.Map[specification.Ref[SchemaType]]
	ContentJSON Optional[Schema]
}

func NewResponse(handlerName OperationName, status string, response *specification.Response) (*Response, Imports, error) {
	r := &Response{Response: response}
	r.Name = string(handlerName) + "Response" + strings.Title(status)
	if response.Content.Has("application/json") {
		r.Name += "JSON"
	}

	r.StatusCode = status

	var imports Imports

	for _, c := range response.Content.List {
		switch c.Name {
		case "application/json":
			s, ims, err := NewSchema(c.V.Schema)
			if err != nil {
				return nil, nil, fmt.Errorf("schema for %q content response: %w", c.Name, err)
			}
			imports = append(imports, ims...)

			r.ContentJSON = NewOptional[Schema](Schema{
				Spec: c.V.Schema,
				Type: s,
			})
		default:
		}
	}

	for _, header := range response.Headers.List {
		h, ims, err := NewResponseHeader(header.Name, header.V.Value())
		if err != nil {
			return nil, nil, fmt.Errorf("new response %q: %w", header.Name, err)
		}
		imports = append(imports, ims...)
		r.Headers = append(r.Headers, h)
	}
	return r, imports, nil
}

type ResponseHeader struct {
	Spec *specification.Header

	FieldName string
	Key       string
	Schema    SchemaType
}

func NewResponseHeader(name string, spec *specification.Header) (zero ResponseHeader, _ Imports, _ error) {
	s, ims, err := NewSchema(spec.Schema)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}
	h := ResponseHeader{
		Spec: spec,

		FieldName: PublicFieldName(name),
		Key:       name,
		Schema:    s,
	}
	return h, ims, nil
}
