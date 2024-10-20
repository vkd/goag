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

type Environment = pkg.Environment

type Environments = pkg.Nullable[[]Environment]

type GetShop struct {
	Environments pkg.Maybe[Environments] `json:"environments"`
	Metadata     Metadata                `json:"metadata"`
	Settings     pkg.Maybe[Settings]     `json:"settings"`
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
	if vOpt, ok := c.Environments.Get(); ok {
		var v any = nil
		if vPtr, ok := vOpt.Get(); ok {
			v = vPtr
		}
		writeProperty("environments", v)
	}
	{
		var v any
		v = c.Metadata
		writeProperty("metadata", v)
	}
	if vOpt, ok := c.Settings.Get(); ok {
		var v any = nil
		if vPtr, ok := vOpt.Get(); ok {
			v = vPtr
		}
		writeProperty("settings", v)
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
	_ = err
	if raw, ok := m["environments"]; ok {
		if string(raw) != "null" {
			var vs []Environment
			err = json.Unmarshal(raw, &vs)
			if err != nil {
				return fmt.Errorf("'environments' field: unmarshal slice: %w", err)
			}
			v := vs
			var vPtr pkg.Nullable[[]Environment]
			vPtr.Set(v)
			c.Environments.Value = vPtr
		}
		c.Environments.IsSet = true
		delete(m, "environments")
	}
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
	if raw, ok := m["settings"]; ok {
		if string(raw) != "null" {
			var cv pkg.Settings
			err = cv.UnmarshalJSON(raw)
			if err != nil {
				return fmt.Errorf("'settings' field: unmarshal object: %w", err)
			}
			v := cv
			var vPtr pkg.Nullable[pkg.Settings]
			vPtr.Set(v)
			c.Settings.Value = vPtr
		}
		c.Settings.IsSet = true
		delete(m, "settings")
	}
	return nil
}

type Metadata = pkg.Metadata

type PageCustom = pkg.Page

type Settings = pkg.Nullable[pkg.Settings]

type Shop ShopName

func NewShop(v ShopName) Shop {
	return Shop(v)
}

func (c Shop) ShopName() ShopName {
	return ShopName(c)
}

type ShopName = pkg.Page

type EnvironmentSchema struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var _ json.Marshaler = (*EnvironmentSchema)(nil)

func (c EnvironmentSchema) MarshalJSON() ([]byte, error) {
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

func (c EnvironmentSchema) marshalJSONInnerBody(out io.Writer) error {
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
		v = c.Name
		writeProperty("name", v)
	}
	{
		var v any
		v = c.Value
		writeProperty("value", v)
	}

	return err
}

var _ json.Unmarshaler = (*EnvironmentSchema)(nil)

func (c *EnvironmentSchema) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *EnvironmentSchema) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["name"]; ok {
		err = json.Unmarshal(raw, &c.Name)
		if err != nil {
			return fmt.Errorf("'name' field: unmarshal string: %w", err)
		}
		delete(m, "name")
	} else {
		return fmt.Errorf("'name' key is missing")
	}
	if raw, ok := m["value"]; ok {
		err = json.Unmarshal(raw, &c.Value)
		if err != nil {
			return fmt.Errorf("'value' field: unmarshal string: %w", err)
		}
		delete(m, "value")
	} else {
		return fmt.Errorf("'value' key is missing")
	}
	return nil
}

type MetadataSchema struct {
	InnerID pkg.Maybe[string] `json:"inner_id"`
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
	_ = err
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

type SettingsSchema struct {
	Theme pkg.Maybe[string] `json:"theme"`
}

var _ json.Marshaler = (*SettingsSchema)(nil)

func (c SettingsSchema) MarshalJSON() ([]byte, error) {
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

func (c SettingsSchema) marshalJSONInnerBody(out io.Writer) error {
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
	if vOpt, ok := c.Theme.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("theme", v)
	}

	return err
}

var _ json.Unmarshaler = (*SettingsSchema)(nil)

func (c *SettingsSchema) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *SettingsSchema) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["theme"]; ok {
		err = json.Unmarshal(raw, &c.Theme.Value)
		if err != nil {
			return fmt.Errorf("'theme' field: unmarshal string: %w", err)
		}
		c.Theme.IsSet = true
		delete(m, "theme")
	}
	return nil
}

type ShopNameSchema int64

func NewShopNameSchema(v int64) ShopNameSchema {
	return ShopNameSchema(v)
}

func (c ShopNameSchema) Int64() int64 {
	return int64(c)
}
