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

type Error struct {
	Message string `json:"message"`
}

type NewPet struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Pet struct {
	NewPet
	ID int64 `json:"id"`
}

type Pets []Pet

type GetShopsShopPetsResponse200JSONBodyGroups struct {
	AdditionalProperties map[string]Pets `json:"-"`
}

var _ json.Marshaler = (*GetShopsShopPetsResponse200JSONBodyGroups)(nil)

func (c GetShopsShopPetsResponse200JSONBodyGroups) MarshalJSON() ([]byte, error) {
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

func (c GetShopsShopPetsResponse200JSONBodyGroups) marshalJSONInnerBody(out io.Writer) error {
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
	for k, v := range c.AdditionalProperties {
		writeProperty(k, v)
	}

	return err
}

var _ json.Unmarshaler = (*GetShopsShopPetsResponse200JSONBodyGroups)(nil)

func (c *GetShopsShopPetsResponse200JSONBodyGroups) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *GetShopsShopPetsResponse200JSONBodyGroups) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	for k, bs := range m {
		var v Pets
		err = json.Unmarshal(bs, &v)
		if err != nil {
			return fmt.Errorf("additional property %q: %w", k, err)
		}
		c.AdditionalProperties[k] = v
	}
	return nil
}
