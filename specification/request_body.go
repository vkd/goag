package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type RequestBody struct {
	NoRef[RequestBody]

	Description string
	Content     Map[*MediaType]
	Required    bool
}

func NewRequestBody(r *openapi3.RequestBody, components ComponentsSchemas, opts SchemaOptions) (*RequestBody, error) {
	contentMap, err := NewMap[*MediaType, *openapi3.MediaType](r.Content, func(mt *openapi3.MediaType) (*MediaType, error) {
		return NewMediaType(mt, components, opts)
	})
	if err != nil {
		return nil, fmt.Errorf("new map: %w", err)
	}
	return &RequestBody{
		Description: r.Description,
		Content:     contentMap,
		Required:    r.Required,
	}, nil
}

var _ Ref[RequestBody] = (*RequestBody)(nil)

func (r *RequestBody) Value() *RequestBody { return r }
