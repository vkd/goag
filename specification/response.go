package specification

import "github.com/getkin/kin-openapi/openapi3"

type Response struct {
	NoRef[Response]

	s *openapi3.Response

	Description string
	Headers     Map[Ref[Header]]
	Content     Map[*MediaType]
	Links       Map[Ref[Link]]
}

func NewResponse(s *openapi3.Response, components Components) *Response {
	return &Response{
		s: s,

		Description: *s.Description,

		Headers: NewMapRefSource[Header, *openapi3.HeaderRef](s.Headers, func(h *openapi3.HeaderRef) (ref string, _ Ref[Header]) {
			if h.Ref != "" {
				return h.Ref, nil
			}
			return "", NewHeader(h.Value, components.Schemas)
		}, components.Headers, ""),

		Content: NewMap[*MediaType, *openapi3.MediaType](s.Content, func(mt *openapi3.MediaType) *MediaType {
			return NewMediaType(mt, components.Schemas)
		}),

		Links: NewMapRefSource[Link, *openapi3.LinkRef](s.Links, func(lr *openapi3.LinkRef) (ref string, _ Ref[Link]) {
			if lr.Ref != "" {
				return lr.Ref, nil
			}
			return "", NewLink(lr.Value)
		}, components.Links, ""),
	}
}

var _ Ref[Response] = (*Response)(nil)

func (r *Response) Value() *Response { return r }

type Link struct {
	NoRef[Link]
}

func NewLink(s *openapi3.Link) *Link {
	return &Link{}
}

var _ Ref[Link] = (*Link)(nil)

func (l *Link) Value() *Link { return l }
