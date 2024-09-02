package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type SliceType struct {
	Multivalue
	Items Schema
}

func (s SliceType) Kind() SchemaKind { return SchemaKindArray }

func (s SliceType) FuncTypeName() string {
	return s.Items.FuncTypeName() + "s"
}

func (s SliceType) Render() (string, error) {
	return ExecuteTemplate("SliceType", TData{
		"ItemsRender": s.Items.Render,
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

		"ItemsRenderFormat": s.Items.RenderFormat,
	})
}

func (s SliceType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_ParseString", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"ItemsRender":      s.Items.Render,
		"ItemsParseString": s.Items.ParseString,
	})
}

func (s SliceType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_ParseStrings", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"ItemsRender":      s.Items.Render,
		"ItemsParseString": s.Items.ParseString,
	})
}

type StructureType struct {
	SingleValue
	Fields []StructureField

	AdditionalProperties *Render
}

func NewStructureType(s *specification.Schema, components Componenter, cfg Config) (zero StructureType, _ Imports, _ error) {
	var stype StructureType
	var imports Imports
	for _, p := range s.Properties {
		t, ims, err := NewSchema(p.Schema, NamedComponenter{components, p.Name}, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("new schema: %w", err)
		}
		imports = append(imports, ims...)

		if t.Ref == nil && !t.IsCustom() && t.Kind() == SchemaKindObject {
			sc := components.AddSchema(PublicFieldName(p.Name), t, cfg)
			t.Ref = sc
		}

		f := NewStructureField(p.Schema, p.Name, t, p.Required, cfg)

		stype.Fields = append(stype.Fields, f)
	}
	if additionalProperties, ok := s.AdditionalProperties.Get(); ok {
		additional, ims, err := NewSchema(additionalProperties, NamedComponenter{components, "AdditionalProperties"}, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("additional properties: %w", err)
		}
		imports = append(imports, ims...)
		render := Render(additional)
		stype.AdditionalProperties = &render
	}

	return stype, imports, nil
}

var _ SchemaType = StructureType{}

func (s StructureType) Kind() SchemaKind { return SchemaKindObject }

func (s StructureType) FuncTypeName() string { return "Structure" }

func (s StructureType) Render() (string, error) { return ExecuteTemplate("StructureType", s) }

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
	Type     Render
	Schema   SchemaType
	Tags     []StructureFieldTag
	JSONTag  string
	Embedded bool
}

func NewStructureField(s specification.Ref[specification.Schema], name string, t Schema, required bool, cfg Config) StructureField {
	var tp SchemaType = t
	if !required {
		ot := NewOptionalType(t, cfg)
		tp = ot
	}
	return StructureField{
		Comment: s.Value().Description,
		Name:    PublicFieldName(name),
		Type:    tp,
		Schema:  tp,
		Tags:    []StructureFieldTag{{Key: "json", Values: []string{name}}},
		JSONTag: name,
	}
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
	Type  SchemaType
}

func NewCustomType(specCustom string, st SchemaType) (CustomType, Imports) {
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

var _ SchemaType = (*CustomType)(nil)

func (c CustomType) Kind() SchemaKind { return c.Type.Kind() }

func (c CustomType) FuncTypeName() string { return strings.ReplaceAll(c.Value, ".", "") }

var _ Render = (*CustomType)(nil)

func (c CustomType) Render() (string, error) {
	return c.Value, nil
}

var _ Parser = (*CustomType)(nil)

func (c CustomType) IsMultivalue() bool { return c.Type.IsMultivalue() }

func (c CustomType) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomType_ParseString", TData{
		"Base":         c.Type,
		"Type":         StringRender(c.Value),
		"FuncTypeName": c.Type.FuncTypeName(),

		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (c CustomType) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("CustomType_ParseStrings", TData{
		"Base":         c.Type,
		"Type":         StringRender(c.Value),
		"FuncTypeName": c.Type.FuncTypeName(),

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
		"Base":         c.Type,
		"IsMultivalue": c.Type.IsMultivalue(),

		"To":    to,
		"From":  from + "." + c.Type.FuncTypeName() + "()",
		"IsNew": isNew,
	})
}

type OptionalType struct {
	V         Schema
	MaybeType string
}

func NewOptionalType(v Schema, cfg Config) OptionalType {
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

var _ Render = OptionalType{}

func (p OptionalType) Render() (string, error) {
	out, err := p.V.Render()
	return p.MaybeType + "[" + out + "]", err
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
