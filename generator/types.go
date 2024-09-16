package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type SliceType struct {
	Items Schema
}

func (s SliceType) Kind() SchemaKind { return SchemaKindArray }

func (s SliceType) FuncTypeName() string {
	return s.Items.FuncTypeName() + "s"
}

func (s SliceType) RenderGoType() (string, error) {
	return ExecuteTemplate("SliceType_GoType", TData{
		"GoTypeFn": s.Items.RenderGoType,
	})
}

func (s SliceType) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf(".RenderFormat() function for SliceType is not supported for type %T", s.Items)
}

func (s SliceType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("Slice_RenderFormatStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,

		"ItemsRenderFormatFn": s.Items.RenderFormat,
	})
}

func (s SliceType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_ParseString", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"GoTypeFn":         s.Items.RenderGoType,
		"ItemsParseString": s.Items.ParseString,
	})
}

func (s SliceType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_ParseStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"GoTypeFn":         s.Items.RenderGoType,
		"ItemsParseString": s.Items.ParseString,
	})
}

type StructureType struct {
	Fields []StructureField

	AdditionalProperties *GoTypeRender
}

func NewStructureType(s *specification.Schema, components Componenter, cfg Config) (zero StructureType, _ Imports, _ error) {
	var stype StructureType
	var imports Imports
	for _, p := range s.Properties {
		f, ims, err := NewStructureField(p, components, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("new structure field %q: %w", p.Name, err)
		}
		imports = append(imports, ims...)

		stype.Fields = append(stype.Fields, f)
	}
	if additionalProperties, ok := s.AdditionalProperties.Get(); ok {
		additional, ims, err := NewSchema(additionalProperties, NamedComponenter{components, "AdditionalProperties"}, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("additional properties: %w", err)
		}
		imports = append(imports, ims...)
		render := GoTypeRender(additional)
		stype.AdditionalProperties = &render
	}

	return stype, imports, nil
}

var _ InternalSchemaType = StructureType{}

func (s StructureType) Kind() SchemaKind { return SchemaKindObject }

func (s StructureType) FuncTypeName() string { return "Structure" }

func (s StructureType) RenderGoType() (string, error) { return ExecuteTemplate("StructureType", s) }

func (s StructureType) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf(".RenderFormat() function for StructureType is not supported")
}

func (s StructureType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return "", fmt.Errorf(".RenderFormatStrings() function for StructureType is not supported")
}

func (s StructureType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return "", fmt.Errorf(".ParseString() function for StructureType is not supported")
}

func (s StructureType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return "", fmt.Errorf(".ParseStrings() function for StructureType is not supported")
}

type StructureField struct {
	Comment  string
	Name     string
	GoTypeFn GoTypeRenderFunc
	Type     SchemaType
	Schema   SchemaType
	Tags     []StructureFieldTag
	JSONTag  string
	Embedded bool

	RenderBaseFromFn func(prefix, from, suffix string) (string, error)
}

