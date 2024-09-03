package generator

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	lower := strings.ToLower(s)
	switch {
	case strings.HasSuffix(lower, "id"):
		return stringsTitle(strings.TrimSuffix(lower, "id")) + "ID"
	case strings.HasSuffix(lower, "ids"):
		return stringsTitle(strings.TrimSuffix(lower, "ids")) + "IDs"
	}
	return stringsTitle(strings.ToLower(s))
}

func PrivateFieldName(name string) string {
	if len(name) == 0 {
		return ""
	}
	return strings.ToLower(name[:1]) + name[1:]
}

var undCaser = cases.Title(language.Und, cases.NoLower)

func stringsTitle(s string) string {
	return undCaser.String(s)
}
