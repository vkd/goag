package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// ------------------------
//         Schemas
// ------------------------

type Organization int

func NewOrganization(v int) Organization { return Organization(v) }

func (c Organization) Int() int { return int(c) }

type Page int32

func NewPage(v int32) Page { return Page(v) }

func (c Page) Int32() int32 { return int32(c) }

type Pages []int32

func NewPages(v []int32) Pages { return Pages(v) }

func (c Pages) Int32s() []int32 { return []int32(c) }

var _ json.Marshaler = (*Pages)(nil)

func (c Pages) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`[`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`]`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Pages) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeItem := func(v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(`null`)
		} else {
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
	}
	_ = writeItem

	for i, cv := range c {
		_ = i
		if err != nil {
			return err
		}

		write(comma)
		comma = ","

		var v int32
		v = cv
		writeItem(v)
	}

	return err
}

var _ json.Unmarshaler = (*Pages)(nil)

func (c *Pages) UnmarshalJSON(bs []byte) error {
	var m []json.RawMessage
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Pages) unmarshalJSONInnerBody(m []json.RawMessage) error {
	out := make(Pages, 0, len(m))
	var err error
	_ = err
	for _, vm := range m {
		var v int32
		err = json.Unmarshal(vm, &v)
		if err != nil {
			return fmt.Errorf("unmarshal int32: %w", err)
		}
		vItem := v
		out = append(out, vItem)
	}
	*c = out
	return nil
}

type Shop Shopa

func NewShop(v Shopa) Shop { return Shop(v) }

func (c Shop) Shopa() Shopa { return Shopa(c) }

type Shopa Shopb

func NewShopa(v Shopb) Shopa { return Shopa(v) }

func (c Shopa) Shopb() Shopb { return Shopb(c) }

type Shopb Shopc

func NewShopb(v Shopc) Shopb { return Shopb(v) }

func (c Shopb) Shopc() Shopc { return Shopc(c) }

type Shopc string

func NewShopc(v string) Shopc { return Shopc(v) }

func (c Shopc) String() string { return string(c) }

type Shops []string

func NewShops(v []string) Shops { return Shops(v) }

func (c Shops) Strings() []string { return []string(c) }

var _ json.Marshaler = (*Shops)(nil)

func (c Shops) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`[`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`]`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Shops) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeItem := func(v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(`null`)
		} else {
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
	}
	_ = writeItem

	for i, cv := range c {
		_ = i
		if err != nil {
			return err
		}

		write(comma)
		comma = ","

		var v string
		v = cv
		writeItem(v)
	}

	return err
}

var _ json.Unmarshaler = (*Shops)(nil)

func (c *Shops) UnmarshalJSON(bs []byte) error {
	var m []json.RawMessage
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Shops) unmarshalJSONInnerBody(m []json.RawMessage) error {
	out := make(Shops, 0, len(m))
	var err error
	_ = err
	for _, vm := range m {
		var vItem string
		err = json.Unmarshal(vm, &vItem)
		if err != nil {
			return fmt.Errorf("unmarshal string: %w", err)
		}
		out = append(out, vItem)
	}
	*c = out
	return nil
}
