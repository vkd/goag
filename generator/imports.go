package generator

type Imports []Import

func NewImportsS(ss ...string) Imports {
	out := make(Imports, 0, len(ss))
	for _, s := range ss {
		out = append(out, Import(s))
	}
	return out
}

func (i Imports) AppendS(s string) Imports { return append(i, Import(s)) }
