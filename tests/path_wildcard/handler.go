package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// GetPetsPetID -
// ---------------------------------------------

type GetPetsPetIDHandlerFunc func(ctx context.Context, r GetPetsPetIDRequest) GetPetsPetIDResponse

func (f GetPetsPetIDHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsPetIDHTTPRequest(r)).writeGetPetsPetID(w)
}

type GetPetsPetIDRequest interface {
	HTTP() *http.Request
	Parse() (GetPetsPetIDParams, error)
}

func GetPetsPetIDHTTPRequest(r *http.Request) GetPetsPetIDRequest {
	return getPetsPetIDHTTPRequest{r}
}

type getPetsPetIDHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsPetIDHTTPRequest) Parse() (GetPetsPetIDParams, error) {
	return newGetPetsPetIDParams(r.Request)
}

type GetPetsPetIDParams struct {
	Path GetPetsPetIDParamsPath
}

type GetPetsPetIDParamsPath struct {
	PetID int32
}

func newGetPetsPetIDParams(r *http.Request) (zero GetPetsPetIDParams, _ error) {
	var params GetPetsPetIDParams

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}'")
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

			vInt64, err := strconv.ParseInt(vPath, 10, 32)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "pet_id", Reason: "parse int32", Err: err}
			}
			params.Path.PetID = int32(vInt64)
		}
	}

	return params, nil
}

func (r GetPetsPetIDParams) HTTP() *http.Request { return nil }

func (r GetPetsPetIDParams) Parse() (GetPetsPetIDParams, error) { return r, nil }

type GetPetsPetIDResponse interface {
	writeGetPetsPetID(http.ResponseWriter)
}

func NewGetPetsPetIDResponse200JSON(body Pet) GetPetsPetIDResponse {
	var out GetPetsPetIDResponse200JSON
	out.Body = body
	return out
}

type GetPetsPetIDResponse200JSON struct {
	Body Pet
}

func (r GetPetsPetIDResponse200JSON) writeGetPetsPetID(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsPetIDResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsPetIDResponse200JSON")
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
