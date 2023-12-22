package generator

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/vkd/goag/generator"
	"github.com/vkd/goag/generator-v0/source"
	"github.com/vkd/goag/specification"
)

type Handlers struct {
	Handlers []Handler

	IsWriteJSONFunc bool
}

func NewHandlers(s *specification.Spec, basePath string) (zero Handlers, _ error) {
	var out Handlers
	out.Handlers = make([]Handler, 0, len(s.Operations))
	for _, o := range s.Operations {
		h, err := NewHandler(o.Operation, o.PathItem.Path, o.Method, o.PathItem.PathItem.Parameters)
		if err != nil {
			return zero, fmt.Errorf("new handler for [%s]%q: %w", o.Method, o.PathItem.Path.Spec, err)
		}
		h.BasePathPrefix = basePath
		out.Handlers = append(out.Handlers, h)
		if h.IsWriteJSONFunc {
			out.IsWriteJSONFunc = true
		}
	}

	return out, nil
}

type Handler struct {
	generator.HandlerOld

	Path   string
	Method string

	Responses []Response

	Params struct {
		Query []Param
	}
}

func NewHandler(p *openapi3.Operation, path specification.Path, method string, params openapi3.Parameters) (zero Handler, _ error) {
	var out Handler
	out.Name = HandlerName(p.OperationID, path, method)
	out.Path = path.Spec
	out.Method = method
	out.Description = strings.ReplaceAll(strings.TrimSpace(p.Description), "\n", "\n// ")
	out.Summary = p.Summary

	var allParams openapi3.Parameters
	allParams = append(allParams, params...)
	allParams = append(allParams, p.Parameters...)

	var pathParams []PathParameter

	for _, pr := range allParams {
		p := pr.Value
		switch p.In {
		case openapi3.ParameterInQuery:
			par := NewQueryParam(p)
			out.Params.Query = append(out.Params.Query, par)
			out.Parameters.Query = append(out.Parameters.Query, generator.Param{
				Field:  par.Field,
				Parser: par.Parser,
			})
		case openapi3.ParameterInPath:
			pp := NewPathParameter(p)
			pathParams = append(pathParams, pp)
			out.Parameters.Path = append(out.Parameters.Path, generator.Param{
				Field: pp.Field,
				// Parser: NewPathParamsParsers(),
			})
		case openapi3.ParameterInHeader:
			pp := NewHeaderParam(p)
			out.Parameters.Headers = append(out.Parameters.Headers, generator.Param{
				Field:  pp.Field,
				Parser: pp.Parser,
			})
		}
	}

	if len(out.Parameters.Path) > 0 {
		var err error
		out.Parameters.PathParsers, err = NewPathParamsParsers(string(path.Spec), pathParams)
		if err != nil {
			return zero, fmt.Errorf("new path params parsers: %w", err)
		}
	}

	if p.RequestBody != nil {
		br := NewBodyRef(p.RequestBody)
		out.Body.TypeName = br
	}

	out.ResponserInterfaceName = PrivateFieldName(out.Name)

	var responses []Response

	for _, r := range PathResponses(p.Responses) {
		if len(r.Response.Value.Content) == 0 {
			resp := NewResponse(nil, out.Name, out.ResponserInterfaceName, r.Code, "", r.Response.Value, "|", r.Code, r.Response)
			responses = append(responses, resp)
			if resp.IsBody {
				out.IsWriteJSONFunc = true
			}
		} else {
			for mtype, ct := range r.Response.Value.Content {
				resp := NewResponse(ct.Schema, out.Name, out.ResponserInterfaceName, r.Code, mtype, r.Response.Value, "|", r.Code, r.Response)
				responses = append(responses, resp)
				if resp.IsBody {
					out.IsWriteJSONFunc = true
				}
			}
		}
	}
	for _, r := range responses {
		out.Responses = append(out.Responses, r)
		out.HandlerOld.Responses = append(out.HandlerOld.Responses, r)
	}

	out.CanParseError = len(out.Parameters.Query) > 0 || len(out.Parameters.Path) > 0 || len(out.Parameters.Headers) > 0 || out.Body.TypeName != nil

	return out, nil
}

