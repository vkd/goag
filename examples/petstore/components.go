// Code generated by goag (https://github.com/vkd/goag). DO NOT EDIT.
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
	Code    int32
	Message string
}

var _ json.Marshaler = (*Error)(nil)

func (c Error) MarshalJSON() ([]byte, error) {
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

func (c Error) marshalJSONInnerBody(out io.Writer) error {
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
		v = c.Code
		writeProperty("code", v)
	}
	{
		var v any
		v = c.Message
		writeProperty("message", v)
	}

	return err
}

var _ json.Unmarshaler = (*Error)(nil)

func (c *Error) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Error) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["code"]; ok {
		var v int32
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return fmt.Errorf("'code' field: unmarshal int32: %w", err)
		}
		c.Code = v
		delete(m, "code")
	} else {
		return fmt.Errorf("'code' key is missing")
	}
	if raw, ok := m["message"]; ok {
		err = json.Unmarshal(raw, &c.Message)
		if err != nil {
			return fmt.Errorf("'message' field: unmarshal string: %w", err)
		}
		delete(m, "message")
	} else {
		return fmt.Errorf("'message' key is missing")
	}
	return nil
}

type Pet struct {
	ID   int64
	Name string
	Tag  Maybe[string]
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
		v = c.ID
		writeProperty("id", v)
	}
	{
		var v any
		v = c.Name
		writeProperty("name", v)
	}
	if vOpt, ok := c.Tag.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("tag", v)
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
	_ = err
	if raw, ok := m["id"]; ok {
		var v int64
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return fmt.Errorf("'id' field: unmarshal int64: %w", err)
		}
		c.ID = v
		delete(m, "id")
	} else {
		return fmt.Errorf("'id' key is missing")
	}
	if raw, ok := m["name"]; ok {
		err = json.Unmarshal(raw, &c.Name)
		if err != nil {
			return fmt.Errorf("'name' field: unmarshal string: %w", err)
		}
		delete(m, "name")
	} else {
		return fmt.Errorf("'name' key is missing")
	}
	if raw, ok := m["tag"]; ok {
		err = json.Unmarshal(raw, &c.Tag.Value)
		if err != nil {
			return fmt.Errorf("'tag' field: unmarshal string: %w", err)
		}
		c.Tag.IsSet = true
		delete(m, "tag")
	}
	return nil
}

type Pets []Pet

var _ json.Marshaler = (*Pets)(nil)

func (c Pets) MarshalJSON() ([]byte, error) {
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

func (c Pets) marshalJSONInnerBody(out io.Writer) error {
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

		writeItem(cv)
	}

	return err
}

var _ json.Unmarshaler = (*Pets)(nil)

func (c *Pets) UnmarshalJSON(bs []byte) error {
	var m []json.RawMessage
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Pets) unmarshalJSONInnerBody(m []json.RawMessage) error {
	out := make(Pets, 0, len(m))
	var err error
	_ = err
	for _, vm := range m {
		var vItem Pet
		err = vItem.UnmarshalJSON(vm)
		if err != nil {
			return fmt.Errorf("unmarshal ref type 'Pet': %w", err)
		}
		out = append(out, vItem)
	}
	*c = out
	return nil
}
