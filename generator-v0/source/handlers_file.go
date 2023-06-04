package source

type HandlersFile struct {
	Handlers []Handler

	IsWriteJSONFunc bool
}

type Handler struct {
	Name        string
	Description string
	Summary     string

	BasePathPrefix string

	CanParseError bool

	ResponserInterfaceName string

	Parameters struct {
		Query   []Param
		Path    []Param
		Headers []Param

		PathParsers []Templater
	}

	Body struct {
		TypeName Templater
	}

	IsWriteJSONFunc bool

	Responses []Templater
}

type Param struct {
	Field  Templater
	Parser Templater
}

var tmHandlersFile = InitTemplate("HandlersFile", `


{{ range $_, $h := .Handlers }}
{{- $name := $h.Name}}
{{- $handlerFunc := print $name "HandlerFunc" }}
{{- $requester := print $name "Requester" }}
{{- $requestParams := print "request" $name "Params" }}
{{- $request := print $name "Request" }}
{{- $newParams := print "new" $name "Params" }}
{{- $responder := print $name "Responder" }}
{{- $writeResponse := $h.ResponserInterfaceName }}
// ---------------------------------------------
// {{$name}} - {{$h.Description}}
// ---------------------------------------------

{{if $h.Summary}}// {{$handlerFunc}} - {{$h.Summary}}{{end}}
type {{ $handlerFunc }} func(r {{$requester}}) {{$responder}}

func (f {{$handlerFunc}}) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f({{$requestParams}}{Request: r}).{{$writeResponse}}(w)
}

type {{$requester}} interface {
	Parse() {{if $h.CanParseError}}({{end}}{{$request}}{{if $h.CanParseError}}, error){{end}}
}

type {{$requestParams}} struct {
	Request *http.Request
}

func (r {{$requestParams}}) Parse() {{if $h.CanParseError}}({{end}}{{$request}}{{if $h.CanParseError}}, error){{end}} {
	return {{$newParams}}(r.Request)
}

type {{$request}} struct {
	HTTPRequest *http.Request

	{{ if $h.Parameters.Query }}
	Query struct{
		{{ range $_, $p := .Parameters.Query }}
		{{ exec $p.Field }}
		{{ end }}
	}
	{{ end }}

	{{ if $h.Parameters.Path }}
	Path struct{
		{{ range $_, $p := .Parameters.Path }}
		{{ exec $p.Field }}
		{{ end }}
	}
	{{ end }}

	{{ if $h.Parameters.Headers }}
	Headers struct{
		{{ range $_, $p := .Parameters.Headers }}
		{{ exec $p.Field }}
		{{ end }}
	}
	{{ end }}

	{{ if $h.Body.TypeName }}
	Body {{ exec $h.Body.TypeName }}
	{{ end }}
}

func {{ $newParams }}(r *http.Request) (zero {{ $request }}{{ if $h.CanParseError }}, _ error{{ end }}) {
	var params {{ $request }}
	params.HTTPRequest = r

	{{ if $h.Parameters.Query }}
	// Query parameters
	{
		query := r.URL.Query()
		{{- range $_, $q := .Parameters.Query }}
		{
			{{ exec $q.Parser }}
		}
		{{- end }}
	}
	{{ end }}

	{{ if $h.Parameters.Headers }}
	// Headers
	{
		header := r.Header
		{{- range $_, $h := .Parameters.Headers }}
		{
			{{ exec $h.Parser }}
		}
		{{- end}}
	}
	{{ end }}

	{{ if $h.Parameters.Path }}
	// Path parameters
	{
		p := r.URL.Path

		{{- if $h.BasePathPrefix }}
		if !strings.HasPrefix(p, "{{ $h.BasePathPrefix }}") {
			return zero, fmt.Errorf("wrong path: expected '{{ $h.BasePathPrefix }}...'")
		}
		p = p[{{ len $h.BasePathPrefix }}:] // "{{ $h.BasePathPrefix }}"

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '{{ $h.BasePathPrefix }}/...'")
		}

		{{ end }}

		{{- range $_, $p := $h.Parameters.PathParsers }}

		{{ exec $p }}
		{{- end }}
	}
	{{ end }}

	{{- if $h.Body.TypeName }}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}
	{{- end }}

	return params{{ if $h.CanParseError }}, nil{{ end }}
}

type {{ $responder }} interface {
	{{$writeResponse}}(w http.ResponseWriter)
}

{{ range $_, $r := .Responses }}
{{ exec $r }}
{{ end }}

{{ end -}}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

{{ if .IsWriteJSONFunc -}}
func writeJSON(w io.Writer, v interface{}, name string) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		LogError(fmt.Errorf("write json response %q: %w", name, err))
	}
}

{{ end -}}

type ErrParseParam struct {
	In        string
	Parameter string
	Reason    string
	Err       error
}

func (e ErrParseParam) Error() string {
	return fmt.Sprintf("%s parameter '%s': %s: %v", e.In, e.Parameter, e.Reason, e.Err)
}

func (e ErrParseParam) Unwrap() error { return e.Err }
`)

func (h HandlersFile) String() (string, error) { return tmHandlersFile.String(h) }
