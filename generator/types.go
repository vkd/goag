package generator

import (
	"fmt"

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

func (s SliceType) Render() (string, error) { return ExecuteTemplate("SliceType", s) }

func (s SliceType) Base() SchemaType {
	return s
}

func (s SliceType) RenderFormat(from string) (string, error) {
	return "", fmt.Errorf(".RenderFormat() function for SliceType is not supported for type %T", s.Items)
}

func (s SliceType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return ExecuteTemplate("Slice_RenderFormatStrings", TData{
		"From":  from,
		"Items": s.Items,
		"To":    to,
		"IsNew": isNew,
	})
}

func (s SliceType) RenderFormatStringsMultiline(to, from string) (string, error) {
	return ExecuteTemplate("SliceType_RenderFormatMultiline", TData{
		"To":    to,
		"From":  from,
		"Items": s.Items,
	})
}

func (s SliceType) ParseString(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_ParseString", TData{
		"From":  from,
		"To":    to,
		"Items": s.Items,
		"MkErr": mkErr,
		"IsNew": isNew,
	})
}

func (s SliceType) ParseStrings(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_ParseStrings", TData{
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

	AdditionalProperties *Render
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
		render := Render(additional)
		stype.AdditionalProperties = &render
	}

	return stype, imports, nil
}

var _ SchemaType = StructureType{}

func (s StructureType) FuncTypeName() string { return "Structure" }

func (s StructureType) Render() (string, error) { return ExecuteTemplate("StructureType", s) }
func (s StructureType) RenderFormat(from string) (string, error) {
	return ExecuteTemplate("StructureTypeRenderFormat", TData{
		"From":  from,
		"Type":  s,
		"MkErr": newError{},
	})
}

func (s StructureType) Kind() SchemaKind { return SchemaKindObject }

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
	Schema   Schema
	Tags     []StructureFieldTag
	JSONTag  string
	Embedded bool
}

func NewStructureField(s specification.Ref[specification.Schema], name string, t Schema, components Components) (zero StructureField, _ error) {
	return StructureField{
		Comment: s.Value().Description,
		Name:    PublicFieldName(name),
		Type:    t,
		Schema:  t,
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

var _ Render = OptionalType{}

func (p OptionalType) Render() (string, error) {
	out, err := p.V.Render()
	return p.MaybeType + "[" + out + "]", err
}

func (o OptionalType) Base() SchemaType {
	return o.V.BaseSchemaType()
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
	Value Schema
}

func NewMapType(v Schema) MapType {
	return MapType{Key: NewPrimitive(StringType{}), Value: v}
}

var _ SchemaType = MapType{}

func (m MapType) Kind() SchemaKind { return SchemaKindMap }

func (m MapType) FuncTypeName() string { return "Map" }

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
