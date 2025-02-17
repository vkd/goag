package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostPets -
// ---------------------------------------------

type PostPetsHandlerFunc func(ctx context.Context, r PostPetsRequest) PostPetsResponse

func (f PostPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostPetsHTTPRequest(r)).writePostPets(w)
}

type PostPetsRequest interface {
	HTTP() *http.Request
	Parse() PostPetsParams
}

func PostPetsHTTPRequest(r *http.Request) PostPetsRequest {
	return postPetsHTTPRequest{r}
}

type postPetsHTTPRequest struct {
	Request *http.Request
}

func (r postPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postPetsHTTPRequest) Parse() PostPetsParams {
	return newPostPetsParams(r.Request)
}

type PostPetsParams struct {
}

func newPostPetsParams(r *http.Request) (zero PostPetsParams) {
	var params PostPetsParams

	return params
}

func (r PostPetsParams) HTTP() *http.Request { return nil }

func (r PostPetsParams) Parse() PostPetsParams { return r }

type PostPetsResponse interface {
	writePostPets(http.ResponseWriter)
}

func NewPostPetsResponse200() PostPetsResponse {
	var out PostPetsResponse200
	return out
}

type PostPetsResponse200 struct{}

func (r PostPetsResponse200) writePostPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostPetsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
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
