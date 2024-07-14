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
	Name       OperationName
	HTTPMethod specification.HTTPMethod
	Method     specification.HTTPMethodTitle
	PathRaw    string

	RequestTypeName  string
	ResponseTypeName string

	Queries []ClientOperationQueryTemplate
	Headers []ClientOperationHeaderTemplate

	IsRequestBody bool

	Responses       []ClientResponseTemplate
	DefaultResponse *ClientResponseTemplate

	PathFormat []Render
}

func NewClientOperation(o *Operation) ClientOperationTemplate {
	c := ClientOperationTemplate{
		Name:       o.Name,
		HTTPMethod: o.HTTPMethod,
		Method:     o.Method,
		PathRaw:    o.Path.Raw,

		RequestTypeName:  o.RequestTypeName,
		ResponseTypeName: o.ResponseTypeName,
	}
	if requestBody, ok := o.Operation.RequestBody.Get(); ok {
		c.IsRequestBody = requestBody.Value().Content.Has("application/json")
	}

	for _, e := range o.Params.Query.List {
		renderFormatFn := e.V.Type.RenderFormat
		switch eType := e.V.Type.(type) {
		case CustomType:
			renderFormatFn = eType.RenderFormatStrings
		}
		q := ClientOperationQueryTemplate{
			Name:      e.Name,
			Required:  e.V.Required,
			FieldName: e.V.FieldName,

			RenderFormat: renderFormatFn,
			IsMultivalue: e.V.Type.IsMultivalue(),
		}
		switch tp := e.V.Type.(type) {
		case SliceType:
			switch tp.Items.(type) {
			case StringType:
			default:
				q.ExecuteMultilineFormat = tp.RenderFormatStringsMultiline
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
				IsMultivalue:      h.IsMultivalue,
				Key:               h.Key,
				Required:          h.Required,
				SchemaParseString: h.Schema.ParseString,
				FieldName:         h.FieldName,
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
				IsMultivalue:      h.IsMultivalue,
				Key:               h.Key,
				Required:          h.Required,
				SchemaParseString: h.Schema.ParseString,
				FieldName:         h.FieldName,
			})
		}
		c.DefaultResponse = &t
	}

	var v string
	for _, dir := range o.Path.Dirs {
		if !dir.IsVariable {
			v += "/" + dir.V
			continue
		}
		c.PathFormat = append(c.PathFormat, QuotedRender(v+"/"))
		v = ""

		param := dir.Param
		c.PathFormat = append(c.PathFormat, RenderFunc(func() (string, error) {
			return param.Type.RenderFormat("request.Path." + param.FieldName)
		}))
	}
	if v != "" {
		c.PathFormat = append(c.PathFormat, QuotedRender(v))
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

	RenderFormat func(from string) (string, error)
	IsMultivalue bool

	ExecuteMultilineFormat func(to, from string) (string, error)
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
	IsMultivalue      bool
	Key               string
	Required          bool
	SchemaParseString ParserFunc
	FieldName         string
}
