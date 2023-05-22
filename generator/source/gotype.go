package source

import "fmt"

type GoVariable string

type GoStringVar GoVariable

func (v GoStringVar) Parser(to string, from Render, errWr ErrorWrapper) Renders {
	return Renders{
		AssignNew(to, from),
	}
}

type GoIntVar struct{}

func (v GoIntVar) Parser(to string, from string, errWr ErrorWrapper) Renders {
	return Renders{
		ParseIntXX{0, from, "vInt", errWr},
		AssignNew(to, TypeConversion("int", "vInt")),
	}
}

type GoIntXXVar int

func (v GoIntXXVar) Parser(from, to string, errWr ErrorWrapper) Templater {
	return Renders{
		ParseIntXX{int(v), from, "vInt", errWr},
		AssignNew(to, TypeConversion(fmt.Sprintf("int%d", v), "vInt")),
	}
}

func (v GoIntXXVar) Format(s string) Templater { return GoIntXXVarFormat{int(v), s} }

func (v GoIntXXVar) String() (string, error) { return fmt.Sprintf("int%d", v), nil }

type GoIntXXVarFormat struct {
	XX int
	V  string
}

var tmGoIntXXVarFormat = InitTemplate("GoIntXXVarFormat", `fmt.Sprintf("%d", {{.}})`)

func (v GoIntXXVarFormat) String() (string, error) { return tmGoIntXXVarFormat.String(v.V) }
