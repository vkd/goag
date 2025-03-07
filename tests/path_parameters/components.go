package test

// ------------------------
//         Schemas
// ------------------------

type Shop string

func NewShop(v string) Shop { return Shop(v) }

func (c Shop) String() string { return string(c) }
