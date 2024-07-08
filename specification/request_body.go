package specification

import "github.com/getkin/kin-openapi/openapi3"

type RequestBody struct {
	NoRef[RequestBody]

	Description string
	Content     Map[*MediaType]
	Required    bool
}

func NewRequestBody(r *openapi3.RequestBody, components ComponentsSchemas, opts SchemaOptions) *RequestBody {
	return &RequestBody{
		Description: r.Description,
		Content:     NewMap[*MediaType, *openapi3.MediaType](r.Content, func(mt *openapi3.MediaType) *MediaType { return NewMediaType(mt, components, opts) }),
		Required:    r.Required,
	}
}

var _ Ref[RequestBody] = (*RequestBody)(nil)

func (r *RequestBody) Value() *RequestBody { return r }
