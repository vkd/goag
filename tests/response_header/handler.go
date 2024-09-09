package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func NewGetPetsResponse200(xNext Maybe[string], xNextTwo []int) GetPetsResponse {
	var out GetPetsResponse200
	out.Headers.XNext = xNext
	out.Headers.XNextTwo = xNextTwo
	return out
}

type GetPetsResponse200 struct {
	Headers struct {
		XNext    Maybe[string]
		XNextTwo []int
	}
}

func (r GetPetsResponse200) writeGetPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsResponse200) Write(w http.ResponseWriter) {
	if r.Headers.XNext.IsSet {
		hs := []string{r.Headers.XNext.Value}
		for _, h := range hs {
			w.Header().Add("x-next", h)
		}
	}
	{
		qv := make([]string, 0, len(r.Headers.XNextTwo))
		for _, v := range r.Headers.XNextTwo {
			qv = append(qv, strconv.FormatInt(int64(v), 10))
		}
		hs := qv
		for _, h := range hs {
			w.Header().Add("x-next-two", h)
		}
	}
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
