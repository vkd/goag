package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// GetPetsPetIDNames -
// GET /pets/{pet_id}/names
// ---------------------------------------------

type GetPetsPetIDNamesHandlerFunc func(ctx context.Context, r GetPetsPetIDNamesRequest) GetPetsPetIDNamesResponse

func (f GetPetsPetIDNamesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsPetIDNamesHTTPRequest(r)).writeGetPetsPetIDNames(w)
}

func (GetPetsPetIDNamesHandlerFunc) Path() string { return "/pets/{pet_id}/names" }

func (GetPetsPetIDNamesHandlerFunc) Method() string { return http.MethodGet }

type GetPetsPetIDNamesRequest interface {
	HTTP() *http.Request
	Parse() (GetPetsPetIDNamesParams, error)
}

func GetPetsPetIDNamesHTTPRequest(r *http.Request) GetPetsPetIDNamesRequest {
	return getPetsPetIDNamesHTTPRequest{r}
}

type getPetsPetIDNamesHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDNamesHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsPetIDNamesHTTPRequest) Parse() (GetPetsPetIDNamesParams, error) {
	return newGetPetsPetIDNamesParams(r.Request)
}

type GetPetsPetIDNamesParams struct {
	Path GetPetsPetIDNamesParamsPath
}

type GetPetsPetIDNamesParamsPath struct {
	PetID string
}

func newGetPetsPetIDNamesParams(r *http.Request) (zero GetPetsPetIDNamesParams, _ error) {
	var params GetPetsPetIDNamesParams

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/names'")
		}
		p = p[6:] // "/pets/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "pet_id", Reason: "required"}
			}

			params.Path.PetID = vPath
		}

		if !strings.HasPrefix(p, "/names") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/names'")
		}
		p = p[6:] // "/names"
	}

	return params, nil
}

func (r GetPetsPetIDNamesParams) HTTP() *http.Request { return nil }

func (r GetPetsPetIDNamesParams) Parse() (GetPetsPetIDNamesParams, error) { return r, nil }

type GetPetsPetIDNamesResponse interface {
	writeGetPetsPetIDNames(http.ResponseWriter)
}

func NewGetPetsPetIDNamesResponse200() GetPetsPetIDNamesResponse {
	var out GetPetsPetIDNamesResponse200
	return out
}

type GetPetsPetIDNamesResponse200 struct{}

func (r GetPetsPetIDNamesResponse200) writeGetPetsPetIDNames(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsPetIDNamesResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

// ---------------------------------------------
// GetPetsPetIDShops -
// GET /pets/{pet_id}/shops
// ---------------------------------------------

type GetPetsPetIDShopsHandlerFunc func(ctx context.Context, r GetPetsPetIDShopsRequest) GetPetsPetIDShopsResponse

func (f GetPetsPetIDShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsPetIDShopsHTTPRequest(r)).writeGetPetsPetIDShops(w)
}

func (GetPetsPetIDShopsHandlerFunc) Path() string { return "/pets/{pet_id}/shops" }

func (GetPetsPetIDShopsHandlerFunc) Method() string { return http.MethodGet }

type GetPetsPetIDShopsRequest interface {
	HTTP() *http.Request
	Parse() (GetPetsPetIDShopsParams, error)
}

func GetPetsPetIDShopsHTTPRequest(r *http.Request) GetPetsPetIDShopsRequest {
	return getPetsPetIDShopsHTTPRequest{r}
}

type getPetsPetIDShopsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDShopsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsPetIDShopsHTTPRequest) Parse() (GetPetsPetIDShopsParams, error) {
	return newGetPetsPetIDShopsParams(r.Request)
}

type GetPetsPetIDShopsParams struct {
	Path GetPetsPetIDShopsParamsPath
}

type GetPetsPetIDShopsParamsPath struct {
	PetID string
}

func newGetPetsPetIDShopsParams(r *http.Request) (zero GetPetsPetIDShopsParams, _ error) {
	var params GetPetsPetIDShopsParams

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/shops'")
		}
		p = p[6:] // "/pets/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "pet_id", Reason: "required"}
			}

			params.Path.PetID = vPath
		}

		if !strings.HasPrefix(p, "/shops") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/shops'")
		}
		p = p[6:] // "/shops"
	}

	return params, nil
}

func (r GetPetsPetIDShopsParams) HTTP() *http.Request { return nil }

func (r GetPetsPetIDShopsParams) Parse() (GetPetsPetIDShopsParams, error) { return r, nil }

type GetPetsPetIDShopsResponse interface {
	writeGetPetsPetIDShops(http.ResponseWriter)
}

func NewGetPetsPetIDShopsResponse200() GetPetsPetIDShopsResponse {
	var out GetPetsPetIDShopsResponse200
	return out
}

type GetPetsPetIDShopsResponse200 struct{}

func (r GetPetsPetIDShopsResponse200) writeGetPetsPetIDShops(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsPetIDShopsResponse200) Write(w http.ResponseWriter) {
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
