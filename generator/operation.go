package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Operation struct {
	*specification.Operation

	Name        OperationName
	Description string
	Summary     string

	// Path specification.PathOld2

	PathBuilder []OperationPathElement

	APIHandlerFieldName string
	HandlerTypeName     string

	RequestTypeName  string
	ResponseTypeName string

	Params OperationParams

	Body struct {
		TypeName Render
	}

	DefaultResponse *Response
	Responses       []*Response
}

func NewOperation(s *specification.Operation, components specification.Components) (zero *Operation, _ Imports, _ error) {
	name := NewOperationName(s)
	o := Operation{
		Operation: s,

		Name:        name,
		Description: s.Description,
		Summary:     s.Summary,

		// Path: s.PathItem.Path,

		APIHandlerFieldName: string(name) + "Handler",
		HandlerTypeName:     string(name) + "HandlerFunc",

		RequestTypeName:  string(name) + "Params",
		ResponseTypeName: string(name) + "Response",
	}

	var imports Imports
	var err error
	o.Params, imports, err = NewOperationParams(s.Parameters)
	if err != nil {
		return zero, nil, fmt.Errorf("new operation params: %w", err)
	}

	var el OperationPathElement
	for _, pd := range s.Path.Dirs {
		if pd.IsVariable {
			if el.Raw != "" {
				el.Raw += "/"
				o.PathBuilder = append(o.PathBuilder, el)
				el = OperationPathElement{}
			}
			pp, ok := o.Params.Path.Get(pd.Param.Value().Name)
			if !ok {
				return zero, nil, fmt.Errorf("unexpected path parameter %q found in path: not found in 'parameters'", pd.Param.Value().Name)
			}
			o.PathBuilder = append(o.PathBuilder, OperationPathElement{
				Param: NewOptional(pp.V),
			})
		} else {
			el.Raw += "/" + pd.V
		}
	}
	if el.Raw != "" {
		o.PathBuilder = append(o.PathBuilder, el)
		el = OperationPathElement{}
	}

	if s.RequestBody.IsSet {
		rBody := s.RequestBody.Value
		if ref := rBody.Ref(); ref != nil && ref.Name != "" {
			o.Body.TypeName = NewRef(ref.Name)
		} else {
			requestBody := rBody.Value()
			jsonContent, ok := requestBody.Content.Get("application/json")
			if ok {
				body, ims, err := NewSchema(jsonContent.V.Schema)
				if err != nil {
					return nil, nil, fmt.Errorf("request body: %w", err)
				}
				imports = append(imports, ims...)
				o.Body.TypeName = body
			}
		}
	}

	for _, r := range s.Responses.List {
		resp, ims, err := NewResponse(name, r.Name, r.V.Value())
		if err != nil {
			return nil, nil, fmt.Errorf("new response for %q status: %w", r.Name, err)
		}
		imports = append(imports, ims...)

		if r.Name == "default" {
			o.DefaultResponse = resp
		} else {
			o.Responses = append(o.Responses, resp)
		}
	}

	return &o, imports, nil
}

type OperationName string

func NewOperationName(s *specification.Operation) OperationName {
	if s.OperationID != "" {
		return OperationName(PublicFieldName(s.OperationID))
	}

	path := s.PathItem.Path

	var out string
	for _, dir := range path.Dirs {
		out += PrefixTitle(dir.Raw)
	}

	var suffix string
	if len(path.Dirs) > 1 && path.Dirs[len(path.Dirs)-1].Raw == "/" {
		suffix = "RT"
	}

	return OperationName(string(s.Method) + out + suffix)
}

type OperationParams struct {
	Query   specification.Map[*QueryParameter]
	Headers specification.Map[*HeaderParameter]
	Path    specification.Map[*PathParameter]
	Cookie  specification.Map[*CookieParameter]
}

func NewOperationParams(params specification.OperationParameters) (zero OperationParams, _ Imports, _ error) {
	var op OperationParams
	var imports Imports

	for _, p := range params.Query.List {
		param, ims, err := NewQueryParameter(p.V.Value())
		if err != nil {
			return zero, nil, fmt.Errorf("new query parameter: %w", err)
		}
		op.Query.Add(p.Name, param)
		imports = append(imports, ims...)
	}

	for _, p := range params.Headers.List {
		param, ims, err := NewHeaderParameter(p.V.Value())
		if err != nil {
			return zero, nil, fmt.Errorf("new header parameter: %w", err)
		}
		op.Headers.Add(p.Name, param)
		imports = append(imports, ims...)
	}

	for _, p := range params.Path.List {
		param, ims, err := NewPathParameter(p.V)
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
	Param Optional[*PathParameter]
}
