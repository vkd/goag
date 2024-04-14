package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type Handler struct {
	*Operation

	HandlerFuncName        string
	ResponserInterfaceName string
	BasePathPrefix         string

	CanParseError bool

	Parameters  HandlerParameters
	PathParsers []Parser

	DefaultResponse *HandlerResponse
	Responses       []HandlerResponse
}

func NewHandler(o *Operation, basePathPrefix string) (zero *Handler, _ Imports, _ error) {
	out := &Handler{
		Operation: o,

		HandlerFuncName:        string(o.Name) + "HandlerFunc",
		ResponserInterfaceName: PrivateFieldName(string(o.Name)),
		BasePathPrefix:         basePathPrefix,

		CanParseError: len(o.Parameters.Query.List) > 0 || len(o.Parameters.Path.List) > 0 || len(o.Parameters.Headers.List) > 0 || o.Body.TypeName != nil,
	}
	ps, imports, err := NewHandlerParameters(o.Params)
	if err != nil {
		return zero, nil, fmt.Errorf("params: %w", err)
	}
	out.Parameters = ps

	var pathRenders []Parser
	for _, pe := range o.PathBuilder {
		if pe.Param.Set {
			pathRenders = append(pathRenders, PathParserVariable{
				FieldName: pe.Param.Value.FieldName,
				Name:      pe.Param.Value.Name,
				Convert:   pe.Param.Value.Type,
			})
		} else if pe.Raw != "" {
			pathRenders = append(pathRenders, PathParserConstant{
				Prefix:   pe.Raw,
				FullPath: o.Path.Raw,
			})
		}
	}
	out.PathParsers = pathRenders

	for _, sec := range o.Security {
		if sec.Scheme.Type == specification.SecuritySchemeTypeHTTP &&
			sec.Scheme.Scheme == "bearer" &&
			sec.Scheme.BearerFormat == "JWT" {
			out.Parameters.Header = append(out.Parameters.Header, HandlerHeaderParameter{
				HandlerParameter: HandlerParameter{
					FieldName:    "Authorization",
					FieldType:    StringType{},
					FieldComment: sec.Scheme.BearerFormat,
				},
				ParameterName: "Authorization",
				Required:      true,
				Parser:        StringType{},
			})
			out.CanParseError = true
		}
	}

	if o.DefaultResponse != nil {
		resp := NewHandlerResponse(o.DefaultResponse, out)
		out.DefaultResponse = &resp
	}
	for _, r := range o.Responses {
		out.Responses = append(out.Responses, NewHandlerResponse(r, out))
	}

	return out, imports, nil
}

func (h Handler) Render() (string, error) {
	return ExecuteTemplate("Handler", h)
}

type HandlerParameters struct {
	Query  []HandlerQueryParameter
	Path   []HandlerPathParameter
	Header []HandlerHeaderParameter
}

func NewHandlerParameters(p OperationParams) (zero HandlerParameters, _ Imports, _ error) {
	out := HandlerParameters{}
	var imports Imports
	for _, q := range p.Query.List {
		p, ims, err := NewHandlerQueryParameter(q.V)
		if err != nil {
			return zero, nil, fmt.Errorf("query parameter %q: %w", q.Name, err)
		}
		imports = append(imports, ims...)
		out.Query = append(out.Query, p)
	}
	for _, q := range p.Path.List {
		p, ims, err := NewHandlerPathParameter(q.V)
		if err != nil {
			return zero, nil, fmt.Errorf("path parameter %q: %w", q.Name, err)
		}
		imports = append(imports, ims...)
		out.Path = append(out.Path, p)
	}
	for _, q := range p.Headers.List {
		p, ims, err := NewHandlerHeaderParameter(q.V)
		if err != nil {
			return zero, nil, fmt.Errorf("header parameter %q: %w", q.Name, err)
		}
		imports = append(imports, ims...)
		out.Header = append(out.Header, p)
	}
	return out, imports, nil
}

type HandlerParameter struct {
	FieldName    string
	FieldType    Render
	FieldComment string
}

func (h HandlerParameter) Render() (string, error) {
	return ExecuteTemplate("HandlerParameter", h)
}

type HandlerQueryParameter struct {
	HandlerParameter

	ParameterName string
	Required      bool
	Parser        Parser
	IsPointer     bool
}

func NewHandlerQueryParameter(p *QueryParameter) (zero HandlerQueryParameter, _ Imports, _ error) {
	tp := p.Type

	var tpRender Render = tp
	var isPointer bool
	if _, ok := tp.(SliceType); !p.Required && !ok {
		tpRender = NewPointerType(tp)
		isPointer = true
	}

	var parser Parser = tp

	out := HandlerQueryParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    PublicFieldName(p.Name),
			FieldType:    tpRender,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.s.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Required:      p.Required,
		Parser:        parser,
		IsPointer:     isPointer,
	}

	return out, nil, nil
}

func (p HandlerQueryParameter) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch parser := p.Parser.(type) {
	case SliceType:
		return parser.ParseStrings(to, from, isNew, mkErr)
	case Ref[specification.Schema]:
		if parser.SchemaType.Value().Type == "array" {
			return parser.ParseQuery(to, from, isNew, mkErr)
		}
		return parser.ParseSchema(to, from+"[0]", isNew, mkErr)
	case Ref[specification.QueryParameter]:
		return parser.ParseQuery(to, from, isNew, mkErr)
	}
	return p.Parser.ParseString(to, from+"[0]", isNew, mkErr)
}

