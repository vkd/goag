package specification

import "github.com/getkin/kin-openapi/openapi3"

type Header struct {
	NoRef[Header]

	Description string
	Required    bool
	Deprecated  bool

	Schema Ref[Schema]
}

func NewHeader(s *openapi3.Header, schemas ComponentsSchemas) *Header {
	return &Header{
		Description: s.Description,
		Required:    s.Required,
		Deprecated:  s.Deprecated,

		Schema: NewSchemaRef(s.Schema, schemas),
	}
}

var _ Ref[Header] = (*Header)(nil)

func (h *Header) Value() *Header { return h }

func Headers(hs openapi3.Headers) (out []HeaderOld) {
	for _, h := range sortedKeys(hs) {
		out = append(out, NewHeaderOld(h, hs[h]))
	}
	return
}

type HeaderOld struct {
	Spec *openapi3.Header
	Ref  string

	Name      string
	SchemaRef *openapi3.SchemaRef

	Description string
	Required    bool
	Deprecated  bool
}

func NewHeaderOld(name string, h *openapi3.HeaderRef) HeaderOld {
	header := HeaderOld{
		Spec: h.Value,
		Ref:  h.Ref,

		Name:      name,
		SchemaRef: h.Value.Schema,

		Description: h.Value.Description,
		Required:    h.Value.Required,
		Deprecated:  h.Value.Deprecated,
	}
	return header
}
