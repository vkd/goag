package source

func Assign(to string, from Render) Render {
	return assign{to, from}
}

var tmAssign = InitTemplate("Assign", `{{.To}} = {{.From.String}}`)

type assign struct {
	To   string
	From Render
}

func (c assign) String() (string, error) { return tmAssign.String(c) }
