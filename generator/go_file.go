package generator

type GoFile struct {
	SkipDoNotEdit bool
	PackageName   string
	Imports       []string
	Body          Render
}

func (g GoFile) Render() (string, error) {
	return ExecuteTemplate("GoFile", g)
}
