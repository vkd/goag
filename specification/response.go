package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Response struct {
	NoRef[Response]

	Description string
	Headers     Map[Ref[Header]]
	Content     Map[*MediaType]
	Links       Map[Ref[Link]]

	UsedIn []ResponseUsedIn
}

func NewResponse(s *openapi3.Response, components Components, opts SchemaOptions) (*Response, error) {
	out := &Response{
		Content: NewMap[*MediaType, *openapi3.MediaType](s.Content, func(mt *openapi3.MediaType) *MediaType {
			return NewMediaType(mt, components.Schemas, opts)
		}),
	}
	var err error
	out.Headers, err = NewMapSelf[Header, *openapi3.HeaderRef](s.Headers, func(h *openapi3.HeaderRef) (Ref[Header], error) {
		if h.Ref != "" {
			v, ok := components.Headers.Get(h.Ref)
			if !ok {
				return nil, fmt.Errorf("component header %q: not found", h.Ref)
			}
			return NewRef(v), nil
		}
		return NewHeader(h.Value, components.Schemas, opts), nil
	})
	if err != nil {
		return nil, fmt.Errorf("new headers: %w", err)
	}

	out.Links, err = NewMapSelf[Link, *openapi3.LinkRef](s.Links, func(lr *openapi3.LinkRef) (Ref[Link], error) {
		if lr.Ref != "" {
			v, ok := components.Links.Get(lr.Ref)
			if !ok {
				return nil, fmt.Errorf("components link %q: not found", lr.Ref)
			}
			return NewRef(v), nil
		}
		return NewLink(lr.Value), nil
	})
	if err != nil {
		return nil, fmt.Errorf("new links: %w", err)
	}

	if s.Description != nil {
		out.Description = *s.Description
	}

	return out, nil
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

type ResponseUsedIn struct {
	Operation *Operation
	Status    string
}
