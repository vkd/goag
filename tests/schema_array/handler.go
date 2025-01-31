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
// GetPets -
// ---------------------------------------------

type GetPetsHandlerFunc func(ctx context.Context, r GetPetsRequest) GetPetsResponse

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsHTTPRequest(r)).writeGetPets(w)
}

type GetPetsRequest interface {
	HTTP() *http.Request
	Parse() GetPetsParams
}

func GetPetsHTTPRequest(r *http.Request) GetPetsRequest {
	return getPetsHTTPRequest{r}
}

type getPetsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsHTTPRequest) Parse() GetPetsParams {
	return newGetPetsParams(r.Request)
}

type GetPetsParams struct {
}

func newGetPetsParams(r *http.Request) (zero GetPetsParams) {
	var params GetPetsParams

	return params
}

func (r GetPetsParams) HTTP() *http.Request { return nil }

func (r GetPetsParams) Parse() GetPetsParams { return r }

type GetPetsResponse interface {
	writeGetPets(http.ResponseWriter)
}

func NewGetPetsResponse200JSON(body []Pet) GetPetsResponse {
	var out GetPetsResponse200JSON
	out.Body = body
	return out
}

type GetPetsResponse200JSON struct {
	Body []Pet
}

func (r GetPetsResponse200JSON) writeGetPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsResponse200JSON")
}

// ---------------------------------------------
// GetPetsNames -
// ---------------------------------------------

type GetPetsNamesHandlerFunc func(ctx context.Context, r GetPetsNamesRequest) GetPetsNamesResponse

func (f GetPetsNamesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsNamesHTTPRequest(r)).writeGetPetsNames(w)
}

type GetPetsNamesRequest interface {
	HTTP() *http.Request
	Parse() GetPetsNamesParams
}

func GetPetsNamesHTTPRequest(r *http.Request) GetPetsNamesRequest {
	return getPetsNamesHTTPRequest{r}
}

type getPetsNamesHTTPRequest struct {
	Request *http.Request
}

func (r getPetsNamesHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsNamesHTTPRequest) Parse() GetPetsNamesParams {
	return newGetPetsNamesParams(r.Request)
}

type GetPetsNamesParams struct {
}

func newGetPetsNamesParams(r *http.Request) (zero GetPetsNamesParams) {
	var params GetPetsNamesParams

	return params
}

func (r GetPetsNamesParams) HTTP() *http.Request { return nil }

func (r GetPetsNamesParams) Parse() GetPetsNamesParams { return r }

type GetPetsNamesResponse interface {
	writeGetPetsNames(http.ResponseWriter)
}

func NewGetPetsNamesResponse200JSON(body []string) GetPetsNamesResponse {
	var out GetPetsNamesResponse200JSON
	out.Body = body
	return out
}

type GetPetsNamesResponse200JSON struct {
	Body []string
}

func (r GetPetsNamesResponse200JSON) writeGetPetsNames(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsNamesResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsNamesResponse200JSON")
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

func write(w io.Writer, r io.Reader, name string) {
	_, err := io.Copy(w, r)
	if err != nil {
		LogError(fmt.Errorf("write response %q: %w", name, err))
	}
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
