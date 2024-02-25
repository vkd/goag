package specification

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

type Spec struct {
	OpenAPI string
	Info    Info
	Servers []Server

	PathItems  []*PathItem
	Operations []*Operation

	Components Components

	Security []SecurityRequirement
}

func ParseSwagger(spec *openapi3.Swagger) (*Spec, error) {
	s := &Spec{
		OpenAPI:    spec.OpenAPI,
		Info:       NewInfo(spec.Info),
		Servers:    NewServers(spec.Servers),
		Components: NewComponents(spec.Components),
	}

	s.Security = NewSecurityRequirements(spec.Security, s.Components.SecuritySchemes)

	for _, pathKey := range sortedKeys(spec.Paths) {
		p, err := NewPathOld2(pathKey)
		if err != nil {
			return nil, fmt.Errorf("parse path %q: %w", pathKey, err)
		}
		pathItem := spec.Paths[pathKey]
		pi := NewPathItem(pathKey)
		pi.Path = p
		pi.PathItem = pathItem
		pi.PathOld, _ = NewPathOld(pathKey)
		for _, method := range httpMethods() {
			operation := pathItem.GetOperation(string(method.HTTP))
			if operation == nil {
				continue
			}

			o, err := NewOperation(pi, pathKey, method, operation, s.Security, spec.Components, s.Components.SecuritySchemes, s.Components)
			if err != nil {
				return nil, fmt.Errorf("new operation for path=%q method=%q: %w", pi.Path.Spec, method.HTTP, err)
			}
			pi.Operations = append(pi.Operations, o)
			s.Operations = append(s.Operations, o)
		}
		s.PathItems = append(s.PathItems, pi)
	}

	return s, nil
}

type PathParameters []*PathParameterOld

func (ps PathParameters) Get(name string) (*PathParameterOld, bool) {
	for _, p := range ps {
		if p.Name == name {
			return p, true
		}
	}
	return nil, false
}

type PathParameterOld struct {
	RefName     string
	Name        string
	Description string
	Schema      Ref[Schema]
}

type QueryParameterOld struct {
	RefName     string
	Name        string
	Description string
	Required    bool
	Schema      Ref[Schema]
}

type HeaderParameterOld struct {
	RefName     string
	Name        string
	Description string
	Required    bool
	Schema      Ref[Schema]
}

// type Parameter struct {
// 	RefName     string
// 	Name        string
// 	Description string
// 	Required    bool
// }

type ResponseOld struct {
	StatusCode string
	Operation  *Operation
	Spec       *openapi3.Response

	RefName string
	Headers []HeaderOld
}

func NewResponseOld(responseStatusCode string, o *Operation, r *openapi3.ResponseRef) *ResponseOld {
	out := &ResponseOld{
		StatusCode: responseStatusCode,
		Operation:  o,
		Spec:       r.Value,

		RefName: r.Ref,
		Headers: Headers(r.Value.Headers),
	}
	return out
}

type Schema struct {
	NoRef[Schema]
	// Ref         string
	Type        string
	Items       Ref[Schema]
	Properties  []SchemaProperty
	AllOf       []Ref[Schema]
	Description string

	Schema *openapi3.Schema
}

func NewSchemaRef(schema *openapi3.SchemaRef, components ComponentsSchemas) Ref[Schema] {
	if schema.Ref != "" {
		v, ok := components.Get(schema.Ref)
		if !ok {
			panic(fmt.Sprintf("%q: not found in components", schema.Ref))
		}
		return NewRefObject[Schema](v)
	}
	return NewSchema(schema.Value, components)
	// out := Schema{
	// 	// Ref:         schema.Ref,
	// 	Type:        schema.Type,
	// 	Schema:      schema,
	// 	Description: schema.Description,
	// }
	// if schema.Items != nil {
	// 	s := NewSchema(schema.Items)
	// 	out.Items = &s
	// }
	// for _, name := range sortedKeys(schema.Properties) {
	// 	out.Properties = append(out.Properties, SchemaProperty{Name: name, Schema: NewSchema(schema.Properties[name])})
	// }
	// for _, a := range schema.AllOf {
	// 	out.AllOf = append(out.AllOf, NewSchema(a))
	// }
	// return out
}

func NewSchema(schema *openapi3.Schema, components ComponentsSchemas) *Schema {
	out := Schema{
		// Ref:         schema.Ref,
		Type:        schema.Type,
		Schema:      schema,
		Description: schema.Description,
	}
	if schema.Items != nil {
		out.Items = NewSchemaRef(schema.Items, components)
	}
	for _, name := range sortedKeys(schema.Properties) {
		out.Properties = append(out.Properties, SchemaProperty{Name: name, Schema: NewSchemaRef(schema.Properties[name], components)})
	}
	for _, a := range schema.AllOf {
		out.AllOf = append(out.AllOf, NewSchemaRef(a, components))
	}
	return &out
}

var _ Ref[Schema] = (*Schema)(nil)

func (s *Schema) Value() *Schema { return s }

type SchemaProperty struct {
	Name   string
	Schema Ref[Schema]
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
	HTTP  HTTPMethod
	Title HTTPMethodTitle
}

// HTTPMethod - http.MethodGet, http.MethodPost, ...
type HTTPMethod string

// HTTPMethodTitle - Get, Post, ...
type HTTPMethodTitle string
