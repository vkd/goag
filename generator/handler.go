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
		if pe.Param.OK {
			rendParam := pe.Param.Value
			pathRenders = append(pathRenders, PathParserVariable{
				Variable: rendParam.FieldName,
				Error:    PathParseError(rendParam.Name),
				Convert: ParserFunc(func(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
					rs := Renders{
						RenderFunc(func() (string, error) { return rendParam.Type.ParseString("v", from, true, mkErr) }),
						RenderFunc(func() (string, error) {
							return to + " = v", nil
						}),
					}
					return rs.Render()
				}),
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
				Parser: ParserFunc(func(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
					return Renders{
						StringRender("v := " + from),
						RenderFunc(func() (string, error) {
							return StringType{}.ParseString(to+"Authorization", "v", isNew, HeaderParseError("Authorization"))
						}),
					}.Render()
				}),
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
}

func NewHandlerQueryParameter(p *QueryParameter) (zero HandlerQueryParameter, _ Imports, _ error) {
	tp, imports, err := NewSchema(p.s.Schema)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}

	var tpRender Render = tp
	if _, ok := tp.(SliceType); !p.Required && !ok {
		tpRender = NewPointerType(tp)
	}

	var parser Parser = tp
	if _, ok := tp.(SliceType); !ok {
		parser = ParserFunc(func(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
			return Renders{
				RenderFunc(func() (string, error) {
					return tp.ParseString("v", from, true, mkErr)
				}),
				RenderFunc(func() (string, error) {
					from := "v"
					if !p.Required {
						from = "&v"
					}
					return Assign(to, from, isNew), nil
				}),
			}.Render()
		})
	}

	out := HandlerQueryParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    PublicFieldName(p.Name),
			FieldType:    tpRender,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.s.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Required:      p.Required,
		Parser:        parser,
	}

	return out, imports, nil
}

func (p HandlerQueryParameter) ParseStrings(to, from string, isNew bool, _ ErrorRender) (string, error) {
	mkErr := QueryParseError(p.ParameterName)

	switch parser := p.Parser.(type) {
	case SliceType:
		return parser.ParseStrings(to, from, isNew, mkErr)
	}
	return p.Parser.ParseString(to, from+"[0]", isNew, mkErr)
}

type PathParserConstant struct {
	Prefix   string
	FullPath string
}

func (p PathParserConstant) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("PathParserConstant", p)
}

type PathParserVariable struct {
	Variable string
	Error    ErrorRender
	Convert  Parser
}

func (p PathParserVariable) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("PathParserVariable", p)
}

type HandlerPathParameter struct {
	HandlerParameter

	ParameterName string
	Parser        Parser
}

func NewHandlerPathParameter(p *PathParameter) (zero HandlerPathParameter, _ Imports, _ error) {
	tp, imports, err := NewSchema(p.s.Schema)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}

	out := HandlerPathParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    PublicFieldName(p.Name),
			FieldType:    tp,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.s.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Parser:        tp,
	}

	return out, imports, nil
}

func (p HandlerPathParameter) ParseString(from, to string) (string, error) {
	return p.Parser.ParseString(to, from, true, PathParseError(p.ParameterName))
}

type HandlerHeaderParameter struct {
	HandlerParameter

	ParameterName string
	Required      bool
	Parser        Parser
}

func NewHandlerHeaderParameter(p *HeaderParameter) (zero HandlerHeaderParameter, _ Imports, _ error) {
	tp, imports, err := NewSchema(p.s.Schema)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}

	if !p.Required {
		tp = NewPointerType(tp)
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
		Parser: ParserFunc(func(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
			if _, ok := tp.(StringType); ok {
				return Renders{
					StringRender("v := " + from),
					RenderFunc(func() (string, error) { return tp.ParseString(to+fieldName, "v", isNew, HeaderParseError(p.Name)) }),
				}.Render()
			}
			return tp.ParseString(to+fieldName, from, isNew, HeaderParseError(p.Name))
		}),
	}

	return out, imports, nil
}

