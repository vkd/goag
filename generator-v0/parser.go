package generator

import "github.com/vkd/goag/generator-v0/source"

type StringsParser interface {
	StringsParser(from, to string, _ ErrorWrapper) Render
}

func NewStringsParser(s SchemaRender, from, toOrig string, isPointer bool, mkErr ErrorWrapper) Render {

	switch s := s.(type) {
	case StringsParser:
		return s.StringsParser(from, toOrig, mkErr)
	}

	var conv Render
	to := "v"
	conv = s.Parser(from+"[0]", to, mkErr)

	_, optionable := s.(interface{ Optionable() })

	if isPointer && !optionable {
		to = "&" + to
	}
	conv = source.Renders{conv, source.Assign(toOrig, GoValue(to))}

	return conv
}
