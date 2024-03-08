package generator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vkd/goag/specification"
)

type BoolType struct{}

var _ SchemaType = BoolType{}

func (b BoolType) Render() (string, error) { return "bool", nil }

func (b BoolType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("BoolParseString", struct {
		From  string
		To    string
		MkErr ErrorRender
		IsNew bool
	}{from, to, mkErr, isNew})
}

func (b BoolType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("BoolFormat", struct {
		From string
	}{from})
}

type IntType struct {
	BitSize int
}

var _ SchemaType = IntType{}

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

func (i IntType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 0:
		return ExecuteTemplate("IntParser", struct {
			From  string
			To    string
			MkErr ErrorRender
			IsNew bool
		}{from, to, mkErr, isNew})
	case 32:
		return ExecuteTemplate("Int32Parser", struct {
			From  string
			To    string
			MkErr ErrorRender
			IsNew bool
		}{from, to, mkErr, isNew})
	case 64:
		return ExecuteTemplate("Int64Parser", struct {
			From  string
			To    string
			MkErr ErrorRender
			IsNew bool
		}{from, to, mkErr, isNew})
	default:
		return "", fmt.Errorf("unsupported int bit size %d", i.BitSize)
	}
}

func (i IntType) RenderFormat(from string) (string, error) {
	switch i.BitSize {
	case 0:
		return ExecuteTemplate("IntFormat", struct {
			From string
		}{from})
	case 32:
		return ExecuteTemplate("Int32Format", struct {
			From string
		}{from})
	case 64:
		return ExecuteTemplate("Int64Format", struct {
			From string
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

func (i FloatType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 32:
		return ExecuteTemplate("Float32Parser", struct {
			From  string
			To    string
			IsNew bool
			MkErr ErrorRender
		}{from, to, isNew, mkErr})
	case 64:
		return ExecuteTemplate("Float64Parser", struct {
			From  string
			To    string
			IsNew bool
			MkErr ErrorRender
		}{from, to, isNew, mkErr})
	default:
		return "", fmt.Errorf("unsupported float bit size %d", i.BitSize)
	}
}

var _ Formatter = FloatType{}

func (i FloatType) RenderFormat(from string) (string, error) {
	switch i.BitSize {
	case 32:
		return ExecuteTemplate("Float32Format", struct {
			From string
		}{from})
	case 64:
		return ExecuteTemplate("Float64Format", struct {
			From string
		}{from})
	default:
		return "", fmt.Errorf("unsupported float bit size %d", i.BitSize)
	}
}

type StringType struct{}

var _ SchemaType = StringType{}

func (s StringType) Render() (string, error) { return "string", nil }

func (_ StringType) RenderFormat(from string) (string, error) {
	return from, nil
}

func (_ StringType) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if isNew {
		return to + " := " + from, nil
	}
	return to + " = " + from, nil
}

type CustomType string

func NewCustomType(s string) (CustomType, Imports) {
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

	return CustomType(customType), NewImportsS(customImport)
}

var _ SchemaType = CustomType("")

func (c CustomType) Render() (string, error) {
	return string(c), nil
}

func (c CustomType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("CustomTypeFormat", struct {
		From string
	}{From: from})
}

func (c CustomType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomTypeParser", TData{
		"To":    to,
		"Type":  c,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

type SliceType struct {
	Items interface {
		Render
		Parser
	}
}

func (s SliceType) Render() (string, error) { return ExecuteTemplate("SliceType", s) }

func (s SliceType) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf("SliceType doesn't implement Formatter interface")
}

func (s SliceType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceTypeParseString", TData{
		"From": from,
		"To":   to,
		"Items": ParserFunc(func(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
			return Renders{
				RenderFunc(func() (string, error) {
					return s.Items.ParseString("v1", from, true, mkErr)
				}),
				RenderFunc(func() (string, error) {
					return Assign(to, "v1", isNew), nil
				}),
			}.Render()
		}),
		"MkErr": mkErr,
		"IsNew": isNew,
	})
}

