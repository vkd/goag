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

func (b BoolType) Base() SchemaType {
	return b
}

func (b BoolType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("BoolParseString", struct {
		From  string
		To    string
		MkErr ErrorRender
		IsNew bool
	}{from, to, mkErr, isNew})
}

func (b BoolType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return b.ParseString(to, from+"[0]", isNew, mkErr)
}

func (b BoolType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("BoolFormat", struct {
		From string
	}{from})
}

func (b BoolType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("SliceSingleElementFormatStrings", TData{
		"Item":  b,
		"From":  from,
		"To":    to,
		"IsNew": isNew,
	})
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

func (i IntType) Base() SchemaType {
	return i
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

func (i IntType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return i.ParseString(to, from+"[0]", isNew, mkErr)
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

func (i IntType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("SliceSingleElementFormatStrings", TData{
		"Item":  i,
		"From":  from,
		"To":    to,
		"IsNew": isNew,
	})
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

func (f FloatType) Base() SchemaType {
	return f
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

func (i FloatType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return i.ParseString(to, from+"[0]", isNew, mkErr)
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

func (i FloatType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("SliceSingleElementFormatStrings", TData{
		"Item":  i,
		"From":  from,
		"To":    to,
		"IsNew": isNew,
	})
}

type StringType struct {
	SingleValue
}

var _ SchemaType = StringType{}

func (s StringType) Render() (string, error) { return "string", nil }

func (s StringType) Base() SchemaType {
	return s
}

func (_ StringType) RenderFormat(from string) (string, error) {
	return from, nil
}

func (s StringType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("SliceSingleElementFormatStrings", TData{
		"Item":  s,
		"From":  from,
		"To":    to,
		"IsNew": isNew,
	})
}

func (_ StringType) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if isNew {
		return to + " := " + from, nil
	}
	return to + " = " + from, nil
}

func (_ StringType) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if isNew {
		return to + " := " + from + "[0]", nil
	}
	return to + " = " + from + "[0]", nil
}

type CustomType struct {
	Type   string
	Import string

	BaseType SchemaType
}

func NewCustomType(s string, baseSchema SchemaType) (CustomType, Imports) {
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
		Type:     customType,
		Import:   customImport,
		BaseType: baseSchema,
	}, NewImportsS(customImport)
}

var _ SchemaType = (*CustomType)(nil)

func (c CustomType) IsMultivalue() bool {
	return c.Base().IsMultivalue()
}

func (c CustomType) Render() (string, error) {
	return string(c.Type), nil
}

func (c CustomType) Base() SchemaType {
	return c.BaseType
}

func BaseFuncName(st SchemaType) Render {
	switch base := st.Base().(type) {
	case SliceType:
		return base.Items
	default:
		return base
	}
}

func (c CustomType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("CustomTypeRenderFormat", TData{
		"From": from,
		"Base": c.BaseType.Base(),

		"IsMultivalue": c.BaseType.IsMultivalue(),
	})
}

func (c CustomType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("CustomTypeRenderFormatStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"Base":  c.BaseType.Base(),

		"IsMultivalue": c.BaseType.Base().IsMultivalue(),
		"BaseFunc":     BaseFuncName(c.BaseType),
	})
}

func (c CustomType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomTypeParseString", TData{
		"To":           to,
		"Type":         c,
		"From":         from,
		"IsNew":        isNew,
		"MkErr":        mkErr,
		"Base":         c.BaseType.Base(),
		"IsMultivalue": c.BaseType.IsMultivalue(),
		"BaseFunc":     BaseFuncName(c.BaseType),
	})
}

func (c CustomType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomTypeParseStrings", TData{
		"To":           to,
		"Type":         c,
		"From":         from,
		"IsNew":        isNew,
		"MkErr":        mkErr,
		"Base":         c.BaseType,
		"IsMultivalue": c.BaseType.IsMultivalue(),
		"BaseFunc":     BaseFuncName(c.BaseType),
	})
}

type SliceType struct {
	Multivalue
	Items SchemaType
}

func (s SliceType) Render() (string, error) { return ExecuteTemplate("SliceType", s) }

func (s SliceType) Base() SchemaType {
	return s
}

func (s SliceType) RenderFormat(from string) (string, error) {
	switch s.Items.(type) {
	case StringType:
		return from, nil
	}
	return "", fmt.Errorf(".RenderFormat() function for SliceType is not supported for type %T", s.Items)
}

func (s SliceType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	switch s.Items.(type) {
	case StringType:
		return Assign(to, from, isNew), nil
	}
	return ExecuteTemplate("SliceRenderFormatStrings", TData{
		"From":  from,
		"Items": s.Items,
		"To":    to,
		"IsNew": isNew,
	})
}

