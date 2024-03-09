package generator

type GoFile struct {
	SkipDoNotEdit bool
	PackageName   string
	Imports       Imports
	Body          Render
}

func (g GoFile) Render() (string, error) {
	return ExecuteTemplate("GoFile", g)
}

type Import string

type Imports []Import

func NewImportsS(ss ...string) Imports {
	out := make(Imports, 0, len(ss))
	for _, s := range ss {
		out = append(out, Import(s))
	}
	return out
}

func (i Imports) AppendS(s string) Imports { return append(i, Import(s)) }