func (p HandlerHeaderParameter) ParseString(from, to string) (string, error) {
	mkErr := HeaderParseError(p.ParameterName)
	return p.Parser.ParseString(to, from, true, mkErr)
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
	if r.ContentJSON.OK {
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

	if r.ContentJSON.OK {
		out.IsBody = true
		switch contentType := r.ContentJSON.Value.Type.(type) {
		case Ref, SliceType, CustomType:
			out.BodyTypeName = contentType
		default:
			bodyStructName := out.Name + "Body"
			out.BodyTypeName = Ref(bodyStructName)
			bodyType := r.ContentJSON.Value
			out.Body = bodyType.Type

			if bodyType.Spec.Value().AdditionalProperties.IsSet {
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

	// out.ContentType = contentType
	// out.PrivateName = out.Name
	// out.PrivateName = PrivateFieldName(out.Name)
	// // out.Status = status

	// var fields []GoStructField
	// if strings.EqualFold(status, "default") {
	// 	out.IsDefault = true
	// 	fields = append(fields, GoStructField{
	// 		Name: "Code",
	// 		Type: Int,
	// 	})
	// 	out.Args = append(out.Args, ResponseArg{
	// 		FieldName: "Code",
	// 		ArgName:   "code",
	// 		Type:      Int,
	// 	})
	// } else {
	// 	var err error
	// 	out.Status, err = strconv.Atoi(status)
	// 	if err != nil {
	// 		panic(fmt.Errorf("convert response status: %w", err))
	// 	}
	// }

	// if response.Value != nil && response.Value.Description != nil {
	// 	out.Description = *response.Value.Description
	// }
	// if resp.Description != nil {
	// 	out.Description = *resp.Description
	// }

	out.ResponserInterfaceName = h.ResponserInterfaceName

	// if s != nil {
	// 	out.IsBody = true

	// 	sr := NewSchemaRef(s)

	// 	fieldType := sr

	// 	if s.Ref == "" {
	// 		switch sr := sr.(type) {
	// 		case GoStruct:
	// 			bodyType := GoTypeDef{
	// 				Name: out.Name + "Body",
	// 				Type: sr,
	// 			}
	// 			if s.Value.AdditionalProperties != nil {
	// 				bodyType.Methods = append(bodyType.Methods,
	// 					GoVarDef{
	// 						Name:  "_",
	// 						Type:  GoType("json.Marshaler"),
	// 						Value: GoValue(`(*` + bodyType.Name + `)` + `(nil)`),
	// 					},
	// 					MarshalJSONFunc(bodyType.GoTypeRef(), sr),
	// 				)
	// 			}
	// 			out.Body = bodyType
	// 			fieldType = GoType(bodyType.Name)
	// 		}
	// 	}

	// 	fields = append(fields, GoStructField{
	// 		Name: "Body",
	// 		Type: fieldType,
	// 	})
	// 	out.Args = append(out.Args, ResponseArg{
	// 		FieldName: "Body",
	// 		ArgName:   "body",
	// 		Type:      fieldType,
	// 	})
	// }

	// pathHeaders := PathHeaders(resp.Headers)
	// fieldHeaders := make([]GoStructField, 0, len(pathHeaders))
	// for _, h := range pathHeaders {
	// 	sr := NewSchemaRef(h.Header.Value.Schema)

	// 	header := ResponseHeader{
	// 		Name:      h.Name,
	// 		FieldName: PublicFieldName(h.Name),
	// 		Type:      sr,
	// 	}

	// 	out.Headers = append(out.Headers, header)

	// 	fieldHeaders = append(fields, GoStructField{
	// 		Name: header.FieldName,
	// 		Type: header.Type,
	// 	})
	// 	out.Args = append(out.Args, ResponseArg{
	// 		FieldName: header.FieldName,
	// 		ArgName:   PrivateFieldName(header.FieldName),
	// 		IsHeader:  true,
	// 		Type:      header.Type,
	// 	})
	// }
	// if len(fieldHeaders) > 0 {
	// 	fields = append(fields, GoStructField{
	// 		Name: "Headers",
	// 		Type: GoStruct{
	// 			Fields: fieldHeaders,
	// 		},
	// 	})
	// }

	// out.Struct = GoTypeDef{
	// 	Comment: out.Description,
	// 	Name:    out.Name,
	// 	Type: GoStruct{
	// 		Fields: fields,
	// 	},
	// }

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
