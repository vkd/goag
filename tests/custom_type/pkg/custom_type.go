package pkg

import "fmt"

type Page string

func (s Page) String() string { return string(s) }

func (s *Page) ParseString(v string) error {
	*s = Page(v)
	return nil
}

type PageCustomTypeQuery string

func (s PageCustomTypeQuery) String() string { return string(s) }

func (s *PageCustomTypeQuery) ParseString(v string) error {
	*s = PageCustomTypeQuery(v)
	return nil
}

type Shop string

func (s *Shop) ParseString(str string) error {
	*s = Shop(str)
	return nil
}

func (s Shop) String() string {
	return string(s)
}

type Metadata struct {
	InternalID string
	OK         bool
}

type MetadataSchema struct {
	InnerID Maybe[string]
}

func (m *Metadata) SetFromSchemaMetadata(v MetadataSchema) error {
	*m = Metadata{
		InternalID: v.InnerID.Value,
		OK:         v.InnerID.IsSet,
	}
	return nil
}

func (m Metadata) ToSchemaMetadata() MetadataSchema {
	return MetadataSchema{
		InnerID: Just(m.InternalID),
	}
}

type Settings struct {
	Theme Maybe[string] `json:"theme"`
}

type SettingsSchema struct {
	Theme Maybe[string]
}

type GetShopAdditionals struct {
	AdditionalProperties map[string]any
}

func (s *Settings) SetFromSchemaSettings(v SettingsSchema) error {
	*s = Settings{
		Theme: v.Theme,
	}
	return nil
}

func (s *Settings) SetFromSchemaGetShopAdditionals(v GetShopAdditionals) error {
	s.Theme.IsSet = false
	if theme, ok := v.AdditionalProperties["theme"]; ok {
		switch theme := theme.(type) {
		case string:
			s.Theme.Set(theme)
		default:
			return fmt.Errorf("unknown type: %T", theme)
		}
	}
	return nil
}

func (s Settings) ToSchemaSettings() SettingsSchema {
	return SettingsSchema{
		Theme: s.Theme,
	}
}

func (s Settings) ToSchemaGetShopAdditionals() GetShopAdditionals {
	if t, ok := s.Theme.Get(); ok {
		return GetShopAdditionals{
			AdditionalProperties: map[string]any{
				"theme": t,
			},
		}
	}
	return GetShopAdditionals{}
}

type Environments []Environment

type Environment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EnvironmentSchema struct {
	Name  string
	Value string
}

func (e *Environment) SetFromSchemaEnvironment(v EnvironmentSchema) error {
	*e = Environment{
		Name:  v.Name,
		Value: v.Value,
	}
	return nil
}

func (e Environment) ToSchemaEnvironment() EnvironmentSchema {
	return EnvironmentSchema{
		Name:  e.Name,
		Value: e.Value,
	}
}

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

type Nullable[T any] struct {
	IsSet bool
	Value T
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{}
}

func Pointer[T any](v T) Nullable[T] {
	return Nullable[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Nullable[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Nullable[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}
