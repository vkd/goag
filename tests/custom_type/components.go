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

type Environment struct {
	pkg.EnvironmentCreate
	Value string
}

var _ json.Marshaler = (*Environment)(nil)

func (c Environment) MarshalJSON() ([]byte, error) {
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

func (c Environment) marshalJSONInnerBody(out io.Writer) error {
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
		var v EnvironmentCreate
		var vc pkg.EnvironmentCreate
		vc = c.EnvironmentCreate
		v = EnvironmentCreate(vc.ToSchemaEnvironmentCreate())
		mErr := v.marshalJSONInnerBody(out)
		if mErr != nil {
			err = mErr
		}
	}
	comma = ","
	{
		var v any
		v = c.Value
		writeProperty("value", v)
	}

	return err
}

var _ json.Unmarshaler = (*Environment)(nil)

func (c *Environment) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Environment) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	{
		var v EnvironmentCreate
		err := v.unmarshalJSONInnerBody(m)
		if err != nil {
			return fmt.Errorf("embedded 'EnvironmentCreate' field: unmarshal schema: %w", err)
		}
		err = c.EnvironmentCreate.SetFromSchemaEnvironmentCreate(v)
		if err != nil {
			return fmt.Errorf("embedded 'EnvironmentCreate' field: set custom type from schema: %w", err)
		}
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

type EnvironmentCreate struct {
	Name string
}

var _ json.Marshaler = (*EnvironmentCreate)(nil)

func (c EnvironmentCreate) MarshalJSON() ([]byte, error) {
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

func (c EnvironmentCreate) marshalJSONInnerBody(out io.Writer) error {
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

	return err
}

var _ json.Unmarshaler = (*EnvironmentCreate)(nil)

func (c *EnvironmentCreate) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *EnvironmentCreate) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
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
	return nil
}

type Environments []pkg.Environment

var _ json.Marshaler = (*Environments)(nil)

func (c Environments) MarshalJSON() ([]byte, error) {
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

func (c Environments) marshalJSONInnerBody(out io.Writer) error {
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

		var v struct {
			pkg.EnvironmentCreate
			Value string
		}
		v = cv.ToSchemaEnvironment()
		vItem, err := Environment(v).MarshalJSON()
		if err != nil {
			return fmt.Errorf("marshal %d element: %v", i, err)
		}
		writeItem(json.RawMessage(vItem))
	}

	return err
}

var _ json.Unmarshaler = (*Environments)(nil)

func (c *Environments) UnmarshalJSON(bs []byte) error {
	var m []json.RawMessage
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Environments) unmarshalJSONInnerBody(m []json.RawMessage) error {
	out := make(Environments, 0, len(m))
	var err error
	_ = err
	for _, vm := range m {
		var vItem pkg.Environment

		var vRef Environment
		err = vRef.UnmarshalJSON(vm)
		if err != nil {
			return fmt.Errorf("unmarshal ref type 'Environment': %w", err)
		}

		var cv pkg.Environment
		err = cv.SetFromSchemaEnvironment(struct {
			pkg.EnvironmentCreate
			Value string
		}(vRef))
		if err != nil {
			return fmt.Errorf("set from schema: %w", err)
		}
		vItem = cv
		out = append(out, vItem)
	}
	*c = out
	return nil
}

