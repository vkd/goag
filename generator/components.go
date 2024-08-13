package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type Components struct {
	Schemas       []SchemaComponent
	Headers       []HeaderComponent
	RequestBodies []RequestBodyComponent
	Responses     []ResponseComponent

	HasContentJSON bool
}

func NewComponents(spec specification.Components, cfg Config) (zero Components, _ Imports, _ error) {
	var cs Components
	var imports Imports

	schemas := make([]SchemaComponent, len(spec.Schemas.List))
	schemasMap := make(map[*specification.Object[string, specification.Ref[specification.Schema]]]*SchemaComponent, len(spec.Schemas.List))
	for i, c := range spec.Schemas.List {
		schemasMap[c] = &schemas[i]
	}

	cs.Schemas = schemas
	for _, c := range spec.Schemas.List {
		if ref := c.V.Ref(); ref != nil {
			*schemasMap[c] = SchemaComponent{
				Name: c.Name,
				Ref:  Just(schemasMap[ref]),
				Type: NewRef(ref),
			}
			continue
		}

		schema, ims, err := NewSchema(c.V)
		if err != nil {
			return zero, nil, fmt.Errorf("parse schema for %q type: %w", c.Name, err)
		}
		imports = append(imports, ims...)

		sc := SchemaComponent{
			Name:         c.Name,
			Description:  c.V.Value().Description,
			Type:         schema,
			IsMultivalue: schema.IsMultivalue(),
		}

		switch schema := schema.(type) {
		case StructureType:
			sc.IgnoreParseFormat = true
			sc.StructureType = schema
			sc.CustomJSONMarshaler = c.V.Value().AdditionalProperties.Set
		case SliceType:
			sc.RenderFormatStringsMultiline = schema.RenderFormatStringsMultiline
			switch items := schema.Items.(type) {
			case Ref[specification.Schema]:
				switch tp := items.SchemaType.Value(); tp.Type {
				case "object":
					sc.IgnoreParseFormat = true
				case "":
					if len(tp.AllOf) > 0 {
						sc.IgnoreParseFormat = true
					}
				}
			case StructureType:
				sc.IgnoreParseFormat = true
			}
		case CustomType:
			switch schema.Type {
			case "any":
				sc.IgnoreParseFormat = true
			default:
				sc.IsAlias = true
			}
		}

		*schemasMap[c] = sc
	}

	cs.Headers = make([]HeaderComponent, 0, len(spec.Headers.List))
	for _, h := range spec.Headers.List {
		var schema SchemaType
		if ref := h.V.Ref(); ref != nil {
			schema = NewRef(ref)
		} else {
			s, ims, err := NewSchema(h.V.Value().Schema)
			if err != nil {
				return zero, nil, fmt.Errorf("parse header for %q type: %w", h.Name, err)
			}
			imports = append(imports, ims...)
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
					return zero, nil, fmt.Errorf("new schema for %q type, %q content: %w", rb.Name, cnt.Name, err)
				}
				imports = append(imports, ims...)
				cs.RequestBodies = append(cs.RequestBodies, RequestBodyComponent{
					Name:        name,
					Description: rb.V.Value().Description,
					Type:        schema,
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

		response, ims, err := NewResponse(OperationName(r.Name), "", resp, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("new %q response: %w", r.Name, err)
		}
		imports = append(imports, ims...)

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

	return cs, imports, nil
}

func (c Components) Render() (string, error) {
	return ExecuteTemplate("Components", c)
}

func (c Components) LenToRender() int {
	ln := len(c.Schemas) + len(c.Headers) + len(c.RequestBodies) + len(c.Responses)
	return ln
}

type SchemaComponent struct {
	Name              string
	Description       string
	Type              SchemaType
	IgnoreParseFormat bool
	IsMultivalue      bool
	IsAlias           bool

	RenderFormatStringsMultiline func(to, from string) (string, error)

	CustomJSONMarshaler bool
	StructureType       StructureType

	Ref Maybe[*SchemaComponent]
}

func (s SchemaComponent) Base() SchemaType {
	if ref, ok := s.Ref.Get(); ok {
		return ref.Base()
	}
	return s.Type
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
