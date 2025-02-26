package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ---------------------------------------------
// GetPets -
// GET /pets
// ---------------------------------------------

type GetPetsHandlerFunc func(ctx context.Context, r GetPetsRequest) GetPetsResponse

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsHTTPRequest(r)).writeGetPets(w)
}

func (GetPetsHandlerFunc) Path() string { return "/pets" }

func (GetPetsHandlerFunc) Method() string { return http.MethodGet }

type GetPetsRequest interface {
	HTTP() *http.Request
	Parse() (GetPetsParams, error)
}

func GetPetsHTTPRequest(r *http.Request) GetPetsRequest {
	return getPetsHTTPRequest{r}
}

type getPetsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsHTTPRequest) Parse() (GetPetsParams, error) {
	return newGetPetsParams(r.Request)
}

type GetPetsParams struct {
	Query GetPetsParamsQuery
}

type GetPetsParamsQuery struct {
	Tag Maybe[[]string]

	Page Maybe[[]int64]
}

func newGetPetsParams(r *http.Request) (zero GetPetsParams, _ error) {
	var params GetPetsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["tag"]
			if ok && len(q) > 0 {
				vOpt := q
				params.Query.Tag.Set(vOpt)
			}
		}
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				vOpt := make([]int64, len(q))
				for i := range q {
					var err error
					vOpt[i], err = strconv.ParseInt(q[i], 10, 64)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int64", Err: err}
					}
				}
				params.Query.Page.Set(vOpt)
			}
		}
	}

	return params, nil
}

func (r GetPetsParams) HTTP() *http.Request { return nil }

func (r GetPetsParams) Parse() (GetPetsParams, error) { return r, nil }

type GetPetsResponse interface {
	writeGetPets(http.ResponseWriter)
}

func NewGetPetsResponse200() GetPetsResponse {
	var out GetPetsResponse200
	return out
}

type GetPetsResponse200 struct{}

func (r GetPetsResponse200) writeGetPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewGetPetsResponseDefault(code int) GetPetsResponse {
	var out GetPetsResponseDefault
	out.Code = code
	return out
}

type GetPetsResponseDefault struct {
	Code int
}

func (r GetPetsResponseDefault) writeGetPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
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
