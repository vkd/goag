package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Handler struct {
	s *specification.Operation

	Name       string
	Method     string
	HTTPMethod string
	Path       specification.Path

	PathBuilder []PathDir

	HandlerTypeName      string
	HandlerInputTypeName string
	RequestTypeName      string
	RequestVarName       string

	Params struct {
		Query   []QueryParameter
		Path    PathParameters
		Headers []HeaderParameter
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
	h.Name = OperationName(o.OperationID, o.PathItem.Path, h.Method)

	h.HandlerTypeName = h.Name + "HandlerFunc"
	h.HandlerInputTypeName = h.Name + "Parser"
	h.RequestTypeName = h.Name + "Params"
	h.RequestVarName = "request"

	if o.Operation.RequestBody != nil {
		if _, ok := o.Operation.RequestBody.Value.Content["application/json"]; ok {
			h.IsRequestBody = true
		}
	}

	h.ResponseTypeName = h.Name + "Response"
	h.ResponsePrivateFuncName = PrivateFieldName(h.Name)

	if o.DefaultResponse != nil {
		h.DefaultResponse.Set(NewResponse(h.Name, o.DefaultResponse))
	}
	for _, response := range o.Responses {
		h.Responses = append(h.Responses, NewResponse(h.Name, response))
	}

	for _, qp := range o.Parameters.Query {
		p, err := NewQueryParameter(qp)
		if err != nil {
			return zero, fmt.Errorf("schema for query parameter %q: %w", qp.Name, err)
		}
		h.Params.Query = append(h.Params.Query, p)
	}
	for _, qp := range o.Parameters.Path {
		p, err := NewPathParameter(qp)
		if err != nil {
			return zero, fmt.Errorf("schema for path parameter %q: %w", qp.Name, err)
		}
		h.Params.Path = append(h.Params.Path, p)
	}
	for _, qp := range o.Parameters.Headers {
		p, err := NewHeaderParameter(qp)
		if err != nil {
			return zero, fmt.Errorf("schema for header parameter %q: %w", qp.Name, err)
		}
		h.Params.Headers = append(h.Params.Headers, p)
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
				pathBuilder = append(pathBuilder, PathDir{Param: &pp})
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
	s specification.QueryParameter

	Name      string
	FieldName string
	Required  bool
	Type      Formatter
}

func NewQueryParameter(s specification.QueryParameter) (zero QueryParameter, _ error) {
	out := QueryParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	out.Required = s.Required
	var err error
	out.Type, err = NewParameterSchema(s.Schema)
	if err != nil {
		return zero, fmt.Errorf("schema: %w", err)
	}
	return out, nil
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

type PathParameters []PathParameter

func (s PathParameters) Get(name string) (zero PathParameter, _ error) {
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

func NewPathParameter(s *specification.PathParameter) (zero PathParameter, _ error) {
	out := PathParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if s.RefName != "" {
		out.FieldTypeName = s.RefName
	}
	var err error
	out.Type, err = NewParameterSchema(s.Schema)
	if err != nil {
		return zero, fmt.Errorf("schema: %w", err)
	}
	return out, nil
}

type HeaderParameter struct {
	s *specification.HeaderParameter

	Name          string
	FieldName     string
	FieldTypeName string
	Type          Formatter
	Required      bool
}

func NewHeaderParameter(s *specification.HeaderParameter) (zero HeaderParameter, _ error) {
	out := HeaderParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if s.RefName != "" {
		out.FieldTypeName = s.RefName
	}
	var err error
	out.Type, err = NewParameterSchema(s.Schema)
	if err != nil {
		return zero, fmt.Errorf("schema: %w", err)
	}
	out.Required = s.Required
	return out, nil
}
