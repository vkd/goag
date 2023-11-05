package generator

import (
	"strings"

	"github.com/vkd/goag/specification"
)

func Title(s string) string {
	switch s {
	case "id", "Id":
		return "ID"
	case "ids":
		return "IDs"
	}
	ss := strings.Split(s, "_")
	if len(ss) > 1 {
		for i, v := range ss {
			ss[i] = Title(v)
		}
		return strings.Join(ss, "")
	}
	return strings.Title(strings.ToLower(s))
}

func PrefixTitle(s specification.Prefix) string {
	return Title(s.Name())
}
