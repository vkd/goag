package generator

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type Response struct {
	Name        string
	PrivateName string
	HandlerName string

	Description string
	Status      string

	ResponserInterfaceName string

	IsBody  bool
	Body    Render
	Headers []ResponseHeader

	Struct GoTypeDef

	Args []ResponseArg
}

const txtResponse = `{{$response := . -}}
func {{.Name}}({{range $i,$a := .Args}}{{if $i}}, {{end}}{{$a.ArgName}} {{$a.Type.String}}{{end}}) {{.HandlerName}}Responser {
	var out {{.PrivateName}}
	{{- range $_, $a := .Args}}
	out.{{if .IsHeader}}Headers.{{end}}{{.FieldName}} = {{.ArgName}}
	{{- end}}
	return out
}

{{- if .Body }}

{{ .Body.String }}
{{- end }}

{{.Struct.String}}

func (r {{.PrivateName}}) {{.ResponserInterfaceName}}(w http.ResponseWriter) {
	{{if .IsDefault}}w.WriteHeader(r.Code){{else}}w.WriteHeader({{.Status}}){{end}}
	{{range $_, $h := .Headers}}w.Header().Set("{{$h.Name}}", r.Headers.{{$h.FieldName}})
	{{end -}}
	{{if .IsBody}}writeJSON(w, r.Body, "{{.Name}}")
	{{end -}}
}`

var tmRespose = template.Must(template.New("Response").Parse(txtResponse))

func (r Response) String() (string, error) {
	out, err := String(tmRespose, r)
	if err != nil {
		return "", fmt.Errorf("Response %q: %w", r.Name, err)
	}
	return out, nil
}

func NewResponse(s *openapi3.SchemaRef, handlerName string, responserName string, status, contentType string, resp *openapi3.Response, gap string, code string, response *openapi3.ResponseRef) Response {
	var out Response
	out.HandlerName = handlerName
	out.Name = handlerName + "Response" + strings.Title(status)
	switch contentType {
	case "application/json":
		out.Name += "JSON"
	default:
		out.Name += contentType
	}
	out.PrivateName = PrivateFieldName(out.Name)
	out.Status = status

	var fields []GoStructField
	if out.IsDefault() {
		fields = append(fields, GoStructField{
			Name: "Code",
			Type: Int,
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Code",
			ArgName:   "code",
			Type:      Int,
		})
	}

	if response.Value != nil && response.Value.Description != nil {
		out.Description = *response.Value.Description
	}
	if resp.Description != nil {
		out.Description = *resp.Description
	}

	out.ResponserInterfaceName = responserName

	if s != nil {
		out.IsBody = true

		sr := NewSchemaRef(s)

		fieldType := sr

		if s.Ref == "" {
			switch sr := sr.(type) {
			case GoStruct:
				bodyType := GoTypeDef{
					Name: out.Name + "Body",
					Type: sr,
				}
				if s.Value.AdditionalProperties != nil {
					bodyType.Methods = append(bodyType.Methods,
						GoVarDef{Name: "_", Type: GoType("json.Marshaler"), Value: GoValue("(*" + bodyType.Name + ")(nil)")},
						MarshalJSONFunc(bodyType.GoTypeRef(), sr),
					)
				}
				out.Body = bodyType
				fieldType = GoType(bodyType.Name)
			}
		}

		fields = append(fields, GoStructField{
			Name: "Body",
			Type: fieldType,
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: "Body",
			ArgName:   "body",
			Type:      fieldType,
		})
	}

	pathHeaders := PathHeaders(resp.Headers)
	fieldHeaders := make([]GoStructField, 0, len(pathHeaders))
	for _, h := range pathHeaders {
		sr := NewSchemaRef(h.Header.Value.Schema)

		header := ResponseHeader{
			Name:      h.Name,
			FieldName: PublicFieldName(h.Name),
			Type:      sr,
		}

		out.Headers = append(out.Headers, header)

		fieldHeaders = append(fields, GoStructField{
			Name: header.FieldName,
			Type: header.Type,
		})
		out.Args = append(out.Args, ResponseArg{
			FieldName: header.FieldName,
			ArgName:   PrivateFieldName(header.FieldName),
			IsHeader:  true,
			Type:      header.Type,
		})
	}
	if len(fieldHeaders) > 0 {
		fields = append(fields, GoStructField{
			Name: "Headers",
			Type: GoStruct{
				Fields: fieldHeaders,
			},
		})
	}

	out.Struct = GoTypeDef{
		Comment: out.Description,
		Name:    out.PrivateName,
		Type: GoStruct{
			Fields: fields,
		},
	}

	return out
}

func (r Response) IsDefault() bool { return strings.EqualFold(r.Status, "default") }

type ResponseArg struct {
	FieldName string
	ArgName   string
	IsHeader  bool
	Type      Render
}