func NewStructureField(p specification.SchemaProperty, components Componenter, cfg Config) (zero StructureField, _ Imports, _ error) {
	schema, ims, err := NewSchema(p.Schema, NamedComponenter{components, p.Name}, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}

	if schema.Ref == nil && !schema.IsCustom() && schema.Kind() == SchemaKindObject {
		sc := components.AddSchema(PublicFieldName(p.Name), schema, cfg)
		schema.Ref = sc
	}

	var st SchemaType = schema
	if !p.Required {
		st = NewOptionalType(schema, cfg)
	}

	name := p.Name

	return StructureField{
		Comment:  p.Schema.Value().Description,
		Name:     PublicFieldName(name),
		Type:     st,
		Schema:   st,
		Tags:     []StructureFieldTag{{Key: "json", Values: []string{name}}},
		JSONTag:  name,
		GoTypeFn: st.RenderGoType,

		RenderBaseFromFn: st.RenderBaseFrom,
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

type CustomType struct {
	Value string
	Type  InternalSchemaType
}

func NewCustomType(specCustom string, st InternalSchemaType) (CustomType, Imports) {
	var customImport, customType string = "", specCustom
	slIdx := strings.LastIndex(specCustom, "/")
	if slIdx >= 0 {
		customImport = specCustom[:slIdx]
		customType = specCustom[slIdx+1:]

		dotIdx := strings.LastIndex(specCustom, ".")
		if dotIdx >= 0 {
			customImport = specCustom[:dotIdx]
		}
	}

	return CustomType{
		Value: customType,
		Type:  st,
	}, NewImportsS(customImport)
}

var _ InternalSchemaType = (*CustomType)(nil)

func (c CustomType) Kind() SchemaKind { return c.Type.Kind() }

func (c CustomType) FuncTypeName() string { return stringsTitle(strings.ReplaceAll(c.Value, ".", "")) }

var _ GoTypeRender = (*CustomType)(nil)

func (c CustomType) RenderGoType() (string, error) {
	return c.Value, nil
}

var _ Parser = (*CustomType)(nil)

func (c CustomType) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomType_ParseString", TData{
		"Base":       c.Type,
		"CustomType": c.Value,

		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (c CustomType) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomType_ParseStrings", TData{
		"Base":       c.Type,
		"CustomType": c.Value,

		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

var _ Formatter = (*CustomType)(nil)

func (c CustomType) RenderFormat(from string) (string, error) {
	return c.Type.RenderFormat(from + "." + c.Type.FuncTypeName() + "()")
}

func (c CustomType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("CustomType_RenderFormatStrings", TData{
		"Base": c.Type,

		"To":    to,
		"From":  from,
		"IsNew": isNew,
	})
}

func (c CustomType) RenderConvertToBaseSchema(from string) (string, error) {
	return from + "." + c.Type.FuncTypeName() + "()", nil
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

func (p OptionalType) FuncTypeName() string {
	return p.MaybeType + p.V.FuncTypeName()
}

func (p OptionalType) Kind() SchemaKind { return p.V.Kind() }

var _ GoTypeRender = OptionalType{}

func (p OptionalType) RenderGoType() (string, error) {
	out, err := p.V.RenderGoType()
	return p.MaybeType + "[" + out + "]", err
}

var _ SchemaType = OptionalType{}

func (p OptionalType) RenderBaseFrom(prefix, from, suffix string) (string, error) {
	return ExecuteTemplate("OptionalType_RenderBaseFrom", TData{
		"From":   from,
		"Prefix": prefix,
		"Suffix": suffix,
		"Self":   p,
		"Type":   p.V,
	})
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
	return ExecuteTemplate("OptionalType_RenderFormatStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"Self":  p,
		"Type":  p.V,
	})
}

func (p OptionalType) RenderConvertToBaseSchema(from string) (string, error) {
	return ExecuteTemplate("OptionalType_RenderConvertToBaseSchema", TData{
		"From": from,
		"Self": p,
		"Type": p.V,
	})
}

type NullableType struct {
	V        InternalSchemaType
	TypeName string
}

func NewNullableType(v InternalSchemaType, cfg Config) NullableType {
	typename := cfg.Nullable.Type
	if typename == "" {
		typename = "Nullable"
	}
	return NullableType{V: v, TypeName: typename}
}

func (n NullableType) FuncTypeName() string {
	return strings.ReplaceAll(n.TypeName, ".", "") + n.V.FuncTypeName()
}

func (n NullableType) Kind() SchemaKind { return n.V.Kind() }

var _ GoTypeRender = NullableType{}

func (n NullableType) GoType(from string) string {
	return n.TypeName + "[" + from + "]"
}

func (n NullableType) RenderGoType() (string, error) {
	out, err := n.V.RenderGoType()
	return n.GoType(out), err
}

func (n NullableType) RenderBaseFrom(prefix, from, suffix string) (string, error) {
	return ExecuteTemplate("NullableType_RenderBaseFrom", TData{
		"From":   from,
		"Prefix": prefix,
		"Suffix": suffix,
		"Self":   n,
		"Type":   n.V,
	})
}

var _ Parser = NullableType{}

func (n NullableType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("NullableTypeParseString", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
		"Self":  n,
		"Type":  n.V,
	})
}

func (n NullableType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("NullableTypeParseStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
		"Self":  n,
		"Type":  n.V,
	})
}

var _ Formatter = NullableType{}

func (n NullableType) RenderFormat(from string) (string, error) {
	return n.V.RenderFormat(from + ".Value")
}

func (n NullableType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return n.V.RenderFormatStrings(to, from+".Value", isNew)
}
