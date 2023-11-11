package specification

import "github.com/getkin/kin-openapi/openapi3"

func Headers(hs openapi3.Headers) (out []Header) {
	for _, h := range sortedKeys(hs) {
		out = append(out, NewHeader(h, hs[h]))
	}
	return
}

type Header struct {
	Spec *openapi3.Header
	Ref  string

	Name      string
	SchemaRef *openapi3.SchemaRef

	Description string
	Required    bool
	Deprecated  bool
}

func NewHeader(name string, h *openapi3.HeaderRef) Header {
	header := Header{
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
