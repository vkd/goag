package generator

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type Handlers struct {
	PackageName string
	BasePath    string

	Handlers []Handler

	IsWriteJSONFunc bool
}

func NewHandlers(pname string, s *openapi3.Swagger, basePath string) (zero Handlers, _ error) {
	var out Handlers
	out.PackageName = pname
	out.BasePath = basePath

	out.Handlers = make([]Handler, 0, len(s.Paths))
	for _, p := range Paths(s.Paths) {
		for _, o := range PathOperations(p.Item) {
			h, err := NewHandler(o.Operation, p.Path, o.Method, p.Item.Parameters)
			if err != nil {
				return zero, fmt.Errorf("new handler for [%s]%q: %w", o.Method, p.Path, err)
			}
			out.Handlers = append(out.Handlers, h)
			if h.IsWriteJSONFunc {
				out.IsWriteJSONFunc = true
			}
		}
	}

	return out, nil
}

type Handler struct {
	Name        string
	Path        string
	Method      string
	Description string
	Summary     string

	ResponserInterfaceName string

	Parameters struct {
		Queries []Param

		Path        []PathParameter
		PathParsers []Render
	}

	RequestBody Render

	ParamParsers []Render

	// response type which implement handler's responser interface
	Responses []Render

	IsWriteJSONFunc bool
}

// const txtHandler = ``

// var tmHandler = template.Must(template.New("Handler").Parse(txtHandler))

func NewHandler(p *openapi3.Operation, path, method string, params openapi3.Parameters) (zero Handler, _ error) {
	var out Handler
	out.Name = HandlerName(path, method)
	out.Path = path
	out.Method = method
	out.Description = strings.ReplaceAll(strings.TrimSpace(p.Description), "\n", "\n// ")
	out.Summary = p.Summary

	for _, pr := range params {
		p := pr.Value
		switch p.In {
		case openapi3.ParameterInQuery:
			par := NewParam(p)
			out.Parameters.Queries = append(out.Parameters.Queries, par)
		case openapi3.ParameterInPath:
			pp := NewPathParameter(p)
			out.Parameters.Path = append(out.Parameters.Path, pp)
		}
	}
	for _, pr := range p.Parameters {
		p := pr.Value
		switch p.In {
		case openapi3.ParameterInQuery:
			par := NewParam(p)
			out.Parameters.Queries = append(out.Parameters.Queries, par)
		case openapi3.ParameterInPath:
			pp := NewPathParameter(p)
			out.Parameters.Path = append(out.Parameters.Path, pp)
		}
	}

	if len(out.Parameters.Path) > 0 {
		var err error
		out.Parameters.PathParsers, err = NewPathParamsParsers(path, out.Parameters.Path)
		if err != nil {
			return zero, fmt.Errorf("new path params parsers: %w", err)
		}
	}

	if p.RequestBody != nil {
		br := NewBodyRef(p.RequestBody)
		out.RequestBody = br
	}

	out.ResponserInterfaceName = "write" + out.Name + "Response"

	for _, r := range PathResponses(p.Responses) {
		if len(r.Response.Value.Content) == 0 {
			resp := NewResponse(nil, out.Name, out.ResponserInterfaceName, r.Code, "", r.Response.Value, "|", r.Code, r.Response)
			out.Responses = append(out.Responses, resp)
			if resp.IsBody {
				out.IsWriteJSONFunc = true
			}
		} else {
			for mtype, ct := range r.Response.Value.Content {
				resp := NewResponse(ct.Schema, out.Name, out.ResponserInterfaceName, r.Code, mtype, r.Response.Value, "|", r.Code, r.Response)
				out.Responses = append(out.Responses, resp)
				if resp.IsBody {
					out.IsWriteJSONFunc = true
				}
			}
		}
	}

	return out, nil
}

func HandlerName(path, method string) string {
	var suffix string
	if strings.HasSuffix(path, "/") {
		suffix = "RT"
	}
	return strings.Title(strings.ToLower(method)) + PathName(path) + suffix
}

func PathName(path string) string {
	ps := strings.Split(path, "/")
	for i, p := range ps {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			p = p[1 : len(p)-1]
		} else {
			p = strings.ToLower(p)
		}
		ps[i] = PublicFieldName(p)
	}
	return strings.Join(ps, "")
}

type ResponseHeader struct {
	Name      string
	FieldName string
	Type      Render
}

type Param struct {
	Field  Render
	Parser Render
}

