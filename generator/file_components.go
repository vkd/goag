package generator

import (
	"fmt"
	"strings"

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
	Responses        []ResponseComponent

	HasContentJSON bool
}

func NewComponents(spec specification.Components) (zero Components, _ error) {
	var cs Components
	cs.Schemas = make([]SchemaComponent, 0, len(spec.Schemas.List))
	for _, c := range spec.Schemas.List {
		cs.HasComponent = true
		schema, ims, err := NewSchema(c.V)
		if err != nil {
			return zero, fmt.Errorf("parse schema for %q type: %w", c.Name, err)
		}
		cs.Imports = append(cs.Imports, ims...)

		var structureType StructureType
		var isCustomJSONMarshaler bool
		var ignoreParseFormat bool
		var isAlias bool
		switch schema := schema.(type) {
		case StructureType:
			ignoreParseFormat = true
			structureType = schema
			isCustomJSONMarshaler = c.V.Value().AdditionalProperties.Set
		case SliceType:
			switch items := schema.Items.(type) {
			case Ref[specification.Schema]:
				switch tp := items.SchemaType.Value(); tp.Type {
				case "object":
					ignoreParseFormat = true
				case "":
					if len(tp.AllOf) > 0 {
						ignoreParseFormat = true
					}
				}
			case StructureType:
				ignoreParseFormat = true
			}
		case CustomType:
			switch schema.Type {
			case "any":
				ignoreParseFormat = true
			default:
				isAlias = true
			}
		}
		cs.Schemas = append(cs.Schemas, SchemaComponent{
			Name:              c.Name,
			Description:       c.V.Value().Description,
			Type:              schema,
			IgnoreParseFormat: ignoreParseFormat,
			IsMultivalue:      schema.IsMultivalue(),
			IsAlias:           isAlias,

			CustomJSONMarshaler: isCustomJSONMarshaler,
			StructureType:       structureType,
		})
	}

	cs.Headers = make([]HeaderComponent, 0, len(spec.Headers.List))
	for _, h := range spec.Headers.List {
		// cs.HasComponent = true
		var schema SchemaType
		if ref := h.V.Ref(); ref != nil {
			schema = NewRef(ref)
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
				Type:        NewRef(ref),
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
		// cs.HasComponent = true
		if ref := p.V.Ref(); ref != nil {
			cs.QueryParameters = append(cs.QueryParameters, QueryParameterComponent{
				Name:        p.Name,
				Description: p.V.Value().Description,
				Type:        NewRef(p),
				// IsStringsParser: true,
				IsArray: p.V.Value().Schema.Value().Type == "array",
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
				IsArray:     param.Schema.Value().Type == "array",
			})
		}
	}
	for _, p := range spec.HeaderParameters.List {
		// cs.HasComponent = true
		if ref := p.V.Ref(); ref != nil {
			cs.HeaderParameters = append(cs.HeaderParameters, HeaderParameterComponent{
				Name:        p.Name,
				Description: p.V.Value().Description,
				Type:        NewRef(p),
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
		// cs.HasComponent = true
		if ref := p.V.Ref(); ref != nil {
			cs.PathParameters = append(cs.PathParameters, PathParameterComponent{
				Name:        p.Name,
				Description: p.V.Value().Description,
				Type:        NewRef(p),
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

	// Responses
	for _, r := range spec.Responses.List {
		cs.HasComponent = true

		var alias string
		if ref := r.V.Ref(); ref != nil {
			alias = ref.Name + "Response"
		}

		resp := r.V.Value()

		response, ims, err := NewResponse(OperationName(r.Name), "", resp)
		if err != nil {
			return zero, fmt.Errorf("new %q response: %w", r.Name, err)
		}
		cs.Imports = append(cs.Imports, ims...)

		var ifaces []ResponseUsedIn
		var status string
		for _, usedIn := range resp.UsedIn {
			oName := PublicFieldName(usedIn.Operation.OperationID)
			if oName == "" {
				oName = string(usedIn.Operation.Method)
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

		hr := NewHandlerResponse(response, OperationName(r.Name), status, ifaces...)

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

	return cs, nil
}

func (c Components) Render() (string, error) {
	return ExecuteTemplate("Components", c)
}

type SchemaComponent struct {
	Name              string
	Description       string
	Type              SchemaType
	IgnoreParseFormat bool
	IsMultivalue      bool
	IsAlias           bool

	CustomJSONMarshaler bool
	StructureType       StructureType
}

func (s SchemaComponent) Render() (string, error) {
	if s.IsAlias {
		return ExecuteTemplate("SchemaComponent_Alias", s)
	}
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
	Name            string
	Description     string
	Type            SchemaType
	IsStringsParser bool
	IsArray         bool
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
