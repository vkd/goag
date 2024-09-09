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
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, ".", "_")
	ss := strings.Split(s, "_")
	if len(ss) > 1 {
		for i, v := range ss {
			ss[i] = Title(v)
		}
		return strings.Join(ss, "")
	}
	switch {
	case strings.HasSuffix(s, "id"):
		return stringsTitle(strings.TrimSuffix(s, "id")) + "ID"
	case strings.HasSuffix(s, "ids"):
		return stringsTitle(strings.TrimSuffix(s, "ids")) + "IDs"
	}
	return stringsTitle(s)
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
