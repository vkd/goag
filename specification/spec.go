package specification

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

type Spec struct {
	Swagger *openapi3.Swagger

	OpenAPI string
	Info    Info

	Paths      []*PathItem
	Operations []*Operation

	Components Components
}

func ParseSwagger(spec *openapi3.Swagger) (*Spec, error) {
	s := &Spec{
		OpenAPI: spec.OpenAPI,
		Info:    NewInfo(spec.Info),
	}

	for _, path := range sortedKeys(spec.Paths) {
		p, err := NewPath(path)
		if err != nil {
			return nil, fmt.Errorf("parse path %q: %w", path, err)
		}
		pathItem := spec.Paths[path]
		pi := &PathItem{
			Path:     p,
			PathItem: pathItem,
			Spec:     s,
		}
		pi.PathOld, _ = NewPathOld(path)
		for _, method := range httpMethods() {
			operation := pathItem.GetOperation(method.HTTP)
			if operation == nil {
				continue
			}
			o := NewOperation(pi, method, operation)
			pi.Operations = append(pi.Operations, o)
			s.Operations = append(s.Operations, o)
		}
		s.Paths = append(s.Paths, pi)
	}

	s.Components = NewComponents(spec.Components)

	return s, nil
}

type PathItem struct {
	Path     Path
	PathOld  PathOld
	PathItem *openapi3.PathItem
	Spec     *Spec

	Operations []*Operation
}

type Operation struct {
	PathItem *PathItem

	HTTPMethod  string
	Method      string
	OperationID string

	Operation *openapi3.Operation

	Parameters struct {
		Path    PathParameters
		Query   []QueryParameter
		Headers []*HeaderParameter
	}

	DefaultResponse *Response
	Responses       []*Response
}

func NewOperation(pi *PathItem, method httpMethod, operation *openapi3.Operation) *Operation {
	o := &Operation{
		PathItem:    pi,
		HTTPMethod:  method.HTTP,
		Method:      method.Title,
		OperationID: operation.OperationID,

		Operation: operation,
	}

	for _, param := range append(append(openapi3.Parameters{}, pi.PathItem.Parameters...), operation.Parameters...) {
		switch param.Value.In {
		case openapi3.ParameterInPath:
			p, ok := pi.Path.Params.Get(param.Value.Name)
			if !ok {
				p = &PathParameter{Name: param.Value.Name}
			}
			p.RefName = param.Ref
			p.Description = param.Value.Description
			p.Schema = NewSchema(param.Value.Schema)

			o.Parameters.Path = append(o.Parameters.Path, p)
		case openapi3.ParameterInQuery:
			o.Parameters.Query = append(o.Parameters.Query, QueryParameter{
				RefName:     param.Ref,
				Name:        param.Value.Name,
				Description: param.Value.Description,
				Required:    param.Value.Required,
				Schema:      NewSchema(param.Value.Schema),
			})
		case openapi3.ParameterInHeader:
			o.Parameters.Headers = append(o.Parameters.Headers, &HeaderParameter{
				RefName:     param.Ref,
				Name:        param.Value.Name,
				Description: param.Value.Description,
				Required:    param.Value.Required,
				Schema:      NewSchema(param.Value.Schema),
			})
		}
	}

	for _, responseStatusCode := range sortedKeys(operation.Responses) {
		response := operation.Responses[responseStatusCode]
		if responseStatusCode == "default" {
			defaultResponse := NewResponse(responseStatusCode, o, response)
			o.DefaultResponse = defaultResponse
		} else {
			o.Responses = append(o.Responses, NewResponse(responseStatusCode, o, response))
		}
	}

	return o
}

type PathParameters []*PathParameter

func (ps PathParameters) Get(name string) (*PathParameter, bool) {
	for _, p := range ps {
		if p.Name == name {
			return p, true
		}
	}
	return nil, false
}

type PathParameter struct {
	RefName     string
	Name        string
	Description string
	Schema      Schema
}

type QueryParameter struct {
	RefName     string
	Name        string
	Description string
	Required    bool
	Schema      Schema
}

type HeaderParameter struct {
	RefName     string
	Name        string
	Description string
	Required    bool
	Schema      Schema
}

// type Parameter struct {
// 	RefName     string
// 	Name        string
// 	Description string
// 	Required    bool
// }

type Response struct {
	StatusCode string
	Operation  *Operation
	Spec       *openapi3.Response

	RefName string
	Headers []Header
}

func NewResponse(responseStatusCode string, o *Operation, r *openapi3.ResponseRef) *Response {
	out := &Response{
		StatusCode: responseStatusCode,
		Operation:  o,
		Spec:       r.Value,

		RefName: r.Ref,
		Headers: Headers(r.Value.Headers),
	}
	return out
}

type Schema struct {
	Ref         string
	Type        string
	Items       *Schema
	Properties  []SchemaProperty
	AllOf       []Schema
	Description string

	Schema *openapi3.Schema
}

func NewSchema(schema *openapi3.SchemaRef) Schema {
	out := Schema{
		Ref:         schema.Ref,
		Type:        schema.Value.Type,
		Schema:      schema.Value,
		Description: schema.Value.Description,
	}
	if schema.Value.Items != nil {
		s := NewSchema(schema.Value.Items)
		out.Items = &s
	}
	for _, name := range sortedKeys(schema.Value.Properties) {
		out.Properties = append(out.Properties, SchemaProperty{Name: name, Schema: NewSchema(schema.Value.Properties[name])})
	}
	for _, a := range schema.Value.AllOf {
		out.AllOf = append(out.AllOf, NewSchema(a))
	}
	return out
}

type SchemaProperty struct {
	Name string
	Schema
}

func sortedKeys[T any](m map[string]T) (out []string) {
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func httpMethods() []httpMethod {
	return []httpMethod{
		{http.MethodGet, "Get"},
		{http.MethodPost, "Post"},
		{http.MethodPatch, "Patch"},
		{http.MethodPut, "Put"},
		{http.MethodDelete, "Delete"},
		{http.MethodConnect, "Connect"},
		{http.MethodHead, "Head"},
		{http.MethodOptions, "Options"},
		{http.MethodTrace, "Trace"},
	}
}

type httpMethod struct {
	HTTP  string
	Title string
}
