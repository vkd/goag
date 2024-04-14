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
	f(r.Context(), GetPetsHTTPRequest(r)).Write(w)
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
	getPets()
	Write(w http.ResponseWriter)
}

func NewGetPetsResponse200(xNext string, xNextTwo []int) GetPetsResponse {
	var out GetPetsResponse200
	out.Headers.XNext = xNext
	out.Headers.XNextTwo = xNextTwo
	return out
}

type GetPetsResponse200 struct {
	Headers struct {
		XNext    string
		XNextTwo []int
	}
}

func (r GetPetsResponse200) getPets() {}

func (r GetPetsResponse200) Write(w http.ResponseWriter) {
	w.Header().Set("x-next", r.Headers.XNext)
	for _, h := range r.Headers.XNextTwo {
		w.Header().Add("x-next-two", strconv.FormatInt(int64(h), 10))
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

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		Value: v,
		IsSet: true,
	}
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
