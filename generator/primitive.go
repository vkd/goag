package generator

import (
	"strconv"
)

type Primitive struct {
	PrimitiveIface
}

type PrimitiveIface interface {
	GoType() string

	// ParseBaseType(to, from string, isNew bool, mkErr ErrorRender)
	ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error)

	RenderToBaseType(from string) (string, error)
	RenderFormat(from string) (string, error)
}

func NewPrimitive(v PrimitiveIface) Primitive {
	return Primitive{
		PrimitiveIface: v,
	}
}

func (t Primitive) Kind() SchemaKind { return SchemaKindPrimitive }

func (t Primitive) RenderGoType() (string, error) {
	return t.GoType(), nil
}

func (t Primitive) FuncTypeName() string {
	return Title(t.GoType())
}

func (t Primitive) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return t.PrimitiveIface.ParseString(to, from+"[0]", isNew, mkErr)
}

func (t Primitive) RenderToBaseType(to, from string) (string, error) {
	out, err := t.PrimitiveIface.RenderToBaseType(from)
	return to + " = " + out, err
}

func (t Primitive) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("Primitive_RenderFormatStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,

		"RenderFormat": t.PrimitiveIface.RenderFormat,
	})
}

type BoolType struct{}

var _ PrimitiveIface = BoolType{}

func (b BoolType) GoType() string { return "bool" }

func (b BoolType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Bool_ParseString", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (_ BoolType) RenderToBaseType(from string) (string, error) {
	return from, nil
}

func (b BoolType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("Bool_RenderFormat", TData{
		"From": from,
	})
}

type IntType struct {
	BitSize int
}

var _ PrimitiveIface = IntType{}

func (i IntType) GoType() string {
	switch i.BitSize {
	case 0:
		return "int"
	case 32:
		return "int32"
	case 64:
		return "int64"
	default:
		return "int" + strconv.Itoa(i.BitSize)
	}
}

func (i IntType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Int64_Parser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	default:
		return ExecuteTemplate("IntX_Parser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"GoType":  i.GoType(),
			"BitSize": i.BitSize,
		})
	}
}

func (_ IntType) RenderToBaseType(from string) (string, error) {
	return from, nil
}

func (i IntType) RenderFormat(from string) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Int64_Format", TData{
			"From": from,
		})
	default:
		return ExecuteTemplate("IntX_Format", TData{
			"From": from,
		})
	}
}

type FloatType struct {
	BitSize int
}

var _ PrimitiveIface = FloatType{}

func (f FloatType) GoType() string {
	switch f.BitSize {
	case 32:
		return "float32"
	case 64:
		return "float64"
	default:
		return "float" + strconv.Itoa(f.BitSize)
	}
}

func (i FloatType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Float64_Parser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	default:
		return ExecuteTemplate("FloatX_Parser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"GoType":  i.GoType(),
			"BitSize": i.BitSize,
		})
	}
}

func (_ FloatType) RenderToBaseType(from string) (string, error) {
	return from, nil
}

func (i FloatType) RenderFormat(from string) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Float64_Format", TData{
			"From": from,
		})
	default:
		return ExecuteTemplate("FloatX_Format", TData{
			"From": from,

			"BitSize": i.BitSize,
		})
	}
}

type StringType struct{}

var _ PrimitiveIface = StringType{}

func (s StringType) GoType() string { return "string" }

func (_ StringType) RenderToBaseType(from string) (string, error) {
	return from, nil
}

func (_ StringType) RenderFormat(from string) (string, error) {
	return from, nil
}

func (_ StringType) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if isNew {
		return to + " := " + from, nil
	}
	return to + " = " + from, nil
}

type DateTime struct {
	GoLayout   string
	TextLayout string
}

var _ PrimitiveIface = DateTime{}

func (DateTime) GoType() string { return "time.Time" }

func (d DateTime) RenderToBaseType(from string) (string, error) {
	return d.RenderFormat(from)
}

func (d DateTime) RenderFormat(from string) (string, error) {
	layout := "time.RFC3339"
	if d.GoLayout != "" {
		layout = d.GoLayout
	}
	if d.TextLayout != "" {
		layout = `"` + d.TextLayout + `"`
	}
	return ExecuteTemplate("DateTime_RenderFormat", TData{
		"From": from,

		"Layout": layout,
	})
}

func (d DateTime) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	layout := "time.RFC3339"
	if d.GoLayout != "" {
		layout = d.GoLayout
	}
	if d.TextLayout != "" {
		layout = `"` + d.TextLayout + `"`
	}
	return ExecuteTemplate("DateTime_ParseString", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Layout": layout,
	})
}
