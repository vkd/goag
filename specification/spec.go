package specification

import (
	"encoding/json"
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

func ParseSwagger(spec *openapi3.Swagger, opts SchemaOptions) (*Spec, error) {
	s := &Spec{
		OpenAPI: spec.OpenAPI,
		Info:    NewInfo(spec.Info),
		Servers: NewServers(spec.Servers),
	}
	var err error
	s.Components, err = NewComponents(spec.Components, opts)
	if err != nil {
		return nil, fmt.Errorf("new components: %w", err)
	}

	s.Security = NewSecurityRequirements(spec.Security, s.Components.SecuritySchemes)

	for _, pathKey := range sortedKeys(spec.Paths) {
		pathItem := spec.Paths[pathKey]
		pi := NewPathItem(pathKey)
		pi.PathItem = pathItem
		for _, method := range httpMethods() {
			operation := pathItem.GetOperation(string(method.HTTP))
			if operation == nil {
				continue
			}

			o, err := NewOperation(pi, pathKey, method, operation, s.Security, spec.Components, s.Components.SecuritySchemes, s.Components, opts)
			if err != nil {
				return nil, fmt.Errorf("new operation for path=%q method=%q: %w", pathKey, method.HTTP, err)
			}
			pi.Operations = append(pi.Operations, o)
			s.Operations = append(s.Operations, o)
		}
		s.PathItems = append(s.PathItems, pi)
	}

	for _, resp := range s.Components.Responses.List {
		var usedInDefault, usedInPatterned Maybe[ResponseUsedIn]
		for _, usedIn := range resp.V.Value().UsedIn {
			usedIn := usedIn
			switch usedIn.Status {
			case "default":
				if usedInPatterned.Set {
					return nil, fmt.Errorf("found multiple usages of %q response in operation [%s %s '%s'] and [%s %s '%s']: each response object could be used only on 'default' responses or 'non-default' responses: already used in patterned status code", resp.Name, usedIn.Operation.HTTPMethod, usedIn.Operation.PathRaw, usedIn.Status, usedInPatterned.Value.Operation.HTTPMethod, usedInPatterned.Value.Operation.PathRaw, usedInPatterned.Value.Status)
				}
				usedInDefault = Just(usedIn)
			default:
				if usedInDefault.Set {
					return nil, fmt.Errorf("found multiple usages of %q response in operation [%s %s '%s'] and [%s %s '%s']: each response object could be used only on 'default' responses or 'non-default' responses: already used in default status code", resp.Name, usedIn.Operation.HTTPMethod, usedIn.Operation.PathRaw, usedIn.Status, usedInDefault.Value.Operation.HTTPMethod, usedInDefault.Value.Operation.PathRaw, usedInDefault.Value.Status)
				}
				usedInPatterned = Just(usedIn)
			}
		}
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

	Type        string
	Format      string
	Items       Ref[Schema]
	Properties  []SchemaProperty
	AllOf       []Ref[Schema]
	Description string

	AdditionalProperties Maybe[Ref[Schema]]

	Custom Maybe[string]
}

func NewSchemaRef(schema *openapi3.SchemaRef, components ComponentsSchemas, opts SchemaOptions) Ref[Schema] {
	if schema.Ref != "" {
		v, ok := components.Get(schema.Ref)
		if !ok {
			panic(fmt.Sprintf("%q: not found in components", schema.Ref))
		}
		return NewRefObject[Schema](v)
	}
	return NewSchema(schema.Value, components, opts)
}

type SchemaOptions struct {
	IgnoreCustomType bool
}

func NewSchema(schema *openapi3.Schema, components ComponentsSchemas, opts SchemaOptions) *Schema {
	out := Schema{
		Type:        schema.Type,
		Format:      schema.Format,
		Description: schema.Description,
	}
	if schema.Items != nil {
		out.Items = NewSchemaRef(schema.Items, components, opts)
	}
	for _, name := range sortedKeys(schema.Properties) {
		s := NewSchemaRef(schema.Properties[name], components, opts)
		out.Properties = append(out.Properties, SchemaProperty{Name: name, Schema: s})
	}
	for _, a := range schema.AllOf {
		s := NewSchemaRef(a, components, opts)
		out.AllOf = append(out.AllOf, s)
	}

	if schema.AdditionalProperties != nil {
		s := NewSchemaRef(schema.AdditionalProperties, components, opts)
		out.AdditionalProperties = Just(s)
	} else if schema.AdditionalPropertiesAllowed != nil && *schema.AdditionalPropertiesAllowed {
		out.AdditionalProperties = Just[Ref[Schema]](&Schema{
			Type:   "object",
			Custom: Just("json.RawMessage"),
		})
	}

	if !opts.IgnoreCustomType {
		if v, ok := schema.ExtensionProps.Extensions[ExtTagGoType]; ok {
			if raw, ok := v.(json.RawMessage); ok {
				var s string
				_ = json.Unmarshal(raw, &s)
				out.Custom = Just(s)
			}
		}
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
