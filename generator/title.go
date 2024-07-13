package generator

import (
	"strings"
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
		return strings.Title(strings.TrimSuffix(lower, "id")) + "ID"
	case strings.HasSuffix(lower, "ids"):
		return strings.Title(strings.TrimSuffix(lower, "ids")) + "IDs"
	}
	return strings.Title(strings.ToLower(s))
}

func PrivateFieldName(name string) string {
	if len(name) == 0 {
		return ""
	}
	return strings.ToLower(name[:1]) + name[1:]
}
