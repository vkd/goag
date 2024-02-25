package specification

import "github.com/getkin/kin-openapi/openapi3"

type MediaType struct {
	Schema Ref[Schema]
}

func NewMediaType(s *openapi3.MediaType, components ComponentsSchemas) *MediaType {
	return &MediaType{
		Schema: NewSchemaRef(s.Schema, components),
	}
}
