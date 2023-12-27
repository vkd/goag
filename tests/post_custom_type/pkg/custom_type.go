package pkg

import "encoding"

type ShopType struct {
	V string
}

func (s ShopType) String() string { return s.V }

func (s *ShopType) UnmarshalText(v []byte) error {
	s.V = string(v)
	return nil
}

var _ encoding.TextUnmarshaler = (*ShopType)(nil)

type PetTag struct {
	V string
}

func (p *PetTag) UnmarshalText(s string) error {
	p.V = s
	return nil
}
