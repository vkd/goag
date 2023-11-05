package generator

import (
	"strings"
	"unicode"

	"github.com/vkd/goag/specification"
)

type PathItem struct {
	PathItem *specification.PathItem

	Operations []*Operation
}

type Operation struct {
	Operation *specification.Operation

	Name            string
	HandlerTypeName string
	Handler         *Handler
}

func NewOperation(operation *specification.Operation) *Operation {
	o := &Operation{
		Operation: operation,
	}
	o.Name = OperationName(operation.Path, operation.Method)
	o.HandlerTypeName = o.Name + "HandlerFunc"
	return o
}

func OperationName(path specification.Path, method string) string {
	var suffix string
	if strings.HasSuffix(string(path), "/") {
		// "/shops" and "/shops/" need to have separate handlers
		suffix = "RT" // RT = root
	}
	return method + path.Name(PrefixTitle, "") + suffix
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