func (s SliceType) RenderFormatStringsMultiline(to, from string) (string, error) {
	return ExecuteTemplate("SliceTypeRenderFormatMultiline", TData{
		"To":    to,
		"From":  from,
		"Items": s.Items,
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

	AdditionalProperties *SchemaType
}

func NewStructureType(s *specification.Schema, components Components) (zero StructureType, _ Imports, _ error) {
	var stype StructureType
	var imports Imports
	for _, p := range s.Properties {
		t, ims, err := NewSchema(p.Schema, components)
		if err != nil {
			return zero, nil, fmt.Errorf("new schema: %w", err)
		}
		imports = append(imports, ims...)

		f, err := NewStructureField(p.Schema, p.Name, t, components)
		if err != nil {
			return zero, nil, fmt.Errorf("field %q: %w", p.Name, err)
		}

		stype.Fields = append(stype.Fields, f)
	}
	if additionalProperties, ok := s.AdditionalProperties.Get(); ok {
		additional, ims, err := NewSchema(additionalProperties, components)
		if err != nil {
			return zero, nil, fmt.Errorf("additional properties: %w", err)
		}
		imports = append(imports, ims...)
		stype.AdditionalProperties = &additional
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

func (s StructureType) Base() SchemaType {
	return s
}

func (s StructureType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("StructureTypeRenderFormat", TData{
		"From":  from,
		"Type":  s,
		"MkErr": newError{},
		"To":    to,
		"IsNew": isNew,
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

func (s StructureType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
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
	Comment  string
	Name     string
	Type     Render
	Tags     []StructureFieldTag
	JSONTag  string
	Embedded bool
}

func NewStructureField(s specification.Ref[specification.Schema], name string, t SchemaType, components Components) (zero StructureField, _ error) {
	return StructureField{
		Comment: s.Value().Description,
		Name:    PublicFieldName(name),
		Type:    t,
		Tags:    []StructureFieldTag{{Key: "json", Values: []string{name}}},
		JSONTag: name,
	}, nil
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

type RefSchemaType struct {
	Name string
	Ref  *SchemaComponent
}

var _ SchemaType = RefSchemaType{}

func NewRefSchemaType(name string, next *SchemaComponent) RefSchemaType {
	return RefSchemaType{
		Name: name,
		Ref:  next,
	}
}

func (r RefSchemaType) Base() SchemaType { return r.Ref.Type }

func (r RefSchemaType) Render() (string, error) { return r.Name, nil }

func (r RefSchemaType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("RefSchemaType_RenderFormat", TData{
		"SchemaType":   r.Ref.Type,
		"From":         from,
		"FuncName":     BaseFuncName(r),
		"IsMultivalue": r.IsMultivalue(),
	})
}

func (r RefSchemaType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("RefSchemaType_RenderFormatStrings", TData{
		"Ref":          r.Ref,
		"From":         from,
		"To":           to,
		"IsNew":        isNew,
		"FuncName":     BaseFuncName(r),
		"IsMultivalue": r.IsMultivalue(),
	})
}

func (r RefSchemaType) IsMultivalue() bool {
	return r.Base().IsMultivalue()
}

func (r RefSchemaType) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("RefSchemaType_ParseString", TData{
		"Ref":      r.Ref.Type,
		"Name":     r.Name,
		"From":     from,
		"To":       to,
		"IsNew":    isNew,
		"MkErr":    mkErr,
		"FuncName": BaseFuncName(r),
	})
}

func (r RefSchemaType) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("RefSchemaType_ParseStrings", TData{
		"Ref":      r.Ref,
		"Name":     r.Name,
		"From":     from,
		"To":       to,
		"IsNew":    isNew,
		"MkErr":    mkErr,
		"FuncName": BaseFuncName(r),
	})
}

type OptionalType struct {
	V         SchemaType
	MaybeType string
}

func NewOptionalType(v SchemaType, cfg Config) OptionalType {
	typename := cfg.Maybe.Type
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

func (o OptionalType) Base() SchemaType {
	return o.V
}

var _ Parser = OptionalType{}

func (p OptionalType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("OptionalTypeParseString", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
		"Self":  p,
		"Type":  p.V,
	})
}

func (p OptionalType) IsMultivalue() bool { return p.V.IsMultivalue() }

func (p OptionalType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("OptionalTypeParseStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
		"Self":  p,
		"Type":  p.V,
	})
}

var _ Formatter = OptionalType{}

func (p OptionalType) RenderFormat(from string) (string, error) {
	return p.V.RenderFormat(from + ".Value")
}

func (p OptionalType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return p.V.RenderFormatStrings(to, from+".Value", isNew)
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

func (m MapType) Base() SchemaType {
	return m
}

func (m MapType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (m MapType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (m MapType) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (m MapType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return "", fmt.Errorf("not implemented")
}
