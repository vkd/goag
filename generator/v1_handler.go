package generator

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/vkd/goag/specification"
)

type Handler struct {
	s *specification.Operation

	Imports Imports

	Name       OperationName
	Method     specification.HTTPMethodTitle
	HTTPMethod string
	Path       specification.PathOld2

	PathBuilder []PathDir

	HandlerTypeName      string
	HandlerInputTypeName string
	RequestTypeName      string
	RequestVarName       string

	Params struct {
		Query   []*QueryParameter
		Path    PathParameters
		Headers []*HeaderParameter
	}

	IsRequestBody bool

	ResponseTypeName        string
	ResponsePrivateFuncName string

	DefaultResponse Optional[*Response]
	Responses       []*Response
}

func NewHandler(o *specification.Operation) (zero Handler, _ error) {
	h := Handler{s: o}
	h.Method = o.Method
	h.HTTPMethod = o.HTTPMethod
	h.Path = o.PathItem.Path
	h.Name = OperationNameOld(o.OperationID, o.PathItem.Path, h.Method)

	h.HandlerTypeName = string(h.Name) + "HandlerFunc"
	h.HandlerInputTypeName = string(h.Name) + "Parser"
	h.RequestTypeName = string(h.Name) + "Params"
	h.RequestVarName = "request"

	if len(o.Security) > 0 {
		for _, ss := range o.Security {
			if ss.Scheme.Type != specification.SecuritySchemeTypeHTTP {
				continue
			}
			if ss.Scheme.Scheme != "bearer" {
				continue
			}
			p, ims, err := NewHeaderParameter(&specification.HeaderParameter{
				Name:        "Authorization",
				Description: ss.Scheme.BearerFormat,
				Required:    len(o.Security) == 1,
				Schema: &specification.Schema{
					Type:        "string",
					Description: ss.Scheme.BearerFormat,
					Schema: &openapi3.Schema{
						Type:        "string",
						Description: ss.Scheme.BearerFormat,
					},
				},
			})
			if err != nil {
				return zero, fmt.Errorf("schema for header parameter 'Authorization': %w", err)
			}
			h.Params.Headers = append(h.Params.Headers, p)
			h.Imports = append(h.Imports, ims...)
		}
	}

	if o.Operation.RequestBody != nil {
		if _, ok := o.Operation.RequestBody.Value.Content["application/json"]; ok {
			h.IsRequestBody = true
		}
	}

	h.ResponseTypeName = string(h.Name) + "Response"
	h.ResponsePrivateFuncName = PrivateFieldName(string(h.Name))

	if o.DefaultResponse != nil {
		h.DefaultResponse.Set(NewResponse(h.Name, o.DefaultResponse))
	}
	for _, response := range o.Responses {
		h.Responses = append(h.Responses, NewResponse(h.Name, response))
	}

	for _, q := range o.Parameters.Query.List {
		qp := q.V.Value()
		p, ims, err := NewQueryParameter(qp)
		if err != nil {
			return zero, fmt.Errorf("schema for query parameter %q: %w", qp.Name, err)
		}
		h.Params.Query = append(h.Params.Query, p)
		h.Imports = append(h.Imports, ims...)
	}
	for _, q := range o.Parameters.Path.List {
		qp := q.V.Value()
		p, ims, err := NewPathParameter(q.V)
		if err != nil {
			return zero, fmt.Errorf("schema for path parameter %q: %w", qp.Name, err)
		}
		h.Params.Path = append(h.Params.Path, p)
		h.Imports = append(h.Imports, ims...)
	}
	for _, q := range o.Parameters.Headers.List {
		qp := q.V.Value()
		p, ims, err := NewHeaderParameter(qp)
		if err != nil {
			return zero, fmt.Errorf("schema for header parameter %q: %w", qp.Name, err)
		}
		h.Params.Headers = append(h.Params.Headers, p)
		h.Imports = append(h.Imports, ims...)
	}

	{
		var pathBuilder []PathDir
		var last PathDir
		for _, dir := range o.PathItem.Path.Dirs {
			if dir.Raw.IsVariable() {
				last.V += "/"
				pathBuilder = append(pathBuilder, last)
				last = PathDir{}

				pp, err := h.Params.Path.Get(dir.ParamRef.Name)
				if err != nil {
					return zero, fmt.Errorf("path param %q from endpoint: not found in operation path params", dir.ParamRef.Name)
				}
				// varName := h.RequestVarName + ".Path." + pp.FieldName
				// tmp := pp.Type
				pathBuilder = append(pathBuilder, PathDir{Param: pp})
			} else {
				last.V += string(dir.Raw)
			}
		}
		if last.V != "" {
			pathBuilder = append(pathBuilder, last)
		}
		h.PathBuilder = pathBuilder
	}

	return h, nil
}

