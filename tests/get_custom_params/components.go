package test

// ------------------------
//         Schemas
// ------------------------

type PageCustom string

func NewPageCustom(v string) PageCustom { return PageCustom(v) }

func (c PageCustom) String() string { return string(c) }
