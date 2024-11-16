package test

// ------------------------
//         Schemas
// ------------------------

type PageCustom string

func NewPageCustom(v string) PageCustom {
	return PageCustom(v)
}

func (c PageCustom) String() string {
	return string(c)
}

type PageCustom_Schema string

func NewPageCustom_Schema(v string) PageCustom_Schema {
	return PageCustom_Schema(v)
}

func (c PageCustom_Schema) String() string {
	return string(c)
}
