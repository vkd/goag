package generator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vkd/goag/specification"
)

type BoolType struct {
	SingleValue
}

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
	SingleValue
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
	SingleValue
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

type StringType struct {
	SingleValue
}

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

type CustomType struct {
	Multivalue
	Type   string
	Import string
}

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

	return CustomType{
		Type:   customType,
		Import: customImport,
	}, NewImportsS(customImport)
}

var _ SchemaType = (*CustomType)(nil)

func (c CustomType) Render() (string, error) {
	return string(c.Type), nil
}

func (c CustomType) RenderFormat(from string) (string, error) {
	return from + ".String()", nil
}

func (c CustomType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomTypeParserExternal", TData{
		"To":    to,
		"Type":  c,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

type SliceType struct {
	Multivalue
	Items SchemaType
}

func (s SliceType) Render() (string, error) { return ExecuteTemplate("SliceType", s) }

func (s SliceType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("SliceTypeRenderFormat", TData{
		"From": from,
		"Type": s,
	})
}

func (s SliceType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return s.ParseStrings(to, from, isNew, mkErr)
}

func (s SliceType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	switch s.Items.(type) {
	case StringType:
		return Assign(to, from, isNew), nil
	}

	return ExecuteTemplate("SliceTypeParseStrings", TData{
		"From":  from,
		"To":    to,
		"Items": s.Items,
		"MkErr": mkErr,
		"IsNew": isNew,
	})
}

type StructureType struct {
	SingleValue
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
	if s.AdditionalProperties.Set {
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
	return ExecuteTemplate("StructureTypeRenderFormat", TData{
		"From":  from,
		"Type":  s,
		"MkErr": newError{},
	})
}
func (s StructureType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if isNew {
		return "/* isNew == true is not supported */", nil
	}
	return ExecuteTemplate("StructureTypeParseString", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
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

func (s StructureField) Render() (string, error) { return ExecuteTemplate("StructureField", s) }

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

type Ref[T any] struct {
	// Multivalue
	Ref        string
	SchemaType specification.Ref[T]
}

func NewRef[T any](ref *specification.Object[string, specification.Ref[T]]) Ref[T] {
	// ref = ref[strings.LastIndex(ref, "/"):]
	// ref = strings.TrimPrefix(ref, "/")
	return Ref[T]{Ref: ref.Name, SchemaType: ref.V}
}

var _ SchemaType = Ref[any]{}

func (r Ref[T]) Render() (string, error) { return r.Ref, nil }

func (r Ref[T]) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("RefParseString", TData{
		"To":    to,
		"Type":  r.Ref,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (r Ref[T]) IsMultivalue() bool {
	switch tp := any(*r.SchemaType.Value()).(type) {
	case SchemaType:
		return r.IsMultivalue()
	case specification.Schema:
		switch tp.Type {
		case "array":
			return true
		default:
			return false
		}
	case specification.QueryParameter:
		return true
	case specification.PathParameter:
		return false
	case specification.HeaderParameter:
		switch tp.Schema.Value().Type {
		case "array":
			return true
		default:
			return false
		}
	default:
		panic(fmt.Errorf("unsupported Ref[T].IsMultivalue() type: %T", tp))
	}
}

func (r Ref[T]) ParseQuery(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("RefParseQuery", TData{
		"To":    to,
		"Type":  r.Ref,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (r Ref[T]) ParseSchema(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("RefParseSchema", TData{
		"To":    to,
		"Type":  r.Ref,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}
func (r Ref[T]) RenderFormat(from string) (string, error) {
	if r.IsMultivalue() {
		return string(from) + ".Strings()", nil
	}
	return string(from) + ".String()", nil
}

type OptionalType struct {
	V         SchemaType
	MaybeType string
}

func NewOptionalType(v SchemaType, typename string) OptionalType {
	if typename == "" {
		typename = "Maybe"
	}
	return OptionalType{V: v, MaybeType: typename}
}

var _ Render = OptionalType{}

func (p OptionalType) Render() (string, error) {
	out, err := p.V.Render()
	return p.MaybeType + "[" + out + "]", err
}

var _ Parser = OptionalType{}

func (p OptionalType) ParseString(to string, from string, _ bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("OptionalTypeParseString", TData{
		"To":   to,
		"Type": p.V,
		"From": RenderFunc(func() (string, error) {
			return p.V.ParseString("v", from, true, mkErr)
		}),
	})
}

func (p OptionalType) IsMultivalue() bool { return p.V.IsMultivalue() }

var _ Formatter = OptionalType{}

func (p OptionalType) RenderFormat(from string) (string, error) {
	return p.V.RenderFormat(from)
}

type MapType struct {
	SingleValue
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
