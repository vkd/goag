package test

type ShopType struct {
	V string
}

func (s ShopType) String() string { return s.V }

func (s *ShopType) UnmarshalText(v string) error {
	s.V = v
	return nil
}

type PetTag struct {
	V string
}

func (p *PetTag) UnmarshalText(s string) error {
	p.V = s
	return nil
}
