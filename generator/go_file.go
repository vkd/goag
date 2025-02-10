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

type Import struct {
	Value string
	Alias string
}

func NewImport(v, alias string) Import {
	return Import{
		Value: v,
		Alias: alias,
	}
}

type Imports []Import

func NewImportsS(ss ...string) Imports {
	out := make(Imports, 0, len(ss))
	for _, s := range ss {
		if s != "" {
			out = out.AppendS(s)
		}
	}
	return out
}

func (i Imports) AppendS(s string) Imports {
	return append(i, NewImport(s, ""))
}
