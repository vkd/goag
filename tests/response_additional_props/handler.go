package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPet -
// GET /pet
// ---------------------------------------------

type GetPetHandlerFunc func(ctx context.Context, r GetPetRequest) GetPetResponse

func (f GetPetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetHTTPRequest(r)).writeGetPet(w)
}

func (GetPetHandlerFunc) Path() string { return "/pet" }

func (GetPetHandlerFunc) Method() string { return http.MethodGet }

type GetPetRequest interface {
	HTTP() *http.Request
	Parse() GetPetParams
}

func GetPetHTTPRequest(r *http.Request) GetPetRequest {
	return getPetHTTPRequest{r}
}

type getPetHTTPRequest struct {
	Request *http.Request
}

func (r getPetHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetHTTPRequest) Parse() GetPetParams {
	return newGetPetParams(r.Request)
}

type GetPetParams struct {
}

func newGetPetParams(r *http.Request) (zero GetPetParams) {
	var params GetPetParams

	return params
}

func (r GetPetParams) HTTP() *http.Request { return nil }

func (r GetPetParams) Parse() GetPetParams { return r }

type GetPetResponse interface {
	writeGetPet(http.ResponseWriter)
}

func NewGetPetResponse200JSON(body GetPetResponse200JSONBody) GetPetResponse {
	var out GetPetResponse200JSON
	out.Body = body
	return out
}

type GetPetResponse200JSONBody struct {
	Groups GetPetResponse200JSONBodyGroups
}

var _ json.Marshaler = (*GetPetResponse200JSONBody)(nil)

func (c GetPetResponse200JSONBody) MarshalJSON() ([]byte, error) {
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

func (c GetPetResponse200JSONBody) marshalJSONInnerBody(out io.Writer) error {
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
		v = c.Groups
		writeProperty("groups", v)
	}

	return err
}

var _ json.Unmarshaler = (*GetPetResponse200JSONBody)(nil)

func (c *GetPetResponse200JSONBody) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *GetPetResponse200JSONBody) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["groups"]; ok {
		err = c.Groups.UnmarshalJSON(raw)
		if err != nil {
			return fmt.Errorf("'groups' field: unmarshal ref type 'GetPetResponse200JSONBodyGroups': %w", err)
		}
		delete(m, "groups")
	} else {
		return fmt.Errorf("'groups' key is missing")
	}
	return nil
}

type GetPetResponse200JSON struct {
	Body GetPetResponse200JSONBody
}

func (r GetPetResponse200JSON) writeGetPet(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetResponse200JSON")
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

func writeJSON(w io.Writer, v interface{}, name string) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		LogError(fmt.Errorf("write json response %q: %w", name, err))
	}
}

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

type Nullable[T any] struct {
	IsSet bool
	Value T
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{}
}

func Pointer[T any](v T) Nullable[T] {
	return Nullable[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Nullable[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Nullable[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

var _ json.Marshaler = (*Nullable[any])(nil)

func (m Nullable[T]) MarshalJSON() ([]byte, error) {
	if m.IsSet {
		return json.Marshal(&m.Value)
	}
	return []byte(nullValue), nil
}

var _ json.Unmarshaler = (*Nullable[any])(nil)

const nullValue = "null"

var nullValueBs = []byte(nullValue)

func (m *Nullable[T]) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, nullValueBs) {
		m.IsSet = false
		return nil
	}
	m.IsSet = true
	return json.Unmarshal(bs, &m.Value)
}

type ErrParseParam struct {
	In        string
	Parameter string
	Reason    string
	Err       error
}

func (e ErrParseParam) Error() string {
	return fmt.Sprintf("%s parameter '%s': %s: %v", e.In, e.Parameter, e.Reason, e.Err)
}

func (e ErrParseParam) Unwrap() error { return e.Err }
