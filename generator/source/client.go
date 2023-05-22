package source

type ClientFile struct {
	Funcs []Templater
}

var tmClientFile = InitTemplate("ClientFile", `
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

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

{{- range $i, $f := .Funcs }}{{ if $i }}
{{ end }}
{{- $f.String }}
{{- end }}
`)

func (c *ClientFile) String() (string, error) { return tmClientFile.String(c) }

type ClientFunc struct {
	Name            string
	Method          string
	URL             string
	Queries         []ClientFuncQuery
	Responses       []ClientFuncResponse
	DefaultResponse ClientFuncResponse
}

var tmClientFunc = InitTemplate("ClientFunc", `
func (c *Client) {{.Name}}JSON(request {{.Name}}Request) (interface{}, error) {
	var requestURL = c.baseURL + "{{.URL}}"
	var q = url.Values{}
	{{ range $i, $q := .Queries }}{{ if $i }}
	{{ end -}}
	if request.{{$q.RequestFieldName}} != nil {
		q.Add("{{$q.QueryName}}", {{$q.Formatter.String}})
	}
	{{ end }}
	requestURL += "?" + q.Encode()

	req, err := http.NewRequest("{{.Method}}", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	{{ range $_, $r := .Responses }}
	case {{$r.StatusCode}}:
		var response {{$r.ResponseType}}
		{{ if $r.HasBody }}
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode {{$r.StatusCode}} response: %w", err)
		}
		{{ end }}
		return response, nil
	{{ end }}
	{{ if .DefaultResponse.ResponseType }}
	default:
		var response {{.DefaultResponse.ResponseType}}
		response.Code = resp.StatusCode
		{{ if .DefaultResponse.HasBody }}
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode default response (code: %d): %w", resp.StatusCode, err)
		}
		{{ end }}
		return response, nil
	{{ else }}
	default:
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	{{ end }}
	}
}
`)

func (c *ClientFunc) String() (string, error) { return tmClientFunc.String(c) }

type ClientFuncQuery struct {
	QueryName        string
	RequestFieldName string
	Formatter        Templater
}

type ClientFuncResponse struct {
	StatusCode   int
	HasBody      bool
	ResponseType string
}
