package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
}

var _ json.Unmarshaler = (*Metadata)(nil)

func (m *Metadata) UnmarshalJSON(bs []byte) error {
	type tp Metadata
	var v tp
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return fmt.Errorf("unmarshal metadata: %w", err)
	}
	*m = Metadata(v)
	return nil
}

type Settings struct {
	Theme Maybe[string] `json:"theme"`
}

var _ json.Unmarshaler = (*Settings)(nil)

func (m *Settings) UnmarshalJSON(bs []byte) error {
	type tp Settings
	var v tp
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return fmt.Errorf("unmarshal settings: %w", err)
	}
	*m = Settings(v)
	return nil
}

type Environments []Environment

type Environment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var _ json.Unmarshaler = (*Environment)(nil)

func (m *Environment) UnmarshalJSON(bs []byte) error {
	type tp Environment
	var v tp
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return fmt.Errorf("unmarshal environment: %w", err)
	}
	*m = Environment(v)
	return nil
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

var _ json.Marshaler = (*Nullable[any])(nil)

func (m Nullable[T]) MarshalJSON() ([]byte, error) {
	if m.IsSet {
		return json.Marshal(&m.Value)
	}
	return []byte(nullValue), nil
}

var _ json.Unmarshaler = (*Nullable[any])(nil)

const nullValue = "null"

var nullValueBs = []byte(nullValue)

func (m *Nullable[T]) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, nullValueBs) {
		m.IsSet = false
		return nil
	}
	m.IsSet = true
	return json.Unmarshal(bs, &m.Value)
}