type PathParserConstant struct {
	SingleValue
	Prefix   string
	FullPath string
}

func (p PathParserConstant) ParseString(_, _ string, _ bool, _ ErrorRender) (string, error) {
	return ExecuteTemplate("PathParserConstant", p)
}

type PathParserVariable struct {
	SingleValue
	FieldName string
	Name      string
	Convert   Parser
}

func (p PathParserVariable) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("PathParserVariable", TData{
		"From":    from,
		"To":      to + p.FieldName,
		"IsNew":   isNew,
		"Error":   wrappedError{mkErr, parseParamError{"path", p.Name}},
		"Convert": p.Convert,
	})
}

type HandlerPathParameter struct {
	HandlerParameter

	ParameterName string
	Parser        Parser
}

func NewHandlerPathParameter(p *PathParameter) (zero HandlerPathParameter, _ Imports, _ error) {
	out := HandlerPathParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    PublicFieldName(p.Name),
			FieldType:    p.Type,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.s.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Parser:        p.Type,
	}

	return out, nil, nil
}

type HandlerHeaderParameter struct {
	HandlerParameter

	ParameterName string
	Required      bool
	Parser        Parser
	IsPointer     bool
}

func NewHandlerHeaderParameter(p *HeaderParameter) (zero HandlerHeaderParameter, _ Imports, _ error) {
	tp := p.Type

	var isPointer bool
	if !p.Required {
		tp = NewPointerType(tp)
		isPointer = true
	}

	fieldName := PublicFieldName(p.Name)

	out := HandlerHeaderParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    fieldName,
			FieldType:    tp,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.s.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Required:      p.Required,
		Parser:        tp,
		IsPointer:     isPointer,
	}

	return out, nil, nil
}

type HandlerResponse struct {
	*Response

	Name string
	// PrivateName string
	HandlerName OperationName

	Status    string
	IsDefault bool

	ResponserInterfaceName string

	IsBody       bool
	BodyTypeName Render
	Body         Render
	BodyRenders  Renders
	ContentType  string
	// Headers     []ResponseHeader

	Struct StructureType

	Args []ResponseArg
}

func NewHandlerResponse(r *Response, h *Handler) HandlerResponse {
	out := HandlerResponse{
		Response: r,

		HandlerName: h.Name,

		Status:    r.StatusCode,
		IsDefault: r.StatusCode == "default",
	}

	out.Name = string(h.Name) + "Response" + strings.Title(r.StatusCode)
	if r.ContentJSON.Set {
		out.Name += "JSON"
		out.ContentType = "application/json"
	}

	if out.IsDefault {
		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name: "Code",
			Type: IntType{},
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Code",
			ArgName:   "code",
			Type:      IntType{},
		})
	}

	if r.ContentJSON.Set {
		out.IsBody = true
		switch contentType := r.ContentJSON.Value.Type.(type) {
		case Ref[specification.Schema], SliceType, CustomType:
			out.BodyTypeName = contentType
		default:
			bodyStructName := out.Name + "Body"
			out.BodyTypeName = NewRef(&specification.Object[string, specification.Ref[specification.Schema]]{
				Name: bodyStructName,
				V:    r.ContentJSON.Value.Spec,
			})
			bodyType := r.ContentJSON.Value
			out.Body = bodyType.Type

			if bodyType.Spec.Value().AdditionalProperties.Set {
				out.BodyRenders = Renders{
					StringRender("var _ json.Marshaler = (*" + bodyStructName + ")" + "(nil)"),
					RenderFunc(func() (string, error) {
						out := `func (b ` + bodyStructName + `) MarshalJSON() ([]byte, error) {
							m := make(map[string]interface{})
							for k, v := range b.AdditionalProperties {
								m[k] = v
							}
							`
						if st, ok := bodyType.Type.(StructureType); ok {
							for _, f := range st.Fields {
								if t, ok := f.GetTag("json"); ok && len(t.Values) > 0 && t.Values[0] != "-" {
									out += "m[\"" + t.Values[0] + "\"] = b." + f.Name + "\n"
								}
							}
						}
						out += `return json.Marshal(m)

						}`
						return out, nil
					}),
				}
			}
		}

		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name: "Body",
			Type: out.BodyTypeName,
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Body",
			ArgName:   "body",
			Type:      out.BodyTypeName,
		})
	}

	var headersStruct StructureType
	for _, h := range r.Headers {
		headersStruct.Fields = append(headersStruct.Fields, StructureField{
			Name: h.FieldName,
			Type: h.Schema,
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: h.FieldName,
			ArgName:   PrivateFieldName(h.FieldName),
			IsHeader:  true,
			Type:      h.Schema,
		})
	}
	if len(headersStruct.Fields) > 0 {
		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name: "Headers",
			Type: headersStruct,
		})
	}

	out.ResponserInterfaceName = h.ResponserInterfaceName

	return out
}

func (h HandlerResponse) Render() (string, error) {
	return ExecuteTemplate("HandlerResponse", h)
}

type ResponseArg struct {
	FieldName string
	ArgName   string
	IsHeader  bool
	Type      Render
}