type PathDir struct {
	V     string
	Param *PathParameter
}

func (p PathDir) FormatTemplater(varName string) Templater {
	if p.Param != nil {
		return TemplaterFunc(func() (string, error) {
			return p.Param.Type.RenderFormat(RenderFunc(func() (string, error) {
				return RawTemplate(varName + ".Path." + p.Param.FieldName).String()
			}))
		})
	} else {
		return TemplaterFunc(func() (string, error) {
			return "\"" + p.V + "\"", nil
		})
	}
}

// type TypeFormat struct {
// 	FieldName string
// }

// func (t TypeFormat) ExecuteArgs(args ...any) (string, error) {
// 	log.Printf("args: %+v", args)
// 	return t.FieldName, nil
// }

// func (t TypeFormat) Execute() (string, error) {
// 	log.Printf("args: no args")
// 	return t.FieldName, nil
// }
// func (t TypeFormat) String() (string, error) {
// 	log.Printf("args: String")
// 	return t.FieldName, nil
// }

type QueryParameter struct {
	s *specification.QueryParameter

	Name      string
	FieldName string
	Required  bool
	Type      Formatter
}

func NewQueryParameter(s *specification.QueryParameter) (zero *QueryParameter, _ Imports, _ error) {
	out := QueryParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	out.Required = s.Required
	var err error
	var ims Imports
	out.Type, ims, err = NewParameterSchema(s.Schema.Value())
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	return &out, ims, nil
}

func (p QueryParameter) ExecuteFormat(from, to string) (string, error) {
	fromTxt := from + ".Query." + p.FieldName
	fromTm := RawTemplate(fromTxt)
	toTm := RawTemplate(to + "[\"" + p.Name + "\"]")

	if sliceType, ok := p.Type.(SliceType); ok {
		if _, ok := sliceType.Items.(StringType); ok {
			return AssignTemplate(fromTm, toTm, false).String()
		} else {
			return TemplateData("SliceToSliceStrings", TData{
				"Len":   RawTemplate("len(" + fromTxt + ")"),
				"From":  fromTm,
				"Items": sliceType.Items,
				"To":    toTm,
			}).String()
		}
	}

	tm := AssignTemplate(
		ToSliceStrings(TemplaterFunc(func() (string, error) {
			return p.Type.RenderFormat(RenderFunc(fromTm.String))
		})),
		toTm,
		false,
	)

	if !p.Required {
		tm = TemplateData("OptionalAssign", TData{
			"From": fromTm,
			"T": AssignTemplate(
				ToSliceStrings(TemplaterFunc(func() (string, error) {
					return p.Type.RenderFormat(RenderFunc(RawTemplate("*" + fromTxt).String))
				})),
				toTm,
				false,
			),
		})
	}

	return tm.String()

	// out := p.Type.FormatAssignTemplater(fromTm, toTm, false)
	// if !p.Required {
	// 	PointerType{}.FormatAssignTemplater(fromTm, toTm, false)
	// }
	// return RawTemplate("/* <not implemented> */")
}

type PathParameters []*PathParameter

func (s PathParameters) Get(name string) (zero *PathParameter, _ error) {
	for _, p := range s {
		if p.s.Name == name {
			return p, nil
		}
	}
	return zero, fmt.Errorf("path parameter %q: not found", name)
}

type PathParameter struct {
	s *specification.PathParameter

	Name          string
	FieldName     string
	FieldTypeName string
	Type          Formatter
}

func NewPathParameter(rs specification.Ref[specification.PathParameter]) (zero *PathParameter, _ Imports, _ error) {
	s := rs.Value()
	out := PathParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if rs.Ref() != nil {
		out.FieldTypeName = rs.Ref().Name
	}
	var ims Imports
	var err error
	out.Type, ims, err = NewParameterSchema(s.Schema.Value())
	if err != nil {
		return nil, nil, fmt.Errorf("schema: %w", err)
	}
	return &out, ims, nil
}

type HeaderParameter struct {
	s *specification.HeaderParameter

	Name          string
	FieldName     string
	FieldTypeName string
	Type          Formatter
	Required      bool
}

func NewHeaderParameter(sr specification.Ref[specification.HeaderParameter]) (zero *HeaderParameter, _ Imports, _ error) {
	s := sr.Value()
	out := HeaderParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if sr.Ref() != nil {
		out.FieldTypeName = sr.Ref().Name
	}
	var ims Imports
	var err error
	out.Type, ims, err = NewParameterSchema(s.Schema.Value())
	if err != nil {
		return zero, nil, fmt.Errorf("schema: %w", err)
	}
	out.Required = s.Required
	return &out, ims, nil
}

type CookieParameter struct {
	specification.CookieParameter
}
