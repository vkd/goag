package specification

import "github.com/getkin/kin-openapi/openapi3"

type Header struct {
	NoRef[Header]

	Description string
	Required    bool
	Deprecated  bool

	Schema Ref[Schema]
}

func NewHeader(s *openapi3.Header, schemas ComponentsSchemas, opts SchemaOptions) *Header {
	return &Header{
		Description: s.Description,
		Required:    s.Required,
		Deprecated:  s.Deprecated,

		Schema: NewSchemaRef(s.Schema, schemas, opts),
	}
}

var _ Ref[Header] = (*Header)(nil)

func (h *Header) Value() *Header { return h }
