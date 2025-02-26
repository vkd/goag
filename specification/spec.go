package specification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

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
	servers, err := NewServers(spec.Servers)
	if err != nil {
		return nil, fmt.Errorf("new servers: %w", err)
	}

	s := &Spec{
		OpenAPI: spec.OpenAPI,
		Info:    NewInfo(spec.Info),
		Servers: servers,
	}
	s.Components, err = NewComponents(spec.Components, opts)
	if err != nil {
		return nil, fmt.Errorf("new components: %w", err)
	}

	s.Security, err = NewSecurityRequirements(spec.Security, s.Components.SecuritySchemes)
	if err != nil {
		return nil, fmt.Errorf("new security requirements: %w", err)
	}

	for _, pathKey := range sortedKeys(spec.Paths) {
		pathItem := spec.Paths[pathKey]
		pi := NewPathItem(pathKey)
		for _, method := range httpMethods() {
			operation := pathItem.GetOperation(string(method.HTTP))
			if operation == nil {
				continue
			}

			o, err := NewOperation(pi, pathKey, method, operation, s.Security, spec.Components, s.Components.SecuritySchemes, s.Components, pathItem.Parameters, opts)
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
				if usedInPatterned, ok := usedInPatterned.Get(); ok {
					return nil, fmt.Errorf("found multiple usages of %q response in operation [%s %s '%s'] and [%s %s '%s']: each response object could be used only on 'default' responses or 'non-default' responses: already used in patterned status code", resp.Name, usedIn.Operation.Method.HTTP, usedIn.Operation.PathRaw, usedIn.Status, usedInPatterned.Operation.Method.HTTP, usedInPatterned.Operation.PathRaw, usedInPatterned.Status)
				}
				usedInDefault = Just(usedIn)
			default:
				if usedInDefault, ok := usedInDefault.Get(); ok {
					return nil, fmt.Errorf("found multiple usages of %q response in operation [%s %s '%s'] and [%s %s '%s']: each response object could be used only on 'default' responses or 'non-default' responses: already used in default status code", resp.Name, usedIn.Operation.Method.HTTP, usedIn.Operation.PathRaw, usedIn.Status, usedInDefault.Operation.Method.HTTP, usedInDefault.Operation.PathRaw, usedInDefault.Status)
				}
				usedInPatterned = Just(usedIn)
			}
		}
	}

	return s, nil
}

type Schema struct {
	NoRef[Schema]

	Type          string
	Format        string
	Items         Ref[Schema]
	Properties    []SchemaProperty
	AllOf         []Ref[Schema]
	OneOf         []Ref[Schema]
	Discriminator Discriminator
	Description   string
	Nullable      bool

	AdditionalProperties Maybe[Ref[Schema]]

	Custom     Maybe[string]
	Extentions map[string]string
}

func NewSchemaRef(schema *openapi3.SchemaRef, components Sourcer[Schema], opts SchemaOptions) (Ref[Schema], error) {
	if schema.Ref != "" {
		v, ok := components.Get(schema.Ref)
		if !ok {
			return nil, fmt.Errorf("%q: not found in components", schema.Ref)
		}
		return NewRef[Schema](v), nil
	}
	return NewSchema(schema.Value, components, opts)
}

type SchemaOptions struct {
	IgnoreCustomType bool
}

