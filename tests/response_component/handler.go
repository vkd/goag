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

// ---------------------------------------------
// GetV2Pet -
// GET /v2/pet
// ---------------------------------------------

type GetV2PetHandlerFunc func(ctx context.Context, r GetV2PetRequest) GetV2PetResponse

func (f GetV2PetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetV2PetHTTPRequest(r)).writeGetV2Pet(w)
}

func (GetV2PetHandlerFunc) Path() string { return "/v2/pet" }

func (GetV2PetHandlerFunc) Method() string { return http.MethodGet }

type GetV2PetRequest interface {
	HTTP() *http.Request
	Parse() GetV2PetParams
}

func GetV2PetHTTPRequest(r *http.Request) GetV2PetRequest {
	return getV2PetHTTPRequest{r}
}

type getV2PetHTTPRequest struct {
	Request *http.Request
}

func (r getV2PetHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getV2PetHTTPRequest) Parse() GetV2PetParams {
	return newGetV2PetParams(r.Request)
}

type GetV2PetParams struct {
}

func newGetV2PetParams(r *http.Request) (zero GetV2PetParams) {
	var params GetV2PetParams

	return params
}

func (r GetV2PetParams) HTTP() *http.Request { return nil }

func (r GetV2PetParams) Parse() GetV2PetParams { return r }

type GetV2PetResponse interface {
	writeGetV2Pet(http.ResponseWriter)
}

// ---------------------------------------------
// GetV3Pet -
// GET /v3/pet
// ---------------------------------------------

type GetV3PetHandlerFunc func(ctx context.Context, r GetV3PetRequest) GetV3PetResponse

func (f GetV3PetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetV3PetHTTPRequest(r)).writeGetV3Pet(w)
}

func (GetV3PetHandlerFunc) Path() string { return "/v3/pet" }

func (GetV3PetHandlerFunc) Method() string { return http.MethodGet }

type GetV3PetRequest interface {
	HTTP() *http.Request
	Parse() GetV3PetParams
}

func GetV3PetHTTPRequest(r *http.Request) GetV3PetRequest {
	return getV3PetHTTPRequest{r}
}

type getV3PetHTTPRequest struct {
	Request *http.Request
}

func (r getV3PetHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getV3PetHTTPRequest) Parse() GetV3PetParams {
	return newGetV3PetParams(r.Request)
}

type GetV3PetParams struct {
}

func newGetV3PetParams(r *http.Request) (zero GetV3PetParams) {
	var params GetV3PetParams

	return params
}

func (r GetV3PetParams) HTTP() *http.Request { return nil }

func (r GetV3PetParams) Parse() GetV3PetParams { return r }

type GetV3PetResponse interface {
	writeGetV3Pet(http.ResponseWriter)
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
