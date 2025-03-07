package generator

import (
	"fmt"
	"strconv"
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
		"GoTypeFn": s.Items.RenderFieldType,
	})
}

func (s SliceType) RenderToBaseType(to, from string) (string, error) {
	return ExecuteTemplate("Slice_RenderToBaseType", TData{
		"To":   to,
		"From": from,

		"RenderGoTypeFn": s.RenderGoType,
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

func (s SliceType) RenderUnmarshalJSON(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Self":             s,
		"GoTypeFn":         s.Items.RenderGoType,
		"ItemsParseString": s.Items.ParseString,
	})
}

func (s SliceType) RenderMarshalJSON(to string, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("SliceType_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Self":             s,
		"GoTypeFn":         s.Items.RenderGoType,
		"ItemsParseString": s.Items.ParseString,
	})
}

type StructureType struct {
	NotImplementedParser
	NotImplementedFormatter

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

func (s StructureType) FuncTypeName() string { return "struct{}" }

func (s StructureType) RenderGoType() (string, error) { return ExecuteTemplate("StructureType", s) }

func (s StructureType) RenderToBaseType(to, from string) (string, error) {
	return to + " = " + from, nil
}

func (s StructureType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("StructureType_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Self": s,
	})
}

func (s StructureType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("StructureType_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Self": s,
	})
}

type StructureField struct {
	Comment     string
	Name        string
	GoTypeFn    GoTypeRenderFunc
	FieldTypeFn RenderFunc
	Type        SchemaType
	Schema      Schema
	JSONTag     string
	Embedded    bool
	Required    bool

	RenderToBaseTypeFn func(to, from string) (string, error)
}

func NewStructureField(p specification.SchemaProperty, components Componenter, cfg Config) (zero StructureField, _ Imports, _ error) {
	schema, ims, err := NewSchema(p.Schema, NamedComponenter{components, p.Name}, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema: %w", err)
	}

	if schema.Ref == nil && schema.Kind() == SchemaKindObject {
		sc := components.AddSchema(PublicFieldName(p.Name), schema, cfg)
		schema.Ref = sc
	}

	var st SchemaType = schema
	// if schema.IsNullable() {
	// 	st = NewNullableType(st, cfg)
	// }
	if !p.Required {
		st = NewOptionalType(st, cfg)
	}

	name := p.Name

	return StructureField{
		Comment:     p.Schema.Value().Description,
		Name:        PublicFieldName(name),
		Type:        st,
		Schema:      schema,
		JSONTag:     name,
		GoTypeFn:    st.RenderGoType,
		FieldTypeFn: st.RenderFieldType,
		Required:    p.Required,

		RenderToBaseTypeFn: schema.RenderToBaseType,
	}, ims, nil
}

func (s StructureField) Render() (string, error) { return ExecuteTemplate("StructureField", s) }

type CustomType struct {
	Value string
	Type  InternalSchemaType

	Pkg      string
	TypeName string
}

