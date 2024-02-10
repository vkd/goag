package generator

import (
	"strings"
	"unicode"
)

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
			isWord = unicode.IsLetter(r) || (isWord && unicode.IsDigit(r)) // if start a new word
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
