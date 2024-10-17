package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vkd/goag/tests/custom_type/pkg"
)

// ------------------------
//         Schemas
// ------------------------

type GetShop struct {
	Metadata Metadata `json:"metadata"`
}

var _ json.Marshaler = (*GetShop)(nil)

func (c GetShop) MarshalJSON() ([]byte, error) {
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

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c GetShop) marshalJSONInnerBody(out io.Writer) error {
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
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	{
		var v any
		v = c.Metadata
		writeProperty("metadata", v)
	}

	return err
}

var _ json.Unmarshaler = (*GetShop)(nil)

func (c *GetShop) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *GetShop) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	if raw, ok := m["metadata"]; ok {
		var cv pkg.Metadata
		err = cv.UnmarshalJSON(raw)
		if err != nil {
			return fmt.Errorf("'metadata' field: unmarshal object: %w", err)
		}
		c.Metadata = cv
		delete(m, "metadata")
	} else {
		return fmt.Errorf("'metadata' key is missing")
	}
	return nil
}

type Metadata = pkg.Metadata

type PageCustom = pkg.Page

type Shop ShopName

func NewShop(v ShopName) Shop {
	return Shop(v)
}

func (c Shop) ShopName() ShopName {
	return ShopName(c)
}

type ShopName = pkg.Page

type MetadataSchema struct {
	InnerID Maybe[string] `json:"inner_id"`
}

var _ json.Marshaler = (*MetadataSchema)(nil)

func (c MetadataSchema) MarshalJSON() ([]byte, error) {
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

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c MetadataSchema) marshalJSONInnerBody(out io.Writer) error {
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
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	if vOpt, ok := c.InnerID.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("inner_id", v)
	}

	return err
}

var _ json.Unmarshaler = (*MetadataSchema)(nil)

func (c *MetadataSchema) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *MetadataSchema) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	if raw, ok := m["inner_id"]; ok {
		err = json.Unmarshal(raw, &c.InnerID.Value)
		if err != nil {
			return fmt.Errorf("'inner_id' field: unmarshal string: %w", err)
		}
		c.InnerID.IsSet = true
		delete(m, "inner_id")
	}
	return nil
}

type PageCustomSchema string

func NewPageCustomSchema(v string) PageCustomSchema {
	return PageCustomSchema(v)
}

func (c PageCustomSchema) String() string {
	return string(c)
}

type ShopNameSchema int64

func NewShopNameSchema(v int64) ShopNameSchema {
	return ShopNameSchema(v)
}

func (c ShopNameSchema) Int64() int64 {
	return int64(c)
}
