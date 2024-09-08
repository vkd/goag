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

type Pet struct {
	Custom               PetCustom                  `json:"custom"`
	Name                 string                     `json:"name"`
	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

var _ json.Marshaler = (*Pet)(nil)

func (c Pet) MarshalJSON() ([]byte, error) {
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

func (c Pet) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
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
	writeJSON := func(v any) {
		if err != nil {
			return
		}
		werr := encoder.Encode(v)
		if werr != nil {
			err = werr
		}
	}
	_ = writeJSON
	write([]byte(`"custom":`))
	writeJSON(c.Custom)
	write([]byte(`,`))
	write([]byte(`"name":`))
	writeJSON(c.Name)
	write([]byte(`,`))

	for k, v := range c.AdditionalProperties {
		write([]byte(`"` + k + `":`))
		writeJSON(v)
	}

	return err
}

var _ json.Unmarshaler = (*Pet)(nil)

func (c *Pet) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Pet) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	if raw, ok := m["custom"]; ok {
		var v PetCustom
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return fmt.Errorf("'custom' field: %w", err)
		}
		c.Custom = v
		delete(m, "custom")
	}
	if raw, ok := m["name"]; ok {
		var v string
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return fmt.Errorf("'name' field: %w", err)
		}
		c.Name = v
		delete(m, "name")
	}
	for k, bs := range m {
		var v json.RawMessage
		err = json.Unmarshal(bs, &v)
		if err != nil {
			return fmt.Errorf("additional property %q: %w", k, err)
		}
		c.AdditionalProperties[k] = v
	}
	return nil
}

type PetCustom = json.RawMessage

type Pets []Pet
