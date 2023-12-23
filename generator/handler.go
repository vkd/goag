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

	h.IsRequestBody = o.Operation.RequestBody != nil

	h.ResponseTypeName = h.Name + "Response"
	h.ResponsePrivateFuncName = PrivateFieldName(h.Name)

	if o.DefaultResponse != nil {
		h.DefaultResponse.Set(NewResponse(h.Name, o.DefaultResponse))
	}
	for _, response := range o.Responses {
		h.Responses = append(h.Responses, NewResponse(h.Name, response))
	}

	for _, qp := range o.Parameters.Query {
		h.Params.Query = append(h.Params.Query, NewQueryParameter(qp))
	}
	for _, qp := range o.Parameters.Path {
		h.Params.Path = append(h.Params.Path, NewPathParameter(qp))
	}
	for _, qp := range o.Parameters.Headers {
		h.Params.Headers = append(h.Params.Headers, NewHeaderParameter(qp))
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
		return p.Param.Type.TemplateToString(RawTemplate(varName + ".Path." + p.Param.FieldName))
		// p.Param.Type
		// return fieldRef
		// panic("not implemented")
	} else {
		return StringConst(p.V)
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
	Type      Schema
}

func NewQueryParameter(s specification.QueryParameter) QueryParameter {
	out := QueryParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	out.Required = s.Required
	out.Type = NewSchema(s.Schema)
	return out
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
		ToSliceStrings(
			p.Type.TemplateToString(fromTm),
		),
		toTm,
		false,
	)

	if !p.Required {
		tm = TemplateData("OptionalAssign", TData{
			"From": fromTm,
			"T": AssignTemplate(
				ToSliceStrings(
					p.Type.TemplateToString(RawTemplate("*"+fromTxt)),
				),
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
	Type          Schema
}

func NewPathParameter(s *specification.PathParameter) PathParameter {
	out := PathParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if s.RefName != "" {
		out.FieldTypeName = s.RefName
	}
	out.Type = NewSchema(s.Schema)
	return out
}

type HeaderParameter struct {
	s *specification.HeaderParameter

	Name          string
	FieldName     string
	FieldTypeName string
	Type          Schema
	Required      bool
}

func NewHeaderParameter(s *specification.HeaderParameter) HeaderParameter {
	out := HeaderParameter{s: s}
	out.Name = s.Name
	out.FieldName = PublicFieldName(s.Name)
	if s.RefName != "" {
		out.FieldTypeName = s.RefName
	}
	out.Type = NewSchema(s.Schema)
	out.Required = s.Required
	return out
}
