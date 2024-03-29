package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Components struct {
	HasComponent bool

	Imports Imports

	Schemas          []SchemaComponent
	Headers          []HeaderComponent
	RequestBodies    []RequestBodyComponent
	QueryParameters  []QueryParameterComponent
	HeaderParameters []HeaderParameterComponent
	PathParameters   []PathParameterComponent
}

func NewComponents(spec specification.Components) (zero Components, _ error) {
	var cs Components
	cs.Schemas = make([]SchemaComponent, 0, len(spec.Schemas.List))
	for _, c := range spec.Schemas.List {
		cs.HasComponent = true
		var schema SchemaType
		if ref := c.V.Ref(); ref != nil {
			schema = NewRef(ref.Name)
		} else {
			s, ims, err := NewSchema(c.V.Value())
			if err != nil {
				return zero, fmt.Errorf("parse schema for %q type: %w", c.Name, err)
			}
			cs.Imports = append(cs.Imports, ims...)
			schema = s
		}
		var ignoreParseFormat bool
		switch schema := schema.(type) {
		case StructureType:
			ignoreParseFormat = true
		case SliceType:
			ignoreParseFormat = true
		case CustomType:
			switch schema.Type {
			case "any", "json.RawMessage":
				ignoreParseFormat = true
			}
		}
		cs.Schemas = append(cs.Schemas, SchemaComponent{
			Name:              c.Name,
			Type:              schema,
			IgnoreParseFormat: ignoreParseFormat,
		})
	}

	cs.Headers = make([]HeaderComponent, 0, len(spec.Headers.List))
	for _, h := range spec.Headers.List {
		cs.HasComponent = true
		var schema SchemaType
		if ref := h.V.Ref(); ref != nil {
			schema = NewRef(ref.Name)
		} else {
			s, ims, err := NewSchema(h.V.Value().Schema)
			if err != nil {
				return zero, fmt.Errorf("parse header for %q type: %w", h.Name, err)
			}
			cs.Imports = append(cs.Imports, ims...)
			schema = s
		}

		cs.Headers = append(cs.Headers, HeaderComponent{
			Name:        h.Name,
			Description: h.V.Value().Description,
			Type:        schema,
		})
	}

	cs.RequestBodies = make([]RequestBodyComponent, 0, len(spec.RequestBodies.List))
	for _, rb := range spec.RequestBodies.List {
		cs.HasComponent = true
		if ref := rb.V.Ref(); ref != nil {
			cs.RequestBodies = append(cs.RequestBodies, RequestBodyComponent{
				Name:        rb.Name,
				Description: rb.V.Value().Description,
				Type:        NewRef(ref.Name),
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
				schema, ims, err := NewSchema(cnt.V.Schema)
				if err != nil {
					return zero, fmt.Errorf("new schema for %q type, %q content: %w", rb.Name, cnt.Name, err)
				}
				cs.Imports = append(cs.Imports, ims...)
				cs.RequestBodies = append(cs.RequestBodies, RequestBodyComponent{
					Name:        name,
					Description: rb.V.Value().Description,
					Type:        schema,
				})
			}
		}
	}

	for _, p := range spec.QueryParameters.List {
		cs.HasComponent = true
		if ref := p.V.Ref(); ref != nil {
			cs.QueryParameters = append(cs.QueryParameters, QueryParameterComponent{
				Name:        p.Name,
				Description: p.V.Value().Description,
				Type:        NewRef(p.Name),
			})
		} else {
			param := p.V.Value()

			tp, ims, err := NewSchema(param.Schema)
			if err != nil {
				return zero, fmt.Errorf("new schema for %q query param: %w", p.Name, err)
			}
			cs.Imports = append(cs.Imports, ims...)
			cs.QueryParameters = append(cs.QueryParameters, QueryParameterComponent{
				Name:        p.Name,
				Description: param.Description,
				Type:        tp,
			})
		}
	}
	for _, p := range spec.HeaderParameters.List {
		cs.HasComponent = true
		if ref := p.V.Ref(); ref != nil {
			cs.HeaderParameters = append(cs.HeaderParameters, HeaderParameterComponent{
				Name:        p.Name,
				Description: p.V.Value().Description,
				Type:        NewRef(p.Name),
			})
		} else {
			param := p.V.Value()

			tp, ims, err := NewSchema(param.Schema)
			if err != nil {
				return zero, fmt.Errorf("new schema for %q header param: %w", p.Name, err)
			}
			cs.Imports = append(cs.Imports, ims...)
			cs.HeaderParameters = append(cs.HeaderParameters, HeaderParameterComponent{
				Name:        p.Name,
				Description: param.Description,
				Type:        tp,
			})
		}
	}
	for _, p := range spec.PathParameters.List {
		cs.HasComponent = true
		if ref := p.V.Ref(); ref != nil {
			cs.PathParameters = append(cs.PathParameters, PathParameterComponent{
				Name:        p.Name,
				Description: p.V.Value().Description,
				Type:        NewRef(p.Name),
			})
		} else {
			param := p.V.Value()

			tp, ims, err := NewSchema(param.Schema)
			if err != nil {
				return zero, fmt.Errorf("new schema for %q path param: %w", p.Name, err)
			}
			cs.Imports = append(cs.Imports, ims...)
			cs.PathParameters = append(cs.PathParameters, PathParameterComponent{
				Name:        p.Name,
				Description: param.Description,
				Type:        tp,
			})
		}
	}

	return cs, nil
}

func (c Components) Render() (string, error) {
	return ExecuteTemplate("Components", c)
}

type SchemaComponent struct {
	Name              string
	Type              SchemaType
	IgnoreParseFormat bool
}

func (s SchemaComponent) Render() (string, error) {
	return ExecuteTemplate("SchemaComponent", s)
}

type HeaderComponent struct {
	Name        string
	Description string
	Type        Render
}

func (s HeaderComponent) Render() (string, error) {
	return ExecuteTemplate("HeaderComponent", s)
}

type RequestBodyComponent struct {
	Name        string
	Description string
	Type        Render
}

func (s RequestBodyComponent) Render() (string, error) {
	return ExecuteTemplate("RequestBodyComponent", s)
}

type QueryParameterComponent struct {
	Name        string
	Description string
	Type        SchemaType
}

func (c QueryParameterComponent) Render() (string, error) {
	return ExecuteTemplate("QueryParameterComponent", c)
}

type HeaderParameterComponent struct {
	Name        string
	Description string
	Type        SchemaType
}

func (c HeaderParameterComponent) Render() (string, error) {
	return ExecuteTemplate("HeaderParameterComponent", c)
}

type PathParameterComponent struct {
	Name        string
	Description string
	Type        SchemaType
}

func (c PathParameterComponent) Render() (string, error) {
	return ExecuteTemplate("PathParameterComponent", c)
}
