package test

import (
	"github.com/vkd/goag/tests/custom_type/pkg"
)

// ------------------------
//         Schemas
// ------------------------

type PageCustom = pkg.Page

type Shop ShopName

func NewShop(v ShopName) Shop {
	return Shop(v)
}

func (c Shop) ShopName() ShopName {
	return ShopName(c)
}

type ShopName = pkg.Page
