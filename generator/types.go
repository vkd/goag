package generator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vkd/goag/specification"
)

type BoolType struct{}

var _ Parser = BoolType{}

func (b BoolType) Render() (string, error) { return "bool", nil }

func (b BoolType) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("BoolParser", struct {
		From  Render
		To    Render
		MkErr ErrorRender
	}{from, to, mkErr})
}

var _ Formatter = BoolType{}

func (b BoolType) RenderFormat(from Render) (string, error) {
	return ExecuteTemplate("BoolFormat", struct {
		From Render
	}{from})
}

type IntType struct {
	BitSize int
}

func (i IntType) Render() (string, error) {
	switch i.BitSize {
	case 0:
		return "int", nil
	case 32:
		return "int32", nil
	case 64:
		return "int64", nil
	default:
		return "int" + strconv.Itoa(i.BitSize), nil
	}
}

var _ Parser = IntType{}

func (i IntType) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 0:
		return ExecuteTemplate("IntParser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	case 32:
		return ExecuteTemplate("Int32Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	case 64:
		return ExecuteTemplate("Int64Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	default:
		return "", fmt.Errorf("unsupported int bit size %d", i.BitSize)
	}
}

var _ Formatter = IntType{}

func (i IntType) RenderFormat(from Render) (string, error) {
	switch i.BitSize {
	case 0:
		return ExecuteTemplate("IntFormat", struct {
			From Render
		}{from})
	case 32:
		return ExecuteTemplate("Int32Format", struct {
			From Render
		}{from})
	case 64:
		return ExecuteTemplate("Int64Format", struct {
			From Render
		}{from})
	default:
		return "", fmt.Errorf("unsupported int bit size %d", i.BitSize)
	}
}

type FloatType struct {
	BitSize int
}

func (f FloatType) Render() (string, error) {
	switch f.BitSize {
	case 32:
		return "float32", nil
	case 64:
		return "float64", nil
	default:
		return "float" + strconv.Itoa(f.BitSize), nil
	}
}

var _ Parser = FloatType{}

func (i FloatType) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 32:
		return ExecuteTemplate("Float32Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	case 64:
		return ExecuteTemplate("Float64Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	default:
		return "", fmt.Errorf("unsupported float bit size %d", i.BitSize)
	}
}

var _ Formatter = FloatType{}

func (i FloatType) RenderFormat(from Render) (string, error) {
	switch i.BitSize {
	case 32:
		return ExecuteTemplate("Float32Format", struct {
			From Render
		}{from})
	case 64:
		return ExecuteTemplate("Float64Format", struct {
			From Render
		}{from})
	default:
		return "", fmt.Errorf("unsupported float bit size %d", i.BitSize)
	}
}

type StringType struct{}

func (s StringType) Render() (string, error) { return "string", nil }

func (_ StringType) RenderFormat(from Render) (string, error) {
	return from.Render()
}

type CustomType string

func NewCustomType(s string) CustomType {
	var customImport, customType string = "", s
	slIdx := strings.LastIndex(s, "/")
	if slIdx >= 0 {
		customImport = s[:slIdx]
		customType = s[slIdx+1:]

		dotIdx := strings.LastIndex(s, ".")
		if dotIdx >= 0 {
			customImport = s[:dotIdx]
		}
	}

	AddImport(customImport)
	return CustomType(customType)
}

func (c CustomType) Render() (string, error) {
	return string(c), nil
}

func (c CustomType) RenderFormat(from Render) (string, error) {
	return ExecuteTemplate("CustomTypeFormat", struct {
		From Render
	}{From: from})
}

type SliceType struct {
	Items Render
}

func (s SliceType) Render() (string, error) { return ExecuteTemplate("SliceType", s) }

func (s SliceType) RenderFormat(from Render) (string, error) {
	return "", fmt.Errorf("SliceType doesn't implement Formatter interface")
}

type StructureType struct {
	Fields []StructureField
}

func NewStructureType(s specification.Schema) (zero StructureType, _ error) {
	var stype StructureType
	for _, p := range s.Properties {
		f, err := NewStructureField(p)
		if err != nil {
			return zero, fmt.Errorf("field %q: %w", p.Name, err)
		}
		stype.Fields = append(stype.Fields, f)
	}
	return stype, nil
}

func (s StructureType) Render() (string, error) { return ExecuteTemplate("StructureType", s) }

type StructureField struct {
	Comment string
	Name    string
	Type    Render
	Tags    []StructureFieldTag
}

func NewStructureField(s specification.SchemaProperty) (zero StructureField, _ error) {
	t, err := NewSchema(s.Schema)
	if err != nil {
		return zero, fmt.Errorf(": %w", err)
	}
	return StructureField{
		Comment: s.Description,
		Name:    PublicFieldName(s.Name),
		Type:    t,
		Tags:    []StructureFieldTag{{Key: "json", Values: []string{s.Name}}},
	}, nil
}

type StructureFieldTag struct {
	Key    string
	Values []string
}

type Ref string

func NewRef(ref string) Ref {
	ref = ref[strings.LastIndex(ref, "/"):]
	ref = strings.TrimPrefix(ref, "/")
	return Ref(ref)
}

func (r Ref) Render() (string, error) { return string(r), nil }