func NewParam(p *openapi3.Parameter) Param {
	sr := NewSchemaRef(p.Schema)
	if !p.Required {
		sr = NewOptionalParam(sr)
	}
	f := GoStructField{
		Name:    PublicFieldName(p.Name),
		Type:    sr,
		Comment: p.Description,
	}
	prs := NewQueryParser(p, f)
	out := Param{
		Field:  f,
		Parser: prs,
	}
	return out
}

type OptionalParam struct {
	SchemaRender
}

func NewOptionalParam(sr SchemaRender) SchemaRender {
	if _, ok := sr.(interface{ Optionable() }); ok {
		return sr
	}
	return OptionalParam{SchemaRender: sr}
}

func (r OptionalParam) String() (string, error) {
	str, err := r.SchemaRender.String()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("*%s", str), nil
}

type StringsParser interface {
	StringsParser(from, to string, _ FuncNewError) Render
}

func NewQueryParser(p *openapi3.Parameter, field GoStructField) Render {
	s := NewSchemaRef(p.Schema)

	from := "q"
	toOrig := "params." + field.Name
	to := toOrig
	mkErr := NewQueryErrorFunc(p.Name)

	var conv Render
	if sp, ok := s.(StringsParser); ok {
		conv = sp.StringsParser(from, to, mkErr)
	} else {
		to := "v"
		conv = s.Parser(from+"[0]", to, mkErr)

		_, optionable := s.(interface{ Optionable() })

		if !p.Required && !optionable {
			to = "&" + to
		}
		conv = Combine{conv, Assign{to, toOrig}}
	}

	return QueryParser{
		QueryVarName:  "q",
		ParameterName: p.Name,
		Convert:       conv,
		Required:      p.Required,
	}
}

type QueryParser struct {
	QueryVarName  string
	ParameterName string
	Convert       Render
	Required      bool
}

var tmQueryParser = template.Must(template.New("QueryParser").Parse(`
{{- .QueryVarName}}, ok := query["{{.ParameterName}}"]
{{if .Required}}if !ok {
	return zero, fmt.Errorf("query parameter '{{.ParameterName}}': is required")
}
{{end -}}
if ok && len(q) > 0 {
	{{.Convert.String}}
}`))

func (i QueryParser) String() (string, error) {
	return String(tmQueryParser, i)
}

func NewQueryErrorFunc(name string) FuncNewError {
	return func(s string) string {
		return `ErrParseQueryParam{Name: "` + name + `", Err: ` + s + `}`
	}
}

func NewBodyRef(spec *openapi3.RequestBodyRef) Render {
	if spec.Ref != "" {
		return NewRef(spec.Ref)
	}
	return NewSchemaRef(spec.Value.Content.Get("application/json").Schema)
}

// --------------- Helpers: map -> slice ---------------

type PathItem struct {
	Path string
	Item *openapi3.PathItem
}

func Paths(spec openapi3.Paths) []PathItem {
	paths := make([]PathItem, 0, len(spec))
	for k, v := range spec {
		paths = append(paths, PathItem{k, v})
	}
	sort.Slice(paths, func(i, j int) bool { return paths[i].Path < paths[j].Path })
	return paths
}

type PathOperationItem struct {
	Method    string
	Operation *openapi3.Operation
}

func PathOperations(spec *openapi3.PathItem) []PathOperationItem {
	out := make([]PathOperationItem, 0, 9)
	for _, v := range []struct {
		method string
		o      *openapi3.Operation
	}{
		{http.MethodGet, spec.Get},
		{http.MethodPost, spec.Post},
		{http.MethodPatch, spec.Patch},
		{http.MethodPut, spec.Put},
		{http.MethodDelete, spec.Delete},
		{http.MethodConnect, spec.Connect},
		{http.MethodHead, spec.Head},
		{http.MethodOptions, spec.Options},
		{http.MethodTrace, spec.Trace},
	} {
		if v.o != nil {
			out = append(out, PathOperationItem{v.method, v.o})
		}
	}
	return out
}

type ResponseItem struct {
	Code     string
	Response *openapi3.ResponseRef
}

func PathResponses(spec openapi3.Responses) []ResponseItem {
	out := make([]ResponseItem, 0, len(spec))
	var defResponse *openapi3.ResponseRef
	for k, v := range spec {
		if k == "default" {
			defResponse = v
			continue
		}
		out = append(out, ResponseItem{k, v})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Code < out[j].Code })
	if defResponse != nil {
		out = append(out, ResponseItem{"default", defResponse})
	}
	return out
}

type HeaderItem struct {
	Name   string
	Header *openapi3.HeaderRef
}

func PathHeaders(spec openapi3.Headers) []HeaderItem {
	out := make([]HeaderItem, 0, len(spec))
	for name, h := range spec {
		out = append(out, HeaderItem{name, h})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
