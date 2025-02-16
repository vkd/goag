package generator

import (
	"strconv"
)

type Primitive struct {
	PrimitiveIface
}

type PrimitiveIface interface {
	GoType() string
	RenderToBaseTypeInline(from string) (string, error)

	RenderToStringInline(from string) (string, error)
	RenderStringParser(to, from string, isNew bool, mkErr ErrorRender) (string, error)

	RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
	RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
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

func (t Primitive) RenderToBaseType(to, from string) (string, error) {
	out, err := t.PrimitiveIface.RenderToBaseTypeInline(from)
	return to + " = " + out, err
}

func (t Primitive) RenderFormat(from string) (string, error) {
	return t.PrimitiveIface.RenderToStringInline(from)
}

func (t Primitive) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("Primitive_RenderFormatStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,

		"RenderFormat": t.PrimitiveIface.RenderToStringInline,
	})
}

func (t Primitive) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return t.PrimitiveIface.RenderStringParser(to, from, isNew, mkErr)
}

func (t Primitive) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Primitive_ParseStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"GoType":             t.PrimitiveIface.GoType(),
		"RenderStringParser": t.PrimitiveIface.RenderStringParser,
	})
}

func (t Primitive) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return t.PrimitiveIface.RenderUnmarshalJSON(to, from, isNew, mkErr)
}

func (t Primitive) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return t.PrimitiveIface.RenderMarshalJSON(to, from, isNew, mkErr)
}

type BoolType struct{}

func NewBoolType() Primitive { return NewPrimitive(BoolType{}) }

var _ PrimitiveIface = BoolType{}

func (BoolType) GoType() string { return "bool" }

func (BoolType) RenderToBaseTypeInline(from string) (string, error) {
	return from, nil
}

func (b BoolType) RenderStringParser(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Bool_RenderStringParser", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (b BoolType) RenderToStringInline(from string) (string, error) {
	return ExecuteTemplate("Bool_RenderToStringInline", TData{
		"From": from,
	})
}

func (_ BoolType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Bool_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (_ BoolType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Primitive_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

type IntType struct {
	BitSize int
}

func NewIntType() Primitive { return NewPrimitive(IntType{BitSize: 0}) }

func NewIntTypeXX(bitSize int) Primitive { return NewPrimitive(IntType{BitSize: bitSize}) }

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

func (IntType) RenderToBaseTypeInline(from string) (string, error) {
	return from, nil
}

func (i IntType) RenderStringParser(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Int64_RenderStringParser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	default:
		return ExecuteTemplate("IntX_RenderStringParser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"GoType":  i.GoType(),
			"BitSize": i.BitSize,
		})
	}
}

func (i IntType) RenderToStringInline(from string) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Int64_RenderToStringInline", TData{
			"From": from,
		})
	default:
		return ExecuteTemplate("IntX_RenderToStringInline", TData{
			"From": from,
		})
	}
}

func (i IntType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Int64_RenderUnmarshalJSON", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	default:
		return ExecuteTemplate("IntX_RenderUnmarshalJSON", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"GoType": i.GoType(),
		})
	}
}

func (i IntType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Primitive_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

type FloatType struct {
	BitSize int
}

func NewFloatType(bitSize int) Primitive { return NewPrimitive(FloatType{BitSize: bitSize}) }

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

func (FloatType) RenderToBaseTypeInline(from string) (string, error) {
	return from, nil
}

func (i FloatType) RenderStringParser(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Float64_RenderStringParser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	default:
		return ExecuteTemplate("FloatX_RenderStringParser", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"GoType":  i.GoType(),
			"BitSize": i.BitSize,
		})
	}
}

func (i FloatType) RenderToStringInline(from string) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Float64_RenderToStringInline", TData{
			"From": from,
		})
	default:
		return ExecuteTemplate("FloatX_RenderToStringInline", TData{
			"From": from,

			"BitSize": i.BitSize,
		})
	}
}

func (i FloatType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 64:
		return ExecuteTemplate("Float64_RenderUnmarshalJSON", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	default:
		return ExecuteTemplate("FloatX_RenderUnmarshalJSON", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"GoType": i.GoType(),
		})
	}
}

func (i FloatType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Primitive_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

type StringType struct{}

func NewStringType() Primitive { return NewPrimitive(StringType{}) }

var _ PrimitiveIface = StringType{}

func (StringType) GoType() string { return "string" }

func (StringType) RenderToBaseTypeInline(from string) (string, error) {
	return from, nil
}

func (_ StringType) RenderToStringInline(from string) (string, error) {
	return from, nil
}

func (_ StringType) RenderStringParser(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if isNew {
		return to + " := " + from, nil
	}
	return to + " = " + from, nil
}

func (_ StringType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("String_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (_ StringType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("Primitive_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

type DateTime struct {
	GoLayout   string
	TextLayout string
}

func NewDateTime(layout string) (DateTime, Imports) {
	return DateTime{GoLayout: layout}, Imports{NewImport("time", "")}
}

var _ PrimitiveIface = DateTime{}

func (DateTime) GoType() string { return "time.Time" }

func (d DateTime) RenderToBaseTypeInline(from string) (string, error) {
	return d.RenderToStringInline(from)
}

func (d DateTime) RenderToStringInline(from string) (string, error) {
	layout := "time.RFC3339"
	if d.GoLayout != "" {
		layout = d.GoLayout
	}
	if d.TextLayout != "" {
		layout = `"` + d.TextLayout + `"`
	}
	return ExecuteTemplate("DateTime_RenderToStringInline", TData{
		"From": from,

		"Layout": layout,
	})
}

func (d DateTime) RenderStringParser(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	layout := "time.RFC3339"
	if d.GoLayout != "" {
		layout = d.GoLayout
	}
	if d.TextLayout != "" {
		layout = `"` + d.TextLayout + `"`
	}
	return ExecuteTemplate("DateTime_RenderStringParser", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Layout": layout,
	})
}

func (d DateTime) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	layout := "time.RFC3339"
	if d.GoLayout != "" {
		layout = d.GoLayout
	}
	if d.TextLayout != "" {
		layout = `"` + d.TextLayout + `"`
	}
	return ExecuteTemplate("DateTime_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Layout": layout,
	})
}

func (d DateTime) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	layout := "time.RFC3339"
	if d.GoLayout != "" {
		layout = d.GoLayout
	}
	if d.TextLayout != "" {
		layout = `"` + d.TextLayout + `"`
	}
	return ExecuteTemplate("DateTime_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Layout": layout,
	})
}
