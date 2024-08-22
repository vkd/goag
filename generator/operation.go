package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type PathItem struct {
	*specification.PathItem
	Operations []*Operation
}

type Operation struct {
	*specification.Operation

	Name        OperationName
	Description string
	Summary     string

	Path Path

	PathBuilder []OperationPathElement

	APIHandlerFieldName string
	HandlerTypeName     string

	RequestTypeName  string
	ResponseTypeName string

	Params OperationParams

	Body struct {
		TypeName Render
		Type     Maybe[StructureType]
	}

	DefaultResponse *ResponseCode
	Responses       []*ResponseCode
}

func NewOperation(s *specification.Operation, componenets Components, cfg Config) (zero *Operation, _ Imports, _ error) {
	path, err := NewPath(s.PathRaw)
	if err != nil {
		return zero, nil, fmt.Errorf("parse raw path: %w", err)
	}
	name := NewOperationName(s.Method, s, path)

	o := Operation{
		Operation: s,

		Name:        name,
		Description: s.Description,
		Summary:     s.Summary,

		Path: path,

		APIHandlerFieldName: string(name) + "Handler",
		HandlerTypeName:     string(name) + "HandlerFunc",

		RequestTypeName:  string(name) + "Params",
		ResponseTypeName: string(name) + "Response",
	}

	var imports Imports
	o.Params, imports, err = NewOperationParams(s.Parameters, componenets)
	if err != nil {
		return zero, nil, fmt.Errorf("new operation params: %w", err)
	}

	for _, pathParam := range o.Params.Path.List {
		param, ok := o.Path.Parameters.Get(pathParam.Name)
		if !ok {
			return zero, nil, fmt.Errorf("%q path parameter: not found in %q endpoint", pathParam.Name, o.Path.Raw)
		}
		param.V.Param = pathParam.V
	}
	for _, pp := range o.Path.Parameters.List {
		if pp.V.IsVariable && pp.V.Param == nil {
			return zero, nil, fmt.Errorf("%q endpoint: %q param is not defined", o.Path.Raw, pp.V.V)
		}
	}

	var el OperationPathElement
	for _, pd := range o.Path.Dirs {
		if pd.IsVariable {
			el.Raw += "/"
			o.PathBuilder = append(o.PathBuilder, el)
			el = OperationPathElement{}

			o.PathBuilder = append(o.PathBuilder, OperationPathElement{
				Param: Just(pd.Param),
			})
		} else {
			el.Raw += "/" + pd.V
		}
	}
	if el.Raw != "" {
		o.PathBuilder = append(o.PathBuilder, el)
		el = OperationPathElement{}
	}

	if rBody, ok := s.RequestBody.Get(); ok {
		if ref := rBody.Ref(); ref != nil && ref.Name != "" {
			if _, ok := ref.V.Value().Content.Get("application/json"); ok {
				o.Body.TypeName = StringRender(ref.Name + "JSON")
			}
		} else {
			requestBody := rBody.Value()
			jsonContent, ok := requestBody.Content.Get("application/json")
			if ok {
				body, ims, err := NewSchema(jsonContent.V.Schema, componenets)
				if err != nil {
					return nil, nil, fmt.Errorf("request body: %w", err)
				}
				imports = append(imports, ims...)

				if body.Custom.IsSet {
					o.Body.TypeName = body
				} else {
					switch st := body.Type.(type) {
					case StructureType:
						o.Body.Type = Just(st)
					default:
						o.Body.TypeName = body
					}
				}
			}
		}
	}

	for _, r := range s.Responses.List {
		resp, ims, err := NewResponse(name, r.Name, r.V.Value(), componenets, cfg)
		if err != nil {
			return nil, nil, fmt.Errorf("new response for %q status: %w", r.Name, err)
		}
		imports = append(imports, ims...)

		switch r.Name {
		case "default":
			o.DefaultResponse = &ResponseCode{
				Response:     resp,
				StatusCode:   r.Name,
				ComponentRef: r.V.Ref(),
			}
		default:
			o.Responses = append(o.Responses, &ResponseCode{
				Response:     resp,
				StatusCode:   r.Name,
				ComponentRef: r.V.Ref(),
			})
		}
	}

	if len(o.Security) > 0 {
		for _, sr := range o.Security {
			if sr.Scheme.Type != specification.SecuritySchemeTypeHTTP {
				continue
			}
			if sr.Scheme.Scheme != "bearer" {
				continue
			}
			o.Params.Headers.Add("Authorization", &HeaderParameter{
				Name:        "Authorization",
				FieldName:   "Authorization",
				Description: sr.Scheme.BearerFormat,
				Required:    len(o.Security) == 1,
				Type:        NewSchemaWithType(NewPrimitive(StringType{})),
				Schema:      NewSchemaWithType(NewPrimitive(StringType{})),
			})
		}
	}

	return &o, imports, nil
}

type OperationName string

func NewOperationName(method specification.HTTPMethodTitle, s *specification.Operation, path Path) OperationName {
	if s.OperationID != "" {
		return OperationName(PublicFieldName(s.OperationID))
	}

	var out string
	for _, dir := range path.Dirs {
		out += Title(dir.V)
	}

	var suffix string
	if len(path.Dirs) > 1 && path.Dirs[len(path.Dirs)-1].V == "" {
		suffix = "RT"
	}

	return OperationName(string(method) + out + suffix)
}

type OperationParams struct {
	Query   specification.Map[*QueryParameter]
	Headers specification.Map[*HeaderParameter]
	Path    specification.Map[*PathParameter]
	Cookie  specification.Map[*CookieParameter]
}

func NewOperationParams(params specification.OperationParameters, components Components) (zero OperationParams, _ Imports, _ error) {
	var op OperationParams
	var imports Imports

	for _, p := range params.Query.List {
		param, ims, err := NewQueryParameter(p.V, components)
		if err != nil {
			return zero, nil, fmt.Errorf("new query parameter: %w", err)
		}
		imports = append(imports, ims...)
		op.Query.Add(p.Name, param)
	}

	for _, p := range params.Headers.List {
		param, ims, err := NewHeaderParameter(p.V, components)
		if err != nil {
			return zero, nil, fmt.Errorf("new header parameter: %w", err)
		}
		op.Headers.Add(p.Name, param)
		imports = append(imports, ims...)
	}

	for _, p := range params.Path.List {
		param, ims, err := NewPathParameter(p.V, components)
		if err != nil {
			return zero, nil, fmt.Errorf("new path parameter: %w", err)
		}
		op.Path.Add(p.Name, param)
		imports = append(imports, ims...)
	}

	return op, imports, nil
}

type OperationPathElement struct {
	Raw   string
	Param Maybe[*PathParameter]
}

type ResponseCode struct {
	*Response
	StatusCode   string
	ComponentRef *specification.Object[string, specification.Ref[specification.Response]]
}
