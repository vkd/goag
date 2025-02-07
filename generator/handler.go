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

func NewHandler(o *Operation, basePathPrefix string, components Componenter, cfg Config) (zero *Handler, _ Imports, _ error) {
	out := &Handler{
		Operation: o,

		HandlerFuncName: string(o.Name) + "HandlerFunc",
		BasePathPrefix:  basePathPrefix,

		CanParseError: len(o.Params.Query.List) > 0 || len(o.Params.Path.List) > 0 || len(o.Params.Headers.List) > 0 || o.Body.GoTypeFn != nil || o.Body.Type.IsSet,
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
			resp := NewHandlerResponse(o.DefaultResponse.Response, o.Name, o.DefaultResponse.StatusCode, components, cfg, ResponseUsedIn{OperationName: o.Name, Status: o.DefaultResponse.StatusCode})
			out.DefaultResponse = &resp
		}
	}
	for _, r := range o.Responses {
		if r.ComponentRef == nil {
			out.Responses = append(out.Responses, NewHandlerResponse(r.Response, o.Name, r.StatusCode, components, cfg, ResponseUsedIn{OperationName: o.Name, Status: r.StatusCode}))
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
	TypeFn       RenderFunc
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

func NewHandlerQueryParameter(p *QueryParameter, cfg Config) (zero HandlerQueryParameter, _ Imports, _ error) {
	var ims Imports

	out := HandlerQueryParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    PublicFieldName(p.Name),
			TypeFn:       p.Type.RenderFieldType,
			FieldComment: strings.ReplaceAll(strings.TrimRight(p.Description, "\n "), "\n", "\n// "),
		},

		ParameterName: p.Name,
		Required:      p.Required,
		Parser:        p.Type,
	}

	return out, ims, nil
}

type PathParserConstant struct {
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
			TypeFn:       p.Type.RenderFieldType,
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
	var tp GoTypeRender = p.Type
	var parser Parser = p.Type

	fieldName := Title(p.Name)

	out := HandlerHeaderParameter{
		HandlerParameter: HandlerParameter{
			FieldName:    fieldName,
			TypeFn:       tp.RenderGoType,
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
	IsBodyReader bool
	GoTypeFn     GoTypeRenderFunc
	Body         *SchemaComponent
	BodyRenders  Renders
	ContentType  string
	// Headers     []ResponseHeader

	Struct         StructureType
	StructGoTypeFn GoTypeRenderFunc

	Args []ResponseArg
}

func NewHandlerResponse(r *Response, name OperationName, status string, components Componenter, cfg Config, ifaceNames ...ResponseUsedIn) HandlerResponse {
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
	} else if contentType, ok := r.ContentBody.Get(); ok {
		out.ContentType = contentType
	}

	if out.IsDefault {
		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name:        "Code",
			GoTypeFn:    StringRender(IntType{}.GoType()).Render,
			FieldTypeFn: StringRender(IntType{}.GoType()).Render,
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Code",
			ArgName:   "code",
			GoTypeFn:  StringRender(IntType{}.GoType()).Render,
		})
	}

	if contentJSON, ok := r.ContentJSON.Get(); ok {
		out.IsBody = true
		switch {
		case contentJSON.Type.Ref != nil:
			out.GoTypeFn = StringRender(contentJSON.Type.Ref.Name).Render
		case contentJSON.Type.IsCustom():
			out.GoTypeFn = contentJSON.Type.RenderGoType
		default:
			switch contentType := contentJSON.Type.Type.(type) {
			case SliceType:
				out.GoTypeFn = contentType.RenderGoType
			default:
				bodyStructName := out.Name + "Body"
				out.GoTypeFn = StringRender(bodyStructName).Render
				bodyType := contentJSON
				// out.Body = bodyType.Type
				sc := NewSchemaComponent(bodyStructName, bodyType.Type, components, cfg)
				out.Body = &sc
			}
		}

		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name:        "Body",
			GoTypeFn:    out.GoTypeFn,
			FieldTypeFn: RenderFunc(out.GoTypeFn),
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Body",
			ArgName:   "body",
			GoTypeFn:  out.GoTypeFn,
		})
	} else if _, ok := r.ContentBody.Get(); ok {
		out.IsBodyReader = true
		out.GoTypeFn = StringRender("io.ReadCloser").Render

		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name:        "Body",
			GoTypeFn:    out.GoTypeFn,
			FieldTypeFn: RenderFunc(out.GoTypeFn),
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Body",
			ArgName:   "body",
			GoTypeFn:  out.GoTypeFn,
		})
	}

	var headersStruct StructureType
	for _, h := range r.Headers {
		headersStruct.Fields = append(headersStruct.Fields, StructureField{
			Name:        h.FieldName,
			GoTypeFn:    h.Schema.RenderGoType,
			FieldTypeFn: h.Schema.RenderGoType,
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: h.FieldName,
			ArgName:   PrivateFieldName(h.FieldName),
			IsHeader:  true,
			GoTypeFn:  h.Schema.RenderGoType,
		})
	}
	if len(headersStruct.Fields) > 0 {
		out.Struct.Fields = append(out.Struct.Fields, StructureField{
			Name:        "Headers",
			GoTypeFn:    headersStruct.RenderGoType,
			FieldTypeFn: headersStruct.RenderGoType,
		})
	}

	out.UsedIn = ifaceNames
	out.StructGoTypeFn = out.Struct.RenderGoType

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
	GoTypeFn  GoTypeRenderFunc
}

type ResponseUsedIn struct {
	OperationName OperationName
	Status        string
}
