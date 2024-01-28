package source

type GoVariable string

type GoStringVar GoVariable

func (v GoStringVar) Parser(to string, from Render, errWr ErrorWrapper) Renders {
	return Renders{
		AssignNew(to, from),
	}
}

type GoIntXXVarFormat struct {
	XX int
	V  string
}

var tmGoIntXXVarFormat = InitTemplate("GoIntXXVarFormat", `fmt.Sprintf("%d", {{.}})`)

func (v GoIntXXVarFormat) String() (string, error) { return tmGoIntXXVarFormat.String(v.V) }
