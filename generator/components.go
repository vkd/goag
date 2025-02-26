package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type Components struct {
	Schemas       []*SchemaComponent
	Headers       []HeaderComponent
	RequestBodies []RequestBodyComponent
	Responses     []ResponseComponent

	HasContentJSON bool
}

func NewComponents(spec specification.Components, cfg Config) (zero Components, _ Imports, _ error) {
	var components Components
	cs := &components

	var imports Imports

	cs.Schemas = make([]*SchemaComponent, len(spec.Schemas.List))
	cacheSchemas := make(map[specification.Ref[specification.Schema]]*SchemaComponent, len(spec.Schemas.List))
	for i, c := range spec.Schemas.List {
		cs.Schemas[i] = &SchemaComponent{}
		cs.Schemas[i].Name = c.Name
		cacheSchemas[c.V] = cs.Schemas[i]
	}
	for _, c := range spec.Schemas.List {
		schema, ims, err := NewSchema(c.V, NamedComponenter{cs, c.Name}, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("new schema component: %w", err)
		}
		imports = append(imports, ims...)

		s := NewSchemaComponent(c.Name, schema, cs, cfg)

		sc := cacheSchemas[c.V]
		*sc = s
	}

	cs.Headers = make([]HeaderComponent, 0, len(spec.Headers.List))
	for _, h := range spec.Headers.List {
		s, ims, err := NewSchema(h.V.Value().Schema, cs, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("parse header for %q type: %w", h.Name, err)
		}
		imports = append(imports, ims...)

		cs.Headers = append(cs.Headers, HeaderComponent{
			Name:           h.Name,
			Description:    h.V.Value().Description,
			Type:           s,
			GoTypeRenderFn: s.RenderGoType,
		})
	}

	cs.RequestBodies = make([]RequestBodyComponent, 0, len(spec.RequestBodies.List))
	for _, rb := range spec.RequestBodies.List {
		if ref := rb.V.Ref(); ref != nil {
			cs.RequestBodies = append(cs.RequestBodies, RequestBodyComponent{
				Name:        rb.Name + "JSON",
				Description: rb.V.Value().Description,
				GoTypeFn:    StringRender(ref.Name + "JSON").Render,
			})
		} else {
			for _, cnt := range rb.V.Value().Content.List {
				name := rb.Name
				switch cnt.Name {
				case "application/json":
					name += "JSON"
				default:
					name += PublicFieldName(cnt.Name)
				}
				schema, ims, err := NewSchema(cnt.V.Schema, cs, cfg)
				if err != nil {
					return zero, nil, fmt.Errorf("new schema for %q type, %q content: %w", rb.Name, cnt.Name, err)
				}
				imports = append(imports, ims...)
				cs.RequestBodies = append(cs.RequestBodies, RequestBodyComponent{
					Name:        name,
					Description: rb.V.Value().Description,
					GoTypeFn:    schema.RenderGoType,
				})
			}
		}
	}

	// Responses
	for _, r := range spec.Responses.List {
		var alias string
		if ref := r.V.Ref(); ref != nil {
			alias = ref.Name + "Response"
		}

		resp := r.V.Value()

		response, ims, err := NewResponse(OperationName(r.Name), "", resp, cs, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("new %q response: %w", r.Name, err)
		}
		imports = append(imports, ims...)

		var ifaces []ResponseUsedIn
		var status string
		for _, usedIn := range resp.UsedIn {
			oName := PublicFieldName(usedIn.Operation.OperationID)
			if oName == "" {
				oName = string(usedIn.Operation.Method.Title)
				raw := usedIn.Operation.PathRaw
				for _, ss := range strings.Split(raw, "/")[1:] {
					if strings.HasPrefix(ss, "{") && strings.HasSuffix(ss, "}") {
						oName += Title(ss[1 : len(ss)-1])
					} else {
						oName += Title(ss)
					}
				}
				if raw != "" && strings.HasSuffix(raw, "/") {
					oName += "RT"
				}
			}
			switch usedIn.Status {
			case "default":
				status = usedIn.Status
			}

			ifaces = append(ifaces, ResponseUsedIn{
				OperationName: OperationName(oName),
				Status:        usedIn.Status,
			})
		}

		hr := NewHandlerResponse(response, OperationName(r.Name), status, cs, cfg, ifaces...)

		if hr.ContentJSON.IsSet {
			cs.HasContentJSON = true
		}

		cs.Responses = append(cs.Responses, ResponseComponent{
			Name:        r.Name + "Response",
			Description: r.V.Value().Description,
			Alias:       alias,
			IsComponent: true,

			HandlerResponse: hr,
		})
	}

	return components, imports, nil
}

