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
// GetPetsIDs -
// GET /pets/ids
// ---------------------------------------------

type GetPetsIDsHandlerFunc func(ctx context.Context, r GetPetsIDsRequest) GetPetsIDsResponse

func (f GetPetsIDsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsIDsHTTPRequest(r)).writeGetPetsIDs(w)
}

func (GetPetsIDsHandlerFunc) Path() string { return "/pets/ids" }

func (GetPetsIDsHandlerFunc) Method() string { return http.MethodGet }

type GetPetsIDsRequest interface {
	HTTP() *http.Request
	Parse() GetPetsIDsParams
}

func GetPetsIDsHTTPRequest(r *http.Request) GetPetsIDsRequest {
	return getPetsIDsHTTPRequest{r}
}

type getPetsIDsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsIDsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsIDsHTTPRequest) Parse() GetPetsIDsParams {
	return newGetPetsIDsParams(r.Request)
}

type GetPetsIDsParams struct {
}

func newGetPetsIDsParams(r *http.Request) (zero GetPetsIDsParams) {
	var params GetPetsIDsParams

	return params
}

func (r GetPetsIDsParams) HTTP() *http.Request { return nil }

func (r GetPetsIDsParams) Parse() GetPetsIDsParams { return r }

type GetPetsIDsResponse interface {
	writeGetPetsIDs(http.ResponseWriter)
}

func NewGetPetsIDsResponse200JSON(body []float64) GetPetsIDsResponse {
	var out GetPetsIDsResponse200JSON
	out.Body = body
	return out
}

type GetPetsIDsResponse200JSON struct {
	Body []float64
}

func (r GetPetsIDsResponse200JSON) writeGetPetsIDs(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsIDsResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsIDsResponse200JSON")
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
