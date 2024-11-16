package generator

type HandlersFileTemplate struct {
	Handlers []HandlerTemplate

	IsWriteJSONFunc  bool
	IsCustomMaybe    bool
	IsCustomNullable bool
}

func NewHandlersFileTemplate(hs []*Handler, isWriteJSONFunc bool, cfg Config) HandlersFileTemplate {

	handlers := make([]HandlerTemplate, 0, len(hs))
	for _, h := range hs {
		handlers = append(handlers, NewHandlerTemplate(h))
	}

	return HandlersFileTemplate{
		Handlers: handlers,

		IsWriteJSONFunc:  isWriteJSONFunc,
		IsCustomMaybe:    cfg.Maybe.Type != "",
		IsCustomNullable: cfg.Nullable.Type != "",
	}
}

func (t HandlersFileTemplate) Render() (string, error) {
	return ExecuteTemplate("FileHandler", t)
}

type HandlerTemplate struct {
	Name        OperationName
	Description string
	Summary     string

	HandlerFuncName string
	BasePathPrefix  string

	CanParseError bool

	ParametersQuery  []HandlerParameterQueryTemplate
	ParametersPath   []HandlerParameterPathTemplate
	ParametersHeader []HandlerParameterHeaderTemplate

	GoTypeFn GoTypeRenderFunc
	BodyType *SchemaComponent

	PathParsers []Parser

	Responses       []HandlerResponseTemplate
	DefaultResponse Maybe[HandlerResponseTemplate]
}

func NewHandlerTemplate(h *Handler) HandlerTemplate {
	var defaultResponse Maybe[HandlerResponseTemplate]
	if h.DefaultResponse != nil {
		defaultResponse = Just(NewHandlerResponseTemplate(*h.DefaultResponse))
	}
	var bodyType *SchemaComponent
	if h.Body.Type.IsSet {
		bodyType = &h.Body.Type.Value
	}
	return HandlerTemplate{
		Name:        h.Name,
		Description: h.Description,
		Summary:     h.Summary,

		HandlerFuncName: h.HandlerFuncName,
		BasePathPrefix:  h.BasePathPrefix,

		CanParseError: h.CanParseError,

		ParametersQuery:  NewHandlerParameterQueryTemplates(h.Parameters.Query),
		ParametersPath:   NewHandlerParameterPathTemplates(h.Parameters.Path),
		ParametersHeader: NewHandlerParameterHeaderTemplates(h.Parameters.Header),

		GoTypeFn: h.Body.GoTypeFn,
		BodyType: bodyType,

		PathParsers: h.PathParsers,

		Responses:       NewHandlerResponseTemplates(h.Responses),
		DefaultResponse: defaultResponse,
	}
}

func (t HandlerTemplate) Render() (string, error) {
	return ExecuteTemplate("Handler", t)
}

type HandlerParameterQueryTemplate struct {
	HandlerParameter HandlerParameterTemplate
	ParameterName    string
	Required         bool

	ParseStrings func(to, from string, isNew bool, mkErr ErrorRender) (string, error)
}

func NewHandlerParameterQueryTemplate(p HandlerQueryParameter) HandlerParameterQueryTemplate {
	return HandlerParameterQueryTemplate{
		HandlerParameter: NewHandlerParameterTemplate(p.HandlerParameter),
		ParameterName:    p.ParameterName,
		Required:         p.Required,

		ParseStrings: p.Parser.ParseStrings,
	}
}

func NewHandlerParameterQueryTemplates(ps []HandlerQueryParameter) []HandlerParameterQueryTemplate {
	out := make([]HandlerParameterQueryTemplate, 0, len(ps))
	for _, p := range ps {
		out = append(out, NewHandlerParameterQueryTemplate(p))
	}
	return out
}

type HandlerParameterPathTemplate struct {
	HandlerParameter HandlerParameterTemplate
}

func NewHandlerParameterPathTemplate(p HandlerPathParameter) HandlerParameterPathTemplate {
	return HandlerParameterPathTemplate{
		HandlerParameter: NewHandlerParameterTemplate(p.HandlerParameter),
	}
}

func NewHandlerParameterPathTemplates(ps []HandlerPathParameter) []HandlerParameterPathTemplate {
	out := make([]HandlerParameterPathTemplate, 0, len(ps))
	for _, p := range ps {
		out = append(out, NewHandlerParameterPathTemplate(p))
	}
	return out
}

type HandlerParameterHeaderTemplate struct {
	HandlerParameter HandlerParameterTemplate

	ParameterName string
	Required      bool
	Parser        Parser
}

func NewHandlerParameterHeaderTemplate(p HandlerHeaderParameter) HandlerParameterHeaderTemplate {
	return HandlerParameterHeaderTemplate{
		HandlerParameter: NewHandlerParameterTemplate(p.HandlerParameter),

		ParameterName: p.ParameterName,
		Required:      p.Required,
		Parser:        p.Parser,
	}
}

func NewHandlerParameterHeaderTemplates(ps []HandlerHeaderParameter) []HandlerParameterHeaderTemplate {
	out := make([]HandlerParameterHeaderTemplate, 0, len(ps))
	for _, p := range ps {
		out = append(out, NewHandlerParameterHeaderTemplate(p))
	}
	return out
}

type HandlerParameterTemplate struct {
	FieldName    string
	FieldComment string
	TypeFn       RenderFunc
}

func NewHandlerParameterTemplate(h HandlerParameter) HandlerParameterTemplate {
	return HandlerParameterTemplate{
		FieldName:    h.FieldName,
		FieldComment: h.FieldComment,
		TypeFn:       h.TypeFn,
	}
}

func (t HandlerParameterTemplate) Render() (string, error) {
	return ExecuteTemplate("HandlerParameter", t)
}

type HandlerResponseTemplate struct {
	Name        string
	Description string
	HandlerName OperationName

	Status    string
	IsDefault bool

	UsedIn []ResponseUsedIn

	IsBody      bool
	GoTypeFn    GoTypeRenderFunc
	Body        *SchemaComponent
	BodyRenders Renders
	ContentType string

	Struct         StructureType
	StructGoTypeFn GoTypeRenderFunc

	Args    []ResponseArg
	Headers []ResponseHeader

	IsComponent bool
}

func NewHandlerResponseTemplate(r HandlerResponse) HandlerResponseTemplate {
	return HandlerResponseTemplate{
		Name:        r.Name,
		Description: r.Description,
		HandlerName: r.HandlerName,

		Status:    r.Status,
		IsDefault: r.IsDefault,

		UsedIn: r.UsedIn,

		IsBody:      r.IsBody,
		GoTypeFn:    r.GoTypeFn,
		Body:        r.Body,
		BodyRenders: r.BodyRenders,
		ContentType: r.ContentType,

		Struct:         r.Struct,
		StructGoTypeFn: r.StructGoTypeFn,

		Args:    r.Args,
		Headers: r.Headers,

		IsComponent: false,
	}
}

func NewHandlerResponseTemplates(rs []HandlerResponse) []HandlerResponseTemplate {
	out := make([]HandlerResponseTemplate, 0, len(rs))
	for _, r := range rs {
		out = append(out, NewHandlerResponseTemplate(r))
	}
	return out
}

func (t HandlerResponseTemplate) Render() (string, error) {
	return ExecuteTemplate("ResponseComponent", t)
}