func (s SliceType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch s.Items.(type) {
	case StringType:
		return Assign(to, from, isNew), nil
	}

	return ExecuteTemplate("SliceTypeParseStrings", TData{
		"From":        from,
		"To":          to,
		"ItemsRender": s.Items,
		"Items": ParserFunc(func(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
			return Renders{
				RenderFunc(func() (string, error) {
					return s.Items.ParseString("v1", from, true, mkErr)
				}),
				RenderFunc(func() (string, error) {
					return Assign(to, "v1", isNew), nil
				}),
			}.Render()
		}),
		"MkErr": mkErr,
		"IsNew": isNew,
	})
}

type StructureType struct {
	Fields []StructureField
}

func NewStructureType(s *specification.Schema) (zero StructureType, _ Imports, _ error) {
	var stype StructureType
	var imports Imports
	for _, p := range s.Properties {
		f, ims, err := NewStructureField(p)
		if err != nil {
			return zero, nil, fmt.Errorf("field %q: %w", p.Name, err)
		}
		stype.Fields = append(stype.Fields, f)
		imports = append(imports, ims...)
	}
	if s.AdditionalProperties.IsSet {
		additional, ims, err := NewSchema(s.AdditionalProperties.Value)
		if err != nil {
			return zero, nil, fmt.Errorf("additional properties: %w", err)
		}
		imports = append(imports, ims...)
		stype.Fields = append(stype.Fields, StructureField{
			Name: "AdditionalProperties",
			Type: NewMapType(StringType{}, additional),
			Tags: []StructureFieldTag{{Key: "json", Values: []string{"-"}}},
		})
	}
	return stype, imports, nil
}

var _ SchemaType = StructureType{}

func (s StructureType) Render() (string, error) { return ExecuteTemplate("StructureType", s) }
func (s StructureType) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf("not implemented")
}
func (s StructureType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return "", fmt.Errorf("not implemented")
}

type StructureField struct {
	Comment string
	Name    string
	Type    Render
	Tags    []StructureFieldTag
}

func NewStructureField(s specification.SchemaProperty) (zero StructureField, _ Imports, _ error) {
	t, ims, err := NewSchema(s.Schema)
	if err != nil {
		return zero, nil, fmt.Errorf(": %w", err)
	}
	return StructureField{
		Comment: s.Schema.Value().Description,
		Name:    PublicFieldName(s.Name),
		Type:    t,
		Tags:    []StructureFieldTag{{Key: "json", Values: []string{s.Name}}},
	}, ims, nil
}

func (sf StructureField) GetTag(k string) (zero StructureFieldTag, _ bool) {
	for _, t := range sf.Tags {
		if t.Key == k {
			return t, true
		}
	}
	return zero, false
}

type StructureFieldTag struct {
	Key    string
	Values []string
}

type Ref string

func NewRef(ref string) Ref {
	// ref = ref[strings.LastIndex(ref, "/"):]
	// ref = strings.TrimPrefix(ref, "/")
	return Ref(ref)
}

var _ SchemaType = Ref("")

func (r Ref) Render() (string, error) { return string(r), nil }

func (r Ref) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (r Ref) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

type PointerType struct {
	V SchemaType
}

func NewPointerType(v SchemaType) PointerType {
	return PointerType{V: v}
}

var _ Render = PointerType{}

func (p PointerType) Render() (string, error) {
	out, err := p.V.Render()
	return "*" + out, err
}

var _ Parser = PointerType{}

func (p PointerType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("PointerTypeParseString", TData{
		"To": to,
		"From": RenderFunc(func() (string, error) {
			return p.V.ParseString("v", from, true, mkErr)
		}),
		"IsNew": isNew,
	})
}

var _ Formatter = PointerType{}

func (p PointerType) RenderFormat(from string) (string, error) {
	return p.V.RenderFormat(from)
}

type MapType struct {
	Key   SchemaType
	Value SchemaType
}

func NewMapType(k, v SchemaType) MapType {
	return MapType{Key: k, Value: v}
}

var _ SchemaType = MapType{}

func (m MapType) Render() (string, error) {
	return ExecuteTemplate("MapType", m)
}

func (m MapType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (m MapType) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf("not implemented")
}
