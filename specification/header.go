package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Header struct {
	NoRef[Header]

	Description string
	Required    bool
	Deprecated  bool

	Schema Ref[Schema]
}

func NewHeader(s *openapi3.Header, schemas ComponentsSchemas, opts SchemaOptions) (*Header, error) {
	hSchema, err := NewSchemaRef(s.Schema, schemas, opts)
	if err != nil {
		return nil, fmt.Errorf("new schema ref: %w", err)
	}
	return &Header{
		Description: s.Description,
		Required:    s.Required,
		Deprecated:  s.Deprecated,

		Schema: hSchema,
	}, nil
}

var _ Ref[Header] = (*Header)(nil)

func (h *Header) Value() *Header { return h }
