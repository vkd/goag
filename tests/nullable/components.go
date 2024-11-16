package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ------------------------
//         Schemas
// ------------------------

type NewPet struct {
	Name string
	Tag  Nullable[string]
	Tago Maybe[Nullable[string]]
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
	ID      int64
	Name    string
	Parents Nullable[PetParents]
	Tag     Nullable[string]
	Tago    Maybe[Nullable[string]]
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
	{
		var v any
		if vPtr, ok := c.Parents.Get(); ok {
			v = vPtr
		}
		writeProperty("parents", v)
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
	if raw, ok := m["parents"]; ok {
		var vn Nullable[PetParents]
		if len(raw) != 4 || string(raw) != "null" {
			var v PetParents
			err = v.UnmarshalJSON(raw)
			if err != nil {
				return fmt.Errorf("'parents' field: unmarshal nullable ref type 'PetParents': %w", err)
			}
			vn.Set(v)
		}
		c.Parents = vn
		delete(m, "parents")
	} else {
		return fmt.Errorf("'parents' key is missing")
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

type PetParents struct {
	AdditionalProperties map[string]any
}

var _ json.Marshaler = (*PetParents)(nil)

func (c PetParents) MarshalJSON() ([]byte, error) {
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

func (c PetParents) marshalJSONInnerBody(out io.Writer) error {
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

var _ json.Unmarshaler = (*PetParents)(nil)

func (c *PetParents) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *PetParents) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
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
