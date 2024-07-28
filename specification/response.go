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
	responseContent, err := NewMap[*MediaType, *openapi3.MediaType](s.Content, func(mt *openapi3.MediaType) (*MediaType, error) {
		return NewMediaType(mt, components.Schemas, opts)
	})
	if err != nil {
		return nil, fmt.Errorf("new response content map: %w", err)
	}

	out := &Response{
		Content: responseContent,
	}
	out.Headers, err = NewMapRefSelfSource[Header, *openapi3.HeaderRef](s.Headers, func(h *openapi3.HeaderRef, _ Sourcer[Header]) (ref string, _ Ref[Header], _ error) {
		if h.Ref != "" {
			return h.Ref, nil, nil
		}
		header, err := NewHeader(h.Value, components.Schemas, opts)
		if err != nil {
			return "", nil, fmt.Errorf("new header: %w", err)
		}
		return "", header, nil
	}, components.Headers, "")
	if err != nil {
		return nil, fmt.Errorf("new headers: %w", err)
	}

	out.Links, err = NewMapRefSelfSource[Link, *openapi3.LinkRef](s.Links, func(lr *openapi3.LinkRef, _ Sourcer[Link]) (ref string, _ Ref[Link], _ error) {
		if lr.Ref != "" {
			return lr.Ref, nil, nil
		}
		return "", NewLink(lr.Value), nil
	}, components.Links, "")
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
