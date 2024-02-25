package generator

type GoFile struct {
	SkipDoNotEdit bool
	PackageName   string
	Imports       []Import
	Body          any
}

func (g GoFile) Render() (string, error) {
	return ExecuteTemplate("GoFile", g)
}

type Import string
