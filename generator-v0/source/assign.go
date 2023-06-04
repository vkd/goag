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

func AssignNew(to string, from Render) Render {
	return assignNew{to, from}
}

var tmAssignNew = InitTemplate("AssignNew", `{{.To}} := {{.From.String}}`)

type assignNew struct {
	To   string
	From Render
}

func (c assignNew) String() (string, error) { return tmAssignNew.String(c) }