type GetShop struct {
	Additionals  pkg.Maybe[pkg.Nullable[pkg.Settings]]
	Environments pkg.Maybe[pkg.Nullable[pkg.Environments]]
	Metadata     pkg.Metadata
	Settings     pkg.Maybe[pkg.Nullable[pkg.Settings]]
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
	if vOpt, ok := c.Additionals.Get(); ok {
		var v any = nil
		if vPtr, ok := vOpt.Get(); ok {
			var vc pkg.Settings
			vc = vPtr
			v = GetShopAdditionals(vc.ToSchemaGetShopAdditionals())
		}
		writeProperty("additionals", v)
	}
	if vOpt, ok := c.Environments.Get(); ok {
		var v any = nil
		if vPtr, ok := vOpt.Get(); ok {
			var vc pkg.Environments
			vc = vPtr
			v = Environments(vc.ToSchemaEnvironments())
		}
		writeProperty("environments", v)
	}
	{
		var v any
		var vc pkg.Metadata
		vc = c.Metadata
		v = Metadata(vc.ToSchemaMetadata())
		writeProperty("metadata", v)
	}
	if vOpt, ok := c.Settings.Get(); ok {
		var v any = nil
		if vPtr, ok := vOpt.Get(); ok {
			var vc pkg.Settings
			vc = vPtr
			v = Settings(vc.ToSchemaSettings())
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
	if raw, ok := m["additionals"]; ok {
		var vNull pkg.Nullable[pkg.Settings]
		if len(raw) != 4 || string(raw) != "null" {
			var vRef GetShopAdditionals
			err = vRef.UnmarshalJSON(raw)
			if err != nil {
				return fmt.Errorf("'additionals' field: unmarshal ref type 'GetShopAdditionals': %w", err)
			}

			var cv pkg.Settings
			err = cv.SetFromSchemaGetShopAdditionals(struct {
				AdditionalProperties map[string]any
			}(vRef))
			if err != nil {
				return fmt.Errorf("'additionals' field: set from schema: %w", err)
			}
			vNull.Set(cv)
		}
		c.Additionals.Value = vNull
		c.Additionals.IsSet = true
		delete(m, "additionals")
	}
	if raw, ok := m["environments"]; ok {
		var vNull pkg.Nullable[pkg.Environments]
		if len(raw) != 4 || string(raw) != "null" {
			var vRef Environments
			err = vRef.UnmarshalJSON(raw)
			if err != nil {
				return fmt.Errorf("'environments' field: unmarshal ref type 'Environments': %w", err)
			}

			var cv pkg.Environments
			err = cv.SetFromSchemaEnvironments([]pkg.Environment(vRef))
			if err != nil {
				return fmt.Errorf("'environments' field: set from schema: %w", err)
			}
			vNull.Set(cv)
		}
		c.Environments.Value = vNull
		c.Environments.IsSet = true
		delete(m, "environments")
	}
	if raw, ok := m["metadata"]; ok {

		var vRef Metadata
		err = vRef.UnmarshalJSON(raw)
		if err != nil {
			return fmt.Errorf("'metadata' field: unmarshal ref type 'Metadata': %w", err)
		}

		var cv pkg.Metadata
		err = cv.SetFromSchemaMetadata(struct {
			InnerID pkg.Maybe[string]
		}(vRef))
		if err != nil {
			return fmt.Errorf("'metadata' field: set from schema: %w", err)
		}
		c.Metadata = cv
		delete(m, "metadata")
	} else {
		return fmt.Errorf("'metadata' key is missing")
	}
	if raw, ok := m["settings"]; ok {
		var vNull pkg.Nullable[pkg.Settings]
		if len(raw) != 4 || string(raw) != "null" {
			var vRef Settings
			err = vRef.UnmarshalJSON(raw)
			if err != nil {
				return fmt.Errorf("'settings' field: unmarshal ref type 'Settings': %w", err)
			}

			var cv pkg.Settings
			err = cv.SetFromSchemaSettings(struct {
				Theme pkg.Maybe[string]
			}(vRef))
			if err != nil {
				return fmt.Errorf("'settings' field: set from schema: %w", err)
			}
			vNull.Set(cv)
		}
		c.Settings.Value = vNull
		c.Settings.IsSet = true
		delete(m, "settings")
	}
	return nil
}

type Metadata struct {
	InnerID pkg.Maybe[string]
}

var _ json.Marshaler = (*Metadata)(nil)

func (c Metadata) MarshalJSON() ([]byte, error) {
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

func (c Metadata) marshalJSONInnerBody(out io.Writer) error {
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

var _ json.Unmarshaler = (*Metadata)(nil)

func (c *Metadata) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Metadata) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
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

type PageCustom string

func NewPageCustom(v string) PageCustom {
	return PageCustom(v)
}

func (c PageCustom) String() string {
	return string(c)
}

type Settings struct {
	Theme pkg.Maybe[string]
}

var _ json.Marshaler = (*Settings)(nil)

func (c Settings) MarshalJSON() ([]byte, error) {
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

func (c Settings) marshalJSONInnerBody(out io.Writer) error {
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

var _ json.Unmarshaler = (*Settings)(nil)

func (c *Settings) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Settings) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
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

type Shop ShopName

func NewShop(v ShopName) Shop {
	return Shop(v)
}

func (c Shop) ShopName() ShopName {
	return ShopName(c)
}

type ShopName int64

func NewShopName(v int64) ShopName {
	return ShopName(v)
}

func (c ShopName) Int64() int64 {
	return int64(c)
}

type GetShopAdditionals struct {
	AdditionalProperties map[string]any
}

var _ json.Marshaler = (*GetShopAdditionals)(nil)

func (c GetShopAdditionals) MarshalJSON() ([]byte, error) {
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

func (c GetShopAdditionals) marshalJSONInnerBody(out io.Writer) error {
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

var _ json.Unmarshaler = (*GetShopAdditionals)(nil)

func (c *GetShopAdditionals) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *GetShopAdditionals) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if len(m) > 0 {
		c.AdditionalProperties = make(map[string]any)
	}
	for k, bs := range m {
		var v any
		err = json.Unmarshal(bs, &v)
		if err != nil {
			return fmt.Errorf("additional property %q: %w", k, err)
		}
		c.AdditionalProperties[k] = v
	}
	return nil
}