var HandlerName = generator.OperationName

// func HandlerName(path, method string) string {
// 	var suffix string
// 	if strings.HasSuffix(path, "/") {
// 		// "/shops" and "/shops/" need to have separate handlers
// 		suffix = "RT"
// 	}
// 	return strings.Title(strings.ToLower(method)) + PathName(path) + suffix
// }

func PathName(s string) string {
	if s == "" {
		return ""
	}
	return specification.PathOld(s).Name(func(s specification.Prefix) string { return strings.Title(strings.ToLower(s.Name())) }, "")
}

// func PathName(path string) string {
// 	ps := strings.Split(path, "/")
// 	for i, p := range ps {
// 		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
// 			p = p[1 : len(p)-1]
// 			// p = "var"
// 		} else {
// 			p = strings.ToLower(p)
// 		}
// 		ps[i] = PublicFieldName(p)
// 	}
// 	return strings.Join(ps, "")
// }

type ResponseHeader struct {
	Name      string
	FieldName string
	Type      SchemaRender
}

type Param struct {
	Field  GoStructField
	Parser Render

	Parameter *openapi3.Parameter
}

func NewQueryParam(p *openapi3.Parameter) Param {
	sr := NewSchemaRef(p.Schema)
	if !p.Required {
		sr = NewOptionalParam(sr)
	}
	f := GoStructField{
		Name:    PublicFieldName(p.Name),
		Type:    sr,
		Comment: p.Description,
	}
	f.Comment = strings.ReplaceAll(strings.TrimRight(f.Comment, "\n "), "\n", "\n// ")
	prs := NewQueryParser(p, f)
	out := Param{
		Field:     f,
		Parser:    prs,
		Parameter: p,
	}
	return out
}

func NewHeaderParam(p *openapi3.Parameter) Param {
	sr := NewSchemaRef(p.Schema)
	if !p.Required {
		sr = NewOptionalParam(sr)
	}
	f := GoStructField{
		Name:    PublicFieldName(p.Name),
		Type:    sr,
		Comment: p.Description,
	}
	prs := NewHeaderParser(p, f)
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

func (r OptionalParam) Format(s string) source.Templater {
	return r.SchemaRender.Format("*" + s)
}

func NewQueryParser(p *openapi3.Parameter, field GoStructField) Render {
	s := NewSchemaRef(p.Schema)

	from := "q"
	to := "params.Query." + field.Name
	mkErr := source.QueryParseError(p.Name)

	return QueryParser{
		QueryVarName:  from,
		ParameterName: p.Name,
		Convert:       NewStringsParser(s, from, to, !p.Required, mkErr),
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
if ok && len({{.QueryVarName}}) > 0 {
	{{.Convert.String}}
}`))

func (i QueryParser) String() (string, error) {
	return String(tmQueryParser, i)
}

func NewHeaderParser(p *openapi3.Parameter, field GoStructField) Render {
	s := NewSchemaRef(p.Schema)

	from := "hs"
	to := "params.Headers." + field.Name
	mkErr := source.HeaderParseError(p.Name)

	return HeaderParser{
		HeaderVarName: from,
		ParameterName: p.Name,
		Convert:       NewStringsParser(s, from, to, !p.Required, mkErr),
		Required:      p.Required,
	}
}

type HeaderParser struct {
	HeaderVarName string
	ParameterName string
	Convert       Render
	Required      bool
}

var tmHeaderParser = template.Must(template.New("HeaderParser").Parse(`
{{- .HeaderVarName}} := header.Values("{{.ParameterName}}")
{{if .Required}}if len({{.HeaderVarName}}) == 0 {
	return zero, fmt.Errorf("header parameter '{{.ParameterName}}': is required")
}
{{end -}}
if len({{.HeaderVarName}}) > 0 {
	{{.Convert.String}}
}`))

func (i HeaderParser) String() (string, error) {
	return String(tmHeaderParser, i)
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
