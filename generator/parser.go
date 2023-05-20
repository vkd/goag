package generator

import "github.com/vkd/goag/generator/source"

type StringsParser interface {
	StringsParser(from, to string, _ ErrorWrapper) Render
}

func NewStringsParser(s SchemaRender, from, toOrig string, isPointer bool, mkErr ErrorWrapper) Render {
	to := toOrig

	var conv Render
	if sp, ok := s.(StringsParser); ok {
		conv = sp.StringsParser(from, to, mkErr)
	} else {
		to := "v"
		conv = s.Parser(from+"[0]", to, mkErr)

		_, optionable := s.(interface{ Optionable() })

		if isPointer && !optionable {
			to = "&" + to
		}
		conv = source.Renders{conv, Assign{GoValue(to), toOrig}}
	}

	return conv
}
