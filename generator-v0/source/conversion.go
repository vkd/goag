package source

func TypeConversion(goType, variable string) Render {
	return typeConversion{goType, variable}
}

var tmTypeConversion = InitTemplate("TypeConversion", `{{.GoType}}({{.Var}})`)

type typeConversion struct {
	GoType string
	Var    string
}

func (c typeConversion) String() (string, error) { return tmTypeConversion.String(c) }
