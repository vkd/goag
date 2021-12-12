package source

import (
	"text/template"
)

type Assign struct {
	From Render
	To   string
}

var tmAssign = template.Must(template.New("Assign").Parse(`{{.To}} = {{.From.String}}`))

func (c Assign) String() (string, error) { return String(tmAssign, c) }

type AssignNew struct {
	From Render
	To   string
}

var tmAssignNew = template.Must(template.New("AssignNew").Parse(`{{.To}} := {{.From.String}}`))

func (c AssignNew) String() (string, error) { return String(tmAssignNew, c) }