func NewSchema(schema *openapi3.Schema, components Sourcer[Schema], opts SchemaOptions) (*Schema, error) {
	out := Schema{
		Type:        schema.Type,
		Format:      schema.Format,
		Description: schema.Description,
		Nullable:    schema.Nullable,
	}
	if schema.Items != nil {
		items, err := NewSchemaRef(schema.Items, components, opts)
		if err != nil {
			return nil, fmt.Errorf("new schema ref for items: %w", err)
		}
		out.Items = items
	}
	required := make(map[string]struct{})
	for _, r := range schema.Required {
		required[r] = struct{}{}
	}
	for _, name := range sortedKeys(schema.Properties) {
		s, err := NewSchemaRef(schema.Properties[name], components, opts)
		if err != nil {
			return nil, fmt.Errorf("new schema ref for properties: %w", err)
		}
		var req bool
		if _, ok := required[name]; ok {
			req = true
			delete(required, name)
		}
		out.Properties = append(out.Properties, SchemaProperty{
			Name:     name,
			Schema:   s,
			Required: req,
		})
	}
	if len(required) > 0 {
		var rs []string
		for r := range required {
			rs = append(rs, r)
		}
		return nil, fmt.Errorf("unnecessary required fields - %v: not found in 'properties'", rs)
	}
	for _, a := range schema.AllOf {
		s, err := NewSchemaRef(a, components, opts)
		if err != nil {
			return nil, fmt.Errorf("new schema ref for allOf: %w", err)
		}
		out.AllOf = append(out.AllOf, s)
	}
	if len(schema.OneOf) > 0 {
		refMapping := map[string]string{}
		for _, a := range schema.OneOf {
			s, err := NewSchemaRef(a, components, opts)
			if err != nil {
				return nil, fmt.Errorf("new schema ref for oneOf: %w", err)
			}
			out.OneOf = append(out.OneOf, s)
			if a.Ref != "" {
				refMapping[a.Ref] = s.Ref().Name
			}
		}
		if schema.Discriminator != nil {
			out.Discriminator.PropertyKey = Just(schema.Discriminator.PropertyName)
			out.Discriminator.Mapping = make([]DiscriminatorMapping, 0, len(out.OneOf))
			mapMapping := map[string]*DiscriminatorMapping{}

			for _, o := range out.OneOf {
				ref := o.Ref()
				if ref == nil {
					continue
				}
				out.Discriminator.Mapping = append(out.Discriminator.Mapping, DiscriminatorMapping{
					Key:    ref.Name,
					Values: []string{ref.Name},
				})
				mapMapping[ref.Name] = &out.Discriminator.Mapping[len(out.Discriminator.Mapping)-1]
			}
			for k, v := range schema.Discriminator.Mapping {
				if m, ok := refMapping[v]; ok {
					mapMapping[m].Values = append(mapMapping[m].Values, k)
				} else {
					if _, ok := mapMapping[v]; !ok {
						out.Discriminator.Mapping = append(out.Discriminator.Mapping, DiscriminatorMapping{
							Key:    v,
							Values: nil,
						})
						mapMapping[v] = &out.Discriminator.Mapping[len(out.Discriminator.Mapping)-1]
					}
					mapMapping[v].Values = append(mapMapping[v].Values, k)
				}
			}
		}
	}

	if schema.AdditionalProperties != nil {
		s, err := NewSchemaRef(schema.AdditionalProperties, components, opts)
		if err != nil {
			return nil, fmt.Errorf("new schema ref for additional properties: %w", err)
		}
		out.AdditionalProperties = Just(s)
	} else if schema.AdditionalPropertiesAllowed != nil && *schema.AdditionalPropertiesAllowed {
		out.AdditionalProperties = Just[Ref[Schema]](&Schema{
			Type:   "object",
			Custom: Just("any"),
		})
	}

	extentions := make(map[string]string)
	for k, v := range schema.ExtensionProps.Extensions {
		if strings.HasPrefix(k, ExtGoagPrefix) {
			if raw, ok := v.(json.RawMessage); ok {
				var s string
				_ = json.Unmarshal(raw, &s)
				extentions[k] = s
			}
		}
	}
	out.Extentions = extentions

	if !opts.IgnoreCustomType {
		if v, ok := extentions[ExtTagGoType]; ok {
			out.Custom = Just(v)
		}
	}

	return &out, nil
}

var _ Ref[Schema] = (*Schema)(nil)

func (s *Schema) Value() *Schema { return s }

type SchemaProperty struct {
	Name     string
	Schema   Ref[Schema]
	Required bool
}

type Discriminator struct {
	PropertyKey Maybe[string]
	Mapping     []DiscriminatorMapping
}

type DiscriminatorMapping struct {
	Key    string
	Values []string
}

func sortedKeys[T any](m map[string]T) (out []string) {
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func httpMethods() []Method {
	return []Method{
		{http.MethodGet, "Get", "http.MethodGet"},
		{http.MethodPost, "Post", "http.MethodPost"},
		{http.MethodPatch, "Patch", "http.MethodPatch"},
		{http.MethodPut, "Put", "http.MethodPut"},
		{http.MethodDelete, "Delete", "http.MethodDelete"},
		{http.MethodConnect, "Connect", "http.MethodConnect"},
		{http.MethodHead, "Head", "http.MethodHead"},
		{http.MethodOptions, "Options", "http.MethodOptions"},
		{http.MethodTrace, "Trace", "http.MethodTrace"},
	}
}

type Method struct {
	HTTP    HTTPMethod
	Title   HTTPMethodTitle
	GoValue string
}

// HTTPMethod - http.MethodGet, http.MethodPost, ...
type HTTPMethod string

// HTTPMethodTitle - Get, Post, ...
type HTTPMethodTitle string
