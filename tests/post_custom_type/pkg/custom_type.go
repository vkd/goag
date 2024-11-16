package pkg

import (
	"encoding/json"
	"fmt"
)

type ShopType struct {
	V string
}

func (s ShopType) String() string    { return s.V }
func (s ShopType) Strings() []string { return []string{s.String()} }

func (s *ShopType) ParseString(v string) error {
	s.V = v
	return nil
}

type PetTag struct {
	V string
}

func (p *PetTag) SetFromSchemaString(s string) error {
	p.V = s
	return nil
}

func (p PetTag) String() string { return p.V }

func (p PetTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.V)
}

func (p *PetTag) UnmarshalJSON(v []byte) error {
	var s string
	err := json.Unmarshal(v, &s)
	if err != nil {
		return fmt.Errorf("unmarshal string field: %w", err)
	}
	p.V = s
	return nil
}

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}
