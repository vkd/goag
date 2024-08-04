package test

import (
	"fmt"

	"github.com/vkd/goag/tests/custom_type/pkg"
)

// ------------------------
//         Schemas
// ------------------------

type PageCustom = pkg.Page

type Shop ShopName

func (c *Shop) ParseString(s string) error {
	var v ShopName
	err := v.ParseString(s)
	if err != nil {
		return fmt.Errorf("parse ShopName: %w", err)
	}
	*c = Shop(v)
	return nil
}

func (q Shop) String() string {
	return ShopName(q).String()
}

func (q Shop) Strings() []string {
	return []string{q.String()}
}

type ShopName = pkg.Page
