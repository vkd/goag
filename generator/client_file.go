package generator

import "github.com/vkd/goag/specification"

type Client struct {
	ClientHandlers []ClientHandler
}

func (g *Generator) ClientFile() (Templater, error) {
	var hs []ClientHandler
	for _, o := range g.Operations {
		hs = append(hs, NewClientHandler(o))
	}
	return g.goFile([]string{}, Client{ClientHandlers: hs}), nil
}

var tmClient = InitTemplate("Client", `
type Client struct {
	baseURL    string
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

func NewClient(baseURL string, httpClient HTTPClient) *Client {
	return &Client{baseURL: baseURL, HTTPClient: httpClient}
}

{{ range $_, $h := .ClientHandlers }}
{{ exec $h }}
{{ end -}}

func decodeJSON(r io.Reader, v interface{}, name string) {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		LogError(fmt.Errorf("decode json response %q: %w", name, err))
	}
}
`)

func (c Client) Execute() (string, error) { return tmClient.Execute(c) }

type ClientHandler struct {
	Name   string
	Path   string
	Method string

	PathParameters  []Parameter[specification.PathParameter]
	QueryParameters []Parameter[specification.QueryParameter]
	Headers         []Parameter[specification.HeaderParameter]

	DefaultResponse *Response
	Responses       []Response
}

func NewClientHandler(o *Operation) ClientHandler {
	h := ClientHandler{
		Name:   o.Name,
		Path:   string(o.Operation.PathItem.Path),
		Method: o.Operation.Method,

		PathParameters:  o.PathParameters,
		QueryParameters: o.QueryParameters,
		Headers:         o.HeaderParameters,
	}

	if o.Operation.DefaultResponse != nil {
		dr := NewResponse(&h, o, *o.Operation.DefaultResponse)
		h.DefaultResponse = &dr
	}
	for _, response := range o.Operation.Responses {
		h.Responses = append(h.Responses, NewResponse(&h, o, response))
	}

	// for _, header := range o.Handler.Parameters.Headers {
	// 	h.Headers = append(h.Headers, Param{
	// 		// Field:  header.Field,
	// 		// Parser: header.Parser,
	// 	})
	// }
	return h
}

var tmClientHandler = InitTemplate("ClientHandler", `
{{- $handler := . -}}
func (c *Client) {{.Name}}(ctx context.Context, request {{.Name}}Request) ({{ .Name }}Response, error) {
	var requestURL = c.baseURL + "{{.Path}}"
	{{ if .QueryParameters }}
	{
		var q = url.Values{}

		{{ range $_, $q := .QueryParameters -}}
		//q["{{ $q.Spec.Name }}"] = request.Query.{{ $q.FieldName }}
		{{ end }}

		requestURL += "?" + q.Encode()
	}
	{{ end }}

	req, err := http.NewRequestWithContext(ctx, http.Method{{.Method}}, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	{{ range $_, $h := .Headers }}
	{{ if not $h.Spec.Required }}if request.Headers.{{ $h.FieldName }} != nil {
	{{ end -}}
	req.Header.Set("{{$h.Spec.Name}}", {{ if not $h.Spec.Required }}*{{ end }}request.Headers.{{ $h.FieldName }})
	{{- if not $h.Spec.Required }}
	} {{ end -}}
	{{ end }}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	{{ range $_, $response := .Responses -}}
	case {{ $response.StatusCode }}:
		var response {{ $response.PublicTypeName }}
		{{ range $_, $h := $response.Headers }}
		response.Headers.{{ $h.FieldName }} = resp.Header.Get("{{ $h.Key }}")
		{{ end }}
		{{ if $response.HasBody -}}
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode '{{ $response.PublicTypeName }}' response body: %w", err)
		}
		{{- end }}
		return response, nil
	{{ end }}
	{{ if .DefaultResponse }}
	{{ $response := .DefaultResponse }}
	default:
		var response {{ $response.PublicTypeName }}
		response.Code = resp.StatusCode
		{{ range $_, $h := $response.Headers }}
		response.Headers.{{ $h.FieldName }} = resp.Header.Get("{{ $h.Key }}")
		{{ end }}
		{{ if $response.HasBody -}}
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode '{{ $response.PublicTypeName }}' response body: %w", err)
		}
		{{- end }}
		return response, nil
	{{ else }}
	default:
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	{{ end }}
	}
}
`)

func (c ClientHandler) Execute() (string, error) { return tmClientHandler.Execute(c) }
