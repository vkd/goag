package generator

import (
	"github.com/vkd/goag/specification"
)

type ClientTemplate struct {
	Operations []ClientOperationTemplate
}

func NewClient(s *specification.Spec, ops []*Operation) ClientTemplate {
	var c ClientTemplate
	c.Operations = make([]ClientOperationTemplate, 0, len(ops))
	for _, o := range ops {
		co := NewClientOperation(o)
		c.Operations = append(c.Operations, co)
	}
	return c
}

func (c ClientTemplate) Render() (string, error) {
	return ExecuteTemplate("Client", c)
}

type ClientOperationTemplate struct {
	Name        OperationName
	Description string
	HTTPMethod  specification.HTTPMethod
	Method      specification.HTTPMethodTitle
	PathRaw     string

	PathParams []PathStringBuilder

	RequestTypeName  string
	ResponseTypeName string

	Queries []ClientOperationQueryTemplate
	Headers []ClientOperationHeaderTemplate

	IsRequestBody bool

	Responses       []ClientResponseTemplate
	DefaultResponse *ClientResponseTemplate
}

func NewClientOperation(o *Operation) ClientOperationTemplate {
	c := ClientOperationTemplate{
		Name:        o.Name,
		Description: o.Description,
		HTTPMethod:  o.HTTPMethod,
		Method:      o.Method,
		PathRaw:     o.Path.Raw,

		PathParams: o.Path.StringBuilder(),

		RequestTypeName:  o.RequestTypeName,
		ResponseTypeName: o.ResponseTypeName,
	}

	if requestBody, ok := o.Operation.RequestBody.Get(); ok {
		c.IsRequestBody = requestBody.Value().Content.Has("application/json")
	}

	for _, e := range o.Params.Query.List {
		q := ClientOperationQueryTemplate{
			Name:      e.Name,
			Required:  e.V.Required,
			FieldName: e.V.FieldName,

			RenderFormatStrings: e.V.Type.RenderFormatStrings,
			IsMultivalue:        e.V.Type.IsMultivalue(),
		}
		switch tp := e.V.Type.(type) {
		case SliceType:
			switch tp.Items.(type) {
			case StringType:
			default:
			}
		}
		c.Queries = append(c.Queries, q)
	}

	for _, h := range o.Params.Headers.List {
		c.Headers = append(c.Headers, ClientOperationHeaderTemplate{
			Name:      h.V.Name,
			FieldName: h.V.FieldName,
			Required:  h.V.Required,
			Type:      h.V.Type,
		})
	}

	for _, e := range o.Responses {
		t := ClientResponseTemplate{
			Name:             e.Name,
			StatusCode:       e.StatusCode,
			ContentJSON:      e.ContentJSON.Set,
			ComponentRefName: e.Name,
		}
		if e.ComponentRef != nil {
			t.ComponentRefName = e.ComponentRef.Name + "Response"
		}
		for _, h := range e.Headers {
			t.Headers = append(t.Headers, ClientResponseHeaderTemplate{
				IsMultivalue:       h.IsMultivalue,
				Key:                h.Key,
				Required:           h.Required,
				SchemaParseStrings: h.Schema.ParseStrings,
				FieldName:          h.FieldName,
			})
		}
		c.Responses = append(c.Responses, t)
	}
	if o.DefaultResponse != nil {
		e := *o.DefaultResponse

		t := ClientResponseTemplate{
			Name:             e.Name,
			StatusCode:       e.StatusCode,
			ContentJSON:      e.ContentJSON.Set,
			ComponentRefName: e.Name,
		}
		if e.ComponentRef != nil {
			t.ComponentRefName = e.ComponentRef.Name + "Response"
		}
		for _, h := range e.Headers {
			t.Headers = append(t.Headers, ClientResponseHeaderTemplate{
				IsMultivalue:       h.IsMultivalue,
				Key:                h.Key,
				Required:           h.Required,
				SchemaParseStrings: h.Schema.ParseStrings,
				FieldName:          h.FieldName,
			})
		}
		c.DefaultResponse = &t
	}

	return c
}

func (c ClientOperationTemplate) Render() (string, error) {
	return ExecuteTemplate("ClientOperation", c)
}

type ClientOperationQueryTemplate struct {
	Name      string
	Required  bool
	FieldName string

	RenderFormatStrings func(to, from string, isNew bool) (string, error)
	IsMultivalue        bool
}

type ClientOperationHeaderTemplate struct {
	Name      string
	Required  bool
	FieldName string

	Type Formatter
}

type ClientResponseTemplate struct {
	Name             string
	StatusCode       string
	ComponentRefName string
	Headers          []ClientResponseHeaderTemplate
	ContentJSON      bool
}

func (c ClientResponseTemplate) Render() (string, error) {
	return ExecuteTemplate("ClientResponse", c)
}

type ClientResponseHeaderTemplate struct {
	IsMultivalue       bool
	Key                string
	Required           bool
	SchemaParseStrings ParserFunc
	FieldName          string
}
