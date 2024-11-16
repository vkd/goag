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

type Error struct {
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
	return nil
}

// ------------------------------
//         Responses
// ------------------------------

func NewErrorResponse(code int, body Error) ErrorResponse {
	var out ErrorResponse
	out.Code = code
	out.Body = body
	return out
}

// ErrorResponse - Error output response
type ErrorResponse struct {
	Code int
	Body Error
}

func (r ErrorResponse) writeGetV2Pet(w http.ResponseWriter) {
	r.Write(w)
}

func (r ErrorResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "ErrorResponse")
}

func NewPetResponse(body Pet) PetResponse {
	var out PetResponse
	out.Body = body
	return out
}

// PetResponse - Pet output response
type PetResponse struct {
	Body Pet
}

func (r PetResponse) writeGetPet(w http.ResponseWriter) {
	r.Write(w, 200)
}

func (r PetResponse) writeGetV2Pet(w http.ResponseWriter) {
	r.Write(w, 201)
}

func (r PetResponse) writeGetV3Pet(w http.ResponseWriter) {
	r.Write(w, 202)
}

func (r PetResponse) Write(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, r.Body, "PetResponse")
}

// NewPet2Response - Pet output response
func NewPet2Response(body Pet) Pet2Response {
	return NewPetResponse(body)
}

// Pet2Response - Pet output response
type Pet2Response = PetResponse

// NewPet3Response - Pet output response
func NewPet3Response(body Pet) Pet3Response {
	return NewPet2Response(body)
}

// Pet3Response - Pet output response
type Pet3Response = Pet2Response
