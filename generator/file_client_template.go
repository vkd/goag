package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type ClientTemplate struct {
	Operations []ClientOperationTemplate
}

func NewClient(s *specification.Spec, ops []*Operation) (zero ClientTemplate, _ error) {
	var c ClientTemplate
	c.Operations = make([]ClientOperationTemplate, 0, len(ops))
	for _, o := range ops {
		co, err := NewClientOperation(o)
		if err != nil {
			return zero, fmt.Errorf("new client for %q operation: %w", o.Name, err)
		}
		c.Operations = append(c.Operations, co)
	}
	return c, nil
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

func NewClientOperation(o *Operation) (zero ClientOperationTemplate, _ error) {
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
		c.Queries = append(c.Queries, ClientOperationQueryTemplate{
			Name:          e.Name,
			Required:      e.V.Required,
			ExecuteFormat: e.V.ExecuteFormat,
			FieldName:     e.V.FieldName,
		})
	}

	for _, h := range o.Params.Headers.List {
		c.Headers = append(c.Headers, ClientOperationHeaderTemplate{
			Name:      h.V.Name,
			FieldName: h.V.FieldName,
			Required:  h.V.Required,
			Type:      h.V.Type,
		})
	}

	if len(o.Security) > 0 {
		for _, sr := range o.Security {
			if sr.Scheme.Type != specification.SecuritySchemeTypeHTTP {
				continue
			}
			if sr.Scheme.Scheme != "bearer" {
				continue
			}
			c.Headers = append(c.Headers, ClientOperationHeaderTemplate{
				Name:      "Authorization",
				FieldName: "Authorization",
				Required:  len(o.Security) == 1,
				Type:      StringType{},
			})
		}
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
		gp, ok := o.Params.Path.Get(param.Name)
		if !ok {
			return zero, fmt.Errorf("%q path parameter: not found in %q operation", param.Name, o.Name)
		}
		c.PathFormat = append(c.PathFormat, RenderFunc(func() (string, error) {
			return gp.V.Type.RenderFormat("request.Path." + gp.V.FieldName)
		}))
	}
	if v != "" {
		c.PathFormat = append(c.PathFormat, QuotedRender(v))
	}
	return c, nil
}

func (c ClientOperationTemplate) Render() (string, error) {
	return ExecuteTemplate("ClientOperation", c)
}

type ClientOperationQueryTemplate struct {
	Name          string
	Required      bool
	ExecuteFormat func(to, from string) (string, error)
	FieldName     string
}

type ClientOperationHeaderTemplate struct {
	Name      string
	FieldName string
	Required  bool
	Type      Formatter
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
