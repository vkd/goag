package generator

import (
	"text/template"

	"github.com/vkd/goag/generator/source"
)

type StructParser struct {
	From, To string
	Error    ErrorWrapper
}

type ErrorWrapper = source.ErrorWrapper
type ErrorBuilder = source.ErrorBuilder
type ParseError = source.ParseError

var tmStructParser = template.Must(template.New("StructParser").Parse(`err := {{.To}}.UnmarshalText({{.From}})
if err != nil {
	return zero, {{.Error.Wrap "parse struct"}}
}`,
))

func (s StructParser) String() (string, error) { return String(tmStructParser, s) }
