package generator

import "strings"

func PublicFieldName(name string) string {
	if strings.HasSuffix(name, "Id") {
		name = name[:len(name)-2] + "ID"
	}

	name = strings.ReplaceAll(name, "-", "|")
	name = strings.ReplaceAll(name, "_", "|")
	names := strings.Split(name, "|")

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

func PrivateFieldName(name string) string {
	if len(name) == 0 {
		return ""
	}
	return strings.ToLower(name[:1]) + name[1:]
}
