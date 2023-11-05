package generator

type GoStructField struct {
	Field string
	Type  string
}

var tmGoStructField = InitTemplate("GoStructField", `{{ .Field }} {{ .Type }}`)

func (s GoStructField) String() (string, error) { return tmGoStructField.Execute(s) }