func (c Components) Render() (string, error) {
	return ExecuteTemplate("Components", c)
}

func (c Components) LenToRender() int {
	ln := len(c.Schemas) + len(c.Headers) + len(c.RequestBodies) + len(c.Responses)
	return ln
}

func (c Components) GetSchema(key string) (*SchemaComponent, bool) {
	for i, s := range c.Schemas {
		if s.Name == key {
			return c.Schemas[i], true
		}
	}
	return nil, false
}

func (c *Components) AddSchema(name string, s Schema, cfg Config) *SchemaComponent {
	sc := NewSchemaComponent(name, s, c, cfg)
	c.Schemas = append(c.Schemas, &sc)
	return c.Schemas[len(c.Schemas)-1]
}

type NamedComponenter struct {
	Componenter
	Name string
}

func (n NamedComponenter) AddSchema(name string, s Schema, cfg Config) *SchemaComponent {
	return n.Componenter.AddSchema(n.Name+PublicFieldName(name), s, cfg)
}

type SchemaComponent struct {
	Name   string
	Schema Schema

	Description        string
	IgnoreParseFormat  bool
	IsAlias            bool
	WriteJSONFunc      bool
	WriteJSONFuncArray bool
	WriteJSONFuncOneOf bool

	StructureType  StructureType
	SliceType      SliceType
	OneOfStructure OneOfStructure

	GoTypeFn GoTypeRenderFunc
}

func NewSchemaComponent(name string, schema Schema, cs Componenter, cfg Config) SchemaComponent {
	sc := SchemaComponent{
		Name:   name,
		Schema: schema,

		Description: schema.Description,

		GoTypeFn: schema.RenderGoType,
	}

	switch schema := schema.Type.(type) {
	case AnyType:
		sc.IgnoreParseFormat = true
		sc.IsAlias = true
	case StructureType:
		sc.IgnoreParseFormat = true
		sc.StructureType = schema
		sc.WriteJSONFunc = true
	case SliceType:
		sc.WriteJSONFuncArray = true
		sc.SliceType = schema

		switch schema.Items.Base().Type.(type) {
		case StructureType:
			sc.IgnoreParseFormat = true
		}
	case OneOfStructure:
		sc.IgnoreParseFormat = true
		sc.WriteJSONFuncOneOf = true
		sc.StructureType = schema.Struct
		sc.OneOfStructure = schema
	}

	return sc
}

func (s SchemaComponent) Render() (string, error) {
	if s.IsAlias {
		return ExecuteTemplate("SchemaComponent_Alias", s)
	}
	return ExecuteTemplate("SchemaComponent", s)
}

type HeaderComponent struct {
	Name           string
	Description    string
	Type           SchemaType
	GoTypeRenderFn GoTypeRenderFunc
}

func (s HeaderComponent) Render() (string, error) {
	return ExecuteTemplate("HeaderComponent", s)
}

type RequestBodyComponent struct {
	Name        string
	Description string
	GoTypeFn    GoTypeRenderFunc
}

func (s RequestBodyComponent) Render() (string, error) {
	return ExecuteTemplate("RequestBodyComponent", s)
}

type ResponseComponent struct {
	Name        string
	Description string
	Alias       string
	IsComponent bool

	HandlerResponse
}

func (c ResponseComponent) Render() (string, error) {
	if c.Alias != "" {
		return ExecuteTemplate("ResponseComponentAlias", c)
	}
	return ExecuteTemplate("ResponseComponent", c)
}
