package generator

import "github.com/vkd/goag/specification"

type GoType interface {
	ToStringSlice() Templater
}

func NewGoType(schema specification.Schema) GoType {
	switch schema.Schema.Type {
	case "integer":
		switch schema.Schema.Format {
		// case "int32":
		// 	return Int32Type{}
		}
		panic("not implemented")
	}
	panic("not implemented")
}

// func NewGoType(schema *specification.Schema)

type Int32Type struct{}

func (Int32Type) Variable(from Templater) Int32Variable {
	return Int32Variable{Var: from}
}

type Variable struct {
	Var Templater
}

func (v Variable) Execute() (string, error) {
	return v.Var.String()
}

func (v Variable) String() (string, error) {
	return v.Var.String()
}

type Int32Variable Variable

func (v Int32Variable) ToInt64() Int64Variable {
	return Int64Variable{Var: Int32ToInt64(v.Var)}
}

type Int64Variable Variable

func (v Int64Variable) ToString() StringVariable {
	return StringVariable{Var: Int64ToString(v.Var)}
}

type StringVariable Variable

func (v StringVariable) ToStringSlice() StringSliceVariable {
	return StringSliceVariable{Var: StringToStringSlice(v.Var)}
}

type StringSliceVariable Variable

// func (i Int32Variable) FormatToString(valueFrom Templater) Templater {
// 	return Int64ToString(Int32ToInt64(valueFrom))
// }

var tmInt32ToInt64 = InitTemplate("Int32ToInt64", `int64({{ exec . }})`)

func Int32ToInt64(from Templater) Templater {
	return TemplateData(tmInt32ToInt64, from)
}

var tmInt64ToString = InitTemplate("Int64ToString", `strconv.FormatInt({{ exec . }}, 10)`)

func Int64ToString(from Templater) Templater {
	return TemplateData(tmInt64ToString, from)
}

var tmStringToStringSlice = InitTemplate("StringToStringSlice", `[]string{ {{ exec . }} }`)

func StringToStringSlice(from Templater) Templater {
	return TemplateData(tmStringToStringSlice, from)
}
