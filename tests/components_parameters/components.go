package test

// ------------------------
//         Schemas
// ------------------------

type Organization int

func NewOrganization(v int) Organization {
	return Organization(v)
}

func (c Organization) Int() int {
	return int(c)
}

type Page int32

func NewPage(v int32) Page {
	return Page(v)
}

func (c Page) Int32() int32 {
	return int32(c)
}

type Pages []int32

func NewPages(v []int32) Pages {
	return Pages(v)
}

func (c Pages) Int32s() []int32 {
	return []int32(c)
}

type Shop Shopa

func NewShop(v Shopa) Shop {
	return Shop(v)
}

func (c Shop) Shopa() Shopa {
	return Shopa(c)
}

type Shopa Shopb

func NewShopa(v Shopb) Shopa {
	return Shopa(v)
}

func (c Shopa) Shopb() Shopb {
	return Shopb(c)
}

type Shopb Shopc

func NewShopb(v Shopc) Shopb {
	return Shopb(v)
}

func (c Shopb) Shopc() Shopc {
	return Shopc(c)
}

type Shopc string

func NewShopc(v string) Shopc {
	return Shopc(v)
}

func (c Shopc) String() string {
	return string(c)
}

type Shops []string

func NewShops(v []string) Shops {
	return Shops(v)
}

func (c Shops) Strings() []string {
	return []string(c)
}
