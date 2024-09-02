package generator

import (
	"fmt"
	"strings"
)

type Handler struct {
	*Operation

	HandlerFuncName string
	BasePathPrefix  string

	CanParseError bool

	Parameters  HandlerParameters
	PathParsers []Parser

	DefaultResponse *HandlerResponse
	Responses       []HandlerResponse
}

func NewHandler(o *Operation, basePathPrefix string, cfg Config) (zero *Handler, _ Imports, _ error) {
	out := &Handler{
		Operation: o,

		HandlerFuncName: string(o.Name) + "HandlerFunc",
		BasePathPrefix:  basePathPrefix,

		CanParseError: len(o.Params.Query.List) > 0 || len(o.Params.Path.List) > 0 || len(o.Params.Headers.List) > 0 || o.Body.TypeName != nil || o.Body.Type.IsSet,
	}
	ps, imports, err := NewHandlerParameters(o.Params, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("params: %w", err)
	}
	out.Parameters = ps

	var pathRenders []Parser
	for _, pe := range o.PathBuilder {
		if peParam, ok := pe.Param.Get(); ok {
			pathRenders = append(pathRenders, PathParserVariable{
				FieldName: peParam.FieldName,
				Name:      peParam.Name,
				Convert:   peParam.Type,
			})
		} else if pe.Raw != "" {
			pathRenders = append(pathRenders, PathParserConstant{
				Prefix:   pe.Raw,
				FullPath: o.Path.Raw,
			})
		}
	}
	out.PathParsers = pathRenders

	if o.DefaultResponse != nil {
		if o.DefaultResponse.ComponentRef == nil {
			resp := NewHandlerResponse(o.DefaultResponse.Response, o.Name, o.DefaultResponse.StatusCode, ResponseUsedIn{OperationName: o.Name, Status: o.DefaultResponse.StatusCode})
			out.DefaultResponse = &resp
		}
	}
	for _, r := range o.Responses {
		if r.ComponentRef == nil {
			out.Responses = append(out.Responses, NewHandlerResponse(r.Response, o.Name, r.StatusCode, ResponseUsedIn{OperationName: o.Name, Status: r.StatusCode}))
		}
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

func NewHandlerParameters(p OperationParams, cfg Config) (zero HandlerParameters, _ Imports, _ error) {
	out := HandlerParameters{}
	var imports Imports
	for _, q := range p.Query.List {
		p, ims, err := NewHandlerQueryParameter(q.V, cfg)
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
		p, ims, err := NewHandlerHeaderParameter(q.V, cfg)
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
	IsOptional    bool
	Parser        Parser
}

func NewHandlerQueryParameter(p *QueryParameter, cfg Config) (zero HandlerQueryParameter, _ Imports, _ error) {
	var ims Imports

	var tpRender Render = p.Type
	var isOptional bool
	if !p.Required {
		// switch tp := p.Type.(type) {
		// case CustomType:
		// default:
		tpRender = NewOptionalType(p.Type, cfg)
		if cfg.Maybe.Import != "" {
			ims = append(ims, Import(cfg.Maybe.Import))
		}
		// tp = NewOptionalType(p.Type)
		// }
		isOptional = true
	}

	out := HandlerQueryParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    PublicFieldName(p.Name),
			FieldType:    tpRender,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Required:      p.Required,
		Parser:        p.Type,
		IsOptional:    isOptional,
	}

	return out, ims, nil
}

func (p HandlerQueryParameter) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	// switch parser := p.Parser.(type) {
	// case SliceType:
	// 	return parser.ParseStrings(to, from, isNew, mkErr)
	// case Ref[specification.Schema]:
	// 	if parser.SchemaType.Value().Type == "array" {
	// 		return parser.ParseQuery(to, from, isNew, mkErr)
	// 	}
	// 	return parser.ParseSchema(to, from+"[0]", isNew, mkErr)
	// case Ref[specification.QueryParameter]:
	// 	return parser.ParseQuery(to, from, isNew, mkErr)
	// case CustomType:
	// 	return parser.ParseStrings(to, from, isNew, mkErr)
	// }
	return p.Parser.ParseStrings(to, from, isNew, mkErr)
}

type PathParserConstant struct {
	SingleValue
	Prefix   string
	FullPath string
}

func (p PathParserConstant) ParseString(_, _ string, _ bool, _ ErrorRender) (string, error) {
	return ExecuteTemplate("PathParserConstant", p)
}

func (p PathParserConstant) ParseStrings(_, _ string, _ bool, _ ErrorRender) (string, error) {
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

func (p PathParserVariable) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
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
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.Description, "\n "), "\n", "\n// "),
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
}

func NewHandlerHeaderParameter(p *HeaderParameter, cfg Config) (zero HandlerHeaderParameter, _ Imports, _ error) {
	var ims Imports
	var tp Render = p.Schema
	var parser Parser = p.Schema

	if !p.Required {
		ot := NewOptionalType(p.Schema, cfg)
		if cfg.Maybe.Import != "" {
			ims = append(ims, Import(cfg.Maybe.Import))
		}
		tp = ot
		parser = ot
	}

	fieldName := PublicFieldName(p.Name)

	out := HandlerHeaderParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    fieldName,
			FieldType:    tp,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Required:      p.Required,
		Parser:        parser,
	}

	return out, ims, nil
}

type HandlerResponse struct {
	*Response

	Name string
	// PrivateName string
	HandlerName OperationName

	Status    string
	IsDefault bool

	UsedIn []ResponseUsedIn

	IsBody       bool
	BodyTypeName Render
	Body         Render
	BodyRenders  Renders
	ContentType  string
	// Headers     []ResponseHeader

	Struct StructureType

	Args []ResponseArg
}

func NewHandlerResponse(r *Response, name OperationName, status string, ifaceNames ...ResponseUsedIn) HandlerResponse {
	out := HandlerResponse{
		Response: r,

		HandlerName: name,

		Status:    status,
		IsDefault: status == "default",
	}

	out.Name = string(name) + "Response" + strings.Title(status)
	if r.ContentJSON.IsSet {
		out.Name += "JSON"
		out.ContentType = "application/json"
	}

	if out.IsDefault {
		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name: "Code",
			Type: StringRender(IntType{}.GoType()),
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Code",
			ArgName:   "code",
			Type:      StringRender(IntType{}.GoType()),
		})
	}

	if contentJSON, ok := r.ContentJSON.Get(); ok {
		out.IsBody = true
		switch {
		case contentJSON.Type.Ref != nil:
			out.BodyTypeName = StringRender(contentJSON.Type.Ref.Name)
		case contentJSON.Type.IsCustom():
			out.BodyTypeName = contentJSON.Type
		default:
			switch contentType := contentJSON.Type.Type.(type) {
			case SliceType:
				out.BodyTypeName = contentType
			default:
				bodyStructName := out.Name + "Body"
				out.BodyTypeName = StringRender(bodyStructName)
				bodyType := contentJSON
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
							if st, ok := bodyType.Type.Type.(StructureType); ok {
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

	out.UsedIn = ifaceNames

	return out
}

func (h HandlerResponse) Render() (string, error) {
	return ExecuteTemplate("ResponseComponent", ResponseComponent{
		Name:            h.Name,
		Description:     h.Description,
		HandlerResponse: h,
	})
}

type ResponseArg struct {
	FieldName string
	ArgName   string
	IsHeader  bool
	Type      Render
}

type ResponseUsedIn struct {
	OperationName OperationName
	Status        string
}
