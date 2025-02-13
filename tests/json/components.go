package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ------------------------
//         Schemas
// ------------------------

type Metadata struct {
	Owner string
	Tags  Maybe[Tags]
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
	{
		var v any
		v = c.Owner
		writeProperty("owner", v)
	}
	if vOpt, ok := c.Tags.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("tags", v)
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
	if raw, ok := m["owner"]; ok {
		err = json.Unmarshal(raw, &c.Owner)
		if err != nil {
			return fmt.Errorf("'owner' field: unmarshal string: %w", err)
		}
		delete(m, "owner")
	} else {
		return fmt.Errorf("'owner' key is missing")
	}
	if raw, ok := m["tags"]; ok {
		err = c.Tags.Value.UnmarshalJSON(raw)
		if err != nil {
			return fmt.Errorf("'tags' field: unmarshal ref type 'Tags': %w", err)
		}
		c.Tags.IsSet = true
		delete(m, "tags")
	}
	return nil
}

type NewPet struct {
	Birthday time.Time
	Metadata Maybe[Metadata]
	Name     string
	Tag      Nullable[string]
	Tago     Maybe[Nullable[string]]
}

var _ json.Marshaler = (*NewPet)(nil)

func (c NewPet) MarshalJSON() ([]byte, error) {
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

func (c NewPet) marshalJSONInnerBody(out io.Writer) error {
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
		v = c.Birthday.Format(time.RFC3339Nano)
		writeProperty("birthday", v)
	}
	if vOpt, ok := c.Metadata.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("metadata", v)
	}
	{
		var v any
		v = c.Name
		writeProperty("name", v)
	}
	{
		var v any
		if vPtr, ok := c.Tag.Get(); ok {
			v = vPtr
		}
		writeProperty("tag", v)
	}
	if vOpt, ok := c.Tago.Get(); ok {
		var v any = nil
		if vPtr, ok := vOpt.Get(); ok {
			v = vPtr
		}
		writeProperty("tago", v)
	}

	return err
}

var _ json.Unmarshaler = (*NewPet)(nil)

func (c *NewPet) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *NewPet) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["birthday"]; ok {
		var s string
		err = json.Unmarshal(raw, &s)
		if err != nil {
			return fmt.Errorf("'birthday' field: unmarshal string: %w", err)
		}
		c.Birthday, err = time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return fmt.Errorf("'birthday' field: parse time: %w", err)
		}
		delete(m, "birthday")
	} else {
		return fmt.Errorf("'birthday' key is missing")
	}
	if raw, ok := m["metadata"]; ok {
		err = c.Metadata.Value.UnmarshalJSON(raw)
		if err != nil {
			return fmt.Errorf("'metadata' field: unmarshal ref type 'Metadata': %w", err)
		}
		c.Metadata.IsSet = true
		delete(m, "metadata")
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
		var vn Nullable[string]
		if len(raw) != 4 || string(raw) != "null" {
			var v string
			err = json.Unmarshal(raw, &v)
			if err != nil {
				return fmt.Errorf("'tag' field: unmarshal string: %w", err)
			}
			vn.Set(v)
		}
		c.Tag = vn
		delete(m, "tag")
	} else {
		return fmt.Errorf("'tag' key is missing")
	}
	if raw, ok := m["tago"]; ok {
		var vn Nullable[string]
		if len(raw) != 4 || string(raw) != "null" {
			var v string
			err = json.Unmarshal(raw, &v)
			if err != nil {
				return fmt.Errorf("'tago' field: unmarshal string: %w", err)
			}
			vn.Set(v)
		}
		c.Tago.Value = vn
		c.Tago.IsSet = true
		delete(m, "tago")
	}
	return nil
}

type Pet struct {
	NewPet
	ID int64
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
		mErr := c.NewPet.marshalJSONInnerBody(out)
		if mErr != nil {
			err = mErr
		}
	}
	comma = ","
	{
		var v any
		v = c.ID
		writeProperty("id", v)
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
	{
		var v NewPet
		err := v.unmarshalJSONInnerBody(m)
		if err != nil {
			return fmt.Errorf("embedded 'NewPet' field: unmarshal schema: %w", err)
		}
		c.NewPet = v
	}
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
	return nil
}

type Tag struct {
	Name  string
	Value string
}

var _ json.Marshaler = (*Tag)(nil)

func (c Tag) MarshalJSON() ([]byte, error) {
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

func (c Tag) marshalJSONInnerBody(out io.Writer) error {
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

var _ json.Unmarshaler = (*Tag)(nil)

func (c *Tag) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Tag) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
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

type Tags []Tag

var _ json.Marshaler = (*Tags)(nil)

func (c Tags) MarshalJSON() ([]byte, error) {
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

func (c Tags) marshalJSONInnerBody(out io.Writer) error {
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

var _ json.Unmarshaler = (*Tags)(nil)

func (c *Tags) UnmarshalJSON(bs []byte) error {
	var m []json.RawMessage
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Tags) unmarshalJSONInnerBody(m []json.RawMessage) error {
	out := make(Tags, 0, len(m))
	var err error
	_ = err
	for _, vm := range m {
		var vItem Tag
		err = vItem.UnmarshalJSON(vm)
		if err != nil {
			return fmt.Errorf("unmarshal ref type 'Tag': %w", err)
		}
		out = append(out, vItem)
	}
	*c = out
	return nil
}

// ------------------------------
//         Responses
// ------------------------------

func NewPetResponse(body Pet) PetResponse {
	var out PetResponse
	out.Body = body
	return out
}

// PetResponse - Pet output response
type PetResponse struct {
	Body Pet
}

func (r PetResponse) writePostPets(w http.ResponseWriter) {
	r.Write(w, 200)
}

func (r PetResponse) Write(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, r.Body, "PetResponse")
}
