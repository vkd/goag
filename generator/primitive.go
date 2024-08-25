package generator

import (
	"strconv"
)

type Primitive struct {
	SingleValue

	PrimitiveIface
}

type PrimitiveIface interface {
	GoType() string

	ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error)

	RenderFormat(from string) (string, error)
}

func NewPrimitive(v PrimitiveIface) Primitive {
	return Primitive{
		PrimitiveIface: v,
	}
}

func (t Primitive) Kind() SchemaKind { return SchemaKindPrimitive }

func (t Primitive) Render() (string, error) {
	return t.GoType(), nil
}

func (t Primitive) FuncTypeName() string {
	return stringsTitle(t.GoType())
}

func (t Primitive) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return t.PrimitiveIface.ParseString(to, from+"[0]", isNew, mkErr)
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

func (_ StringType) RenderFormat(from string) (string, error) {
	return from, nil
}

func (_ StringType) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if isNew {
		return to + " := " + from, nil
	}
	return to + " = " + from, nil
}
