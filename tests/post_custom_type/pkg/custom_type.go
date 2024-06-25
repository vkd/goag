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

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}
