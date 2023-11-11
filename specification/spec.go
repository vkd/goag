package specification

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

type Spec struct {
	Swagger *openapi3.Swagger

	Paths      []*PathItem
	Operations []*Operation
}

func ParseSwagger(spec *openapi3.Swagger) (*Spec, error) {
	var s Spec

	for _, path := range sortedKeys(spec.Paths) {
		p, err := NewPath(path)
		if err != nil {
			return nil, fmt.Errorf("validate path %q: %w", path, err)
		}
		pathItem := spec.Paths[path]
		pi := &PathItem{
			Path:     p,
			PathItem: pathItem,
		}
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
	return &s, nil
}

type PathItem struct {
	Path     Path
	PathItem *openapi3.PathItem

	Operations []*Operation
}

type Operation struct {
	PathItem   *PathItem
	HTTPMethod string
	Method     string
	Operation  *openapi3.Operation

	Parameters struct {
		Path    []PathParameter
		Query   []QueryParameter
		Headers []HeaderParameter
	}

	DefaultResponse *Response
	Responses       []Response
}

func NewOperation(pi *PathItem, method httpMethod, operation *openapi3.Operation) *Operation {
	o := &Operation{
		PathItem:   pi,
		HTTPMethod: method.HTTP,
		Method:     method.Title,
		Operation:  operation,
	}

	for _, param := range operation.Parameters {
		switch param.Value.In {
		case openapi3.ParameterInPath:
			o.Parameters.Path = append(o.Parameters.Path, PathParameter{
				RefName:     param.Ref,
				Name:        param.Value.Name,
				Description: param.Value.Description,
				Schema:      NewSchema(param.Value.Schema),
			})
		case openapi3.ParameterInQuery:
			o.Parameters.Query = append(o.Parameters.Query, QueryParameter{
				RefName:     param.Ref,
				Name:        param.Value.Name,
				Description: param.Value.Description,
				Required:    param.Value.Required,
				Schema:      NewSchema(param.Value.Schema),
			})
		case openapi3.ParameterInHeader:
			o.Parameters.Headers = append(o.Parameters.Headers, HeaderParameter{
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
			o.DefaultResponse = &defaultResponse
		} else {
			o.Responses = append(o.Responses, NewResponse(responseStatusCode, o, response))
		}
	}

	return o
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

func NewResponse(responseStatusCode string, o *Operation, r *openapi3.ResponseRef) Response {
	out := Response{
		StatusCode: responseStatusCode,
		Operation:  o,
		Spec:       r.Value,

		RefName: r.Ref,
		Headers: Headers(r.Value.Headers),
	}
	return out
}

type Schema struct {
	Ref    string
	Schema *openapi3.Schema
}

func NewSchema(schema *openapi3.SchemaRef) Schema {
	return Schema{
		Ref:    schema.Ref,
		Schema: schema.Value,
	}
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
