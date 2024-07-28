package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type MediaType struct {
	Schema Ref[Schema]
}

func NewMediaType(s *openapi3.MediaType, components ComponentsSchemas, opts SchemaOptions) (*MediaType, error) {
	mediaSchema, err := NewSchemaRef(s.Schema, components, opts)
	if err != nil {
		return nil, fmt.Errorf("new schema ref: %w", err)
	}
	return &MediaType{
		Schema: mediaSchema,
	}, nil
}
