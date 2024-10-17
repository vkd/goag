package test

import (
	"github.com/vkd/goag/tests/get_custom_params/pkg"
)

// ------------------------
//         Schemas
// ------------------------

type PageCustom = pkg.Page

type PageCustomSchema string

func NewPageCustomSchema(v string) PageCustomSchema {
	return PageCustomSchema(v)
}

func (c PageCustomSchema) String() string {
	return string(c)
}
