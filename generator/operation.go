package generator

import (
	"strings"
	"unicode"

	"github.com/vkd/goag/specification"
)

type Operation struct {
	Operation *specification.Operation

	Name            string
	HandlerTypeName string

	PathParameters   []Parameter[*specification.PathParameter]
	QueryParameters  []Parameter[specification.QueryParameter]
	HeaderParameters []Parameter[*specification.HeaderParameter]

	Handler *HandlerOld // Deprecated
}

func NewOperation(operation *specification.Operation) *Operation {
	o := &Operation{
		Operation: operation,
	}
	o.Name = OperationName(operation.PathItem.Path, operation.Method)
	o.HandlerTypeName = o.Name + "HandlerFunc"

	for _, pathParam := range operation.Parameters.Path {
		o.PathParameters = append(o.PathParameters, Parameter[*specification.PathParameter]{
			Spec: pathParam,

			FieldName: PublicFieldName(pathParam.Name),
		})
	}
	for _, query := range operation.Parameters.Query {
		o.QueryParameters = append(o.QueryParameters, Parameter[specification.QueryParameter]{
			Spec: query,

			FieldName: PublicFieldName(query.Name),
		})
	}
	for _, header := range operation.Parameters.Headers {
		o.HeaderParameters = append(o.HeaderParameters, Parameter[*specification.HeaderParameter]{
			Spec: header,

			FieldName: PublicFieldName(header.Name),
		})
	}

	return o
}

func OperationName(path specification.Path, method string) string {
	var out string
	for _, dir := range path.Dirs {
		out += PrefixTitle(dir.Raw)
	}

	var suffix string
	if len(path.Dirs) > 1 && path.Dirs[len(path.Dirs)-1].Raw == "/" {
		suffix = "RT"
	}

	return method + out + suffix
}

func PublicFieldName(name string) string {
	runes := ([]rune)(name)

	names := make([]string, 0)

	var isWord bool
	var li, ri int
	for i, r := range runes {
		isSplit := !unicode.IsLetter(r) || unicode.IsUpper(r)
		if isSplit { // needs to split here
			if isWord { // if currently is building a word
				names = append(names, string(runes[li:ri+1]))
			}
			isWord = unicode.IsLetter(r) // if start a new word
			if isWord {
				li = i
				ri = i
			}
			continue
		}
		if !isWord {
			isWord = true
			li = i
		}
		ri = i
	}
	if isWord {
		names = append(names, string(runes[li:ri+1]))
	}

	for i, name := range names {
		switch name {
		case "id", "Id":
			names[i] = "ID"
		case "ids":
			names[i] = "IDs"
		default:
			names[i] = strings.Title(name)
		}
	}
	return strings.Join(names, "")
}

type Parameter[T interface {
	*specification.PathParameter | specification.QueryParameter | *specification.HeaderParameter
}] struct {
	Spec T

	FieldName string
	// Type      GoType
}

func PrivateFieldName(name string) string {
	if len(name) == 0 {
		return ""
	}
	return strings.ToLower(name[:1]) + name[1:]
}