func NewCustomType(specCustom string, st InternalSchemaType) (CustomType, Imports) {
	var customImport, customType string = "", specCustom
	var customPkg string
	typeName := specCustom
	var hasCustomImport bool
	slIdx := strings.LastIndex(specCustom, "/")
	if slIdx >= 0 {
		hasCustomImport = true
		customType = specCustom[slIdx+1:]
		typeName = customType
	}

	dotIdx := strings.LastIndex(specCustom, ".")
	if dotIdx >= 0 {
		if hasCustomImport {
			customImport = specCustom[:dotIdx]
		}
		customPkg = specCustom[slIdx+1 : dotIdx]
		typeName = specCustom[dotIdx+1:]
	}

	return CustomType{
		Value: customType,
		Type:  st,

		Pkg:      customPkg,
		TypeName: typeName,
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

func (c CustomType) RenderToBaseType(to, from string) (string, error) {
	switch c.Type.Kind() {
	// case SchemaKindObject:
	default:
		from = from + "." + c.Type.FuncTypeName() + "()"
	}
	return to + " = " + from, nil
	// return c.Type.RenderToBaseType(to, from)
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

func (c CustomType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	// from = from + "." + c.Type.FuncTypeName() + "()"
	if c.Type.Kind() == SchemaKindObject {
		return ExecuteTemplate("CustomType_RenderUnmarshalJSON_Object", TData{
			"Base":       c.Type,
			"CustomType": c.Value,

			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	}
	return ExecuteTemplate("CustomType_RenderUnmarshalJSON", TData{
		"Base":       c.Type,
		"CustomType": c.Value,

		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
	})
}

func (c CustomType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	// from = from + "." + c.Type.FuncTypeName() + "()"

	if c.Type.Kind() == SchemaKindObject {
		return ExecuteTemplate("CustomType_RenderMarshalJSON_Object", TData{
			"Base":       c.Type,
			"CustomType": c.Value,
			"Pkg":        c.Pkg,

			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	}
	return ExecuteTemplate("CustomType_RenderMarshalJSON", TData{
		"Base":       c.Type,
		"CustomType": c.Value,
		"Pkg":        c.Pkg,

		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,
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

func (p OptionalType) FuncTypeName() string {
	return p.MaybeType + p.V.FuncTypeName()
}

func (p OptionalType) Kind() SchemaKind { return p.V.Kind() }

var _ GoTypeRender = OptionalType{}

func (p OptionalType) RenderGoType() (string, error) {
	out, err := p.V.RenderGoType()
	return p.MaybeType + "[" + out + "]", err
}

func (p OptionalType) RenderFieldType() (string, error) {
	out, err := p.V.RenderFieldType()
	return p.MaybeType + "[" + out + "]", err
}

var _ SchemaType = OptionalType{}

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

func (p OptionalType) RenderToBaseType(to, from string) (string, error) {
	return ExecuteTemplate("OptionalType_RenderToBaseType", TData{
		"To":   to,
		"From": from,
		"Type": p.V,
	})
}

func (p OptionalType) RenderFormat(from string) (string, error) {
	return p.V.RenderFormat(from)
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

func (p OptionalType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("OptionalType_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"V":    p.V,
		"Self": p,
	})
}

func (p OptionalType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("OptionalType_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"V":    p.V,
		"Self": p,
	})
}

type NullableType struct {
	V        InternalSchemaType
	TypeName string
}

func NewNullableType(v SchemaType, nullableType string) NullableType {
	typename := nullableType
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

func (n NullableType) RenderToBaseType(to, from string) (string, error) {
	return ExecuteTemplate("NullableType_RenderToBaseType", TData{
		"To":   to,
		"From": from,
		"Type": n.V,
	})
}

func (n NullableType) RenderFormat(from string) (string, error) {
	return n.V.RenderFormat(from + ".Value")
}

func (n NullableType) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return n.V.RenderFormatStrings(to, from+".Value", isNew)
}

func (n NullableType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("NullableType_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"V":    n.V,
		"Self": n,
	})
}

func (n NullableType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("NullableType_RenderMarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"V":    n.V,
		"Self": n,
	})
}

type RawBytesType struct {
	NotImplementedParser
	NotImplementedFormatter
}

var _ InternalSchemaType = RawBytesType{}

func (RawBytesType) RenderGoType() (string, error) {
	return "json.RawMessage", nil
}

func (RawBytesType) RenderToBaseType(to, from string) (string, error) {
	return to + " = " + from, nil
}

func (RawBytesType) FuncTypeName() string {
	return "RawMessage"
}

func (RawBytesType) Kind() SchemaKind {
	return SchemaKindRawBytes
}

func (RawBytesType) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return to + " = " + from, nil
}
func (RawBytesType) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	panic("not implemented")
}

type OneOfStructure struct {
	NotImplementedParser
	NotImplementedFormatter

	Elements []OneOfElement
	Fields   []StructureField

	DiscriminatorPropertyKey Maybe[string]
	DiscriminatorMapping     []DiscriminatorMapping
}

type DiscriminatorMapping struct {
	Key    string
	Values []string
}

var _ SchemaType = (*OneOfStructure)(nil)

func NewOneOfStructure(elems []specification.Ref[specification.Schema], d specification.Discriminator, c Componenter, cfg Config) (zero OneOfStructure, _ Imports, _ error) {
	var s OneOfStructure
	var imports Imports
	for i, e := range elems {
		name := "OneOf" + strconv.Itoa(i)
		schema, ims, err := NewSchema(e, NamedComponenter{c, name}, cfg)
		if err != nil {
			return zero, nil, fmt.Errorf("new oneOf schema for %d-th element: %w", i, err)
		}
		imports = append(imports, ims...)
		s.Elements = append(s.Elements, OneOfElement{
			Index:  i,
			Schema: schema,
		})
		var schemaType SchemaType = schema

		if schema.Ref != nil {
			name = schema.Ref.Name
		} else if schema.Type.Kind() == SchemaKindObject {
			sc := c.AddSchema(name, schema, cfg)
			schemaType = NewSchemaRef(sc)
		}
		s.Fields = append(s.Fields, StructureField{
			Name:        name,
			Type:        schemaType,
			FieldTypeFn: NewOptionalType(schemaType, cfg).RenderFieldType,
		})
	}
	if v, ok := d.PropertyKey.Get(); ok {
		s.DiscriminatorPropertyKey = Just(v)
		for _, v := range d.Mapping {
			s.DiscriminatorMapping = append(s.DiscriminatorMapping, DiscriminatorMapping{
				Key:    v.Key,
				Values: v.Values,
			})
		}
	}
	return s, imports, nil
}

var _ GoTypeRender = (*OneOfStructure)(nil)

func (o OneOfStructure) RenderGoType() (string, error) {
	return StructureType{Fields: o.Fields}.RenderGoType()
}

func (o OneOfStructure) RenderToBaseType(to, from string) (string, error) {
	panic("RenderToBaseType: not implemented")
}
func (o OneOfStructure) RenderFieldType() (string, error) {
	panic("RenderFieldType: not implemented")
}

func (o OneOfStructure) FuncTypeName() string {
	return "OneOfStructure"
}
func (o OneOfStructure) Kind() SchemaKind {
	return SchemaKindObject
}

func (o OneOfStructure) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return StructureType{Fields: o.Fields}.RenderUnmarshalJSON(to, from, isNew, mkErr)
}

func (o OneOfStructure) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	panic("RenderMarshalJSON: not implemented")
}

type OneOfElement struct {
	Index  int
	Schema Schema
}
