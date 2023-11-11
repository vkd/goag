package generator

import (
	"github.com/vkd/goag/generator"
)

var PublicFieldName = generator.PublicFieldName
var PrivateFieldName = generator.PrivateFieldName

// func PublicFieldName(name string) string {
// 	if strings.HasSuffix(name, "Id") {
// 		name = name[:len(name)-2] + "ID"
// 	}

// 	name = strings.ReplaceAll(name, "-", "|")
// 	name = strings.ReplaceAll(name, "_", "|")
// 	names := strings.Split(name, "|")

// 	for i, name := range names {
// 		switch name {
// 		case "id", "Id":
// 			names[i] = "ID"
// 		case "ids":
// 			names[i] = "IDs"
// 		default:
// 			names[i] = strings.Title(name)
// 		}
// 	}
// 	return strings.Join(names, "")
// }
