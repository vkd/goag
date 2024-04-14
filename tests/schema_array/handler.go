package test

import (
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

func NewGetPetsResponse200JSON(body []Pet) GetPetsResponse {
	var out GetPetsResponse200JSON
	out.Body = body
	return out
}

type GetPetsResponse200JSON struct {
	Body []Pet
}

func (r GetPetsResponse200JSON) getPets() {}

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
	f(r.Context(), GetPetsNamesHTTPRequest(r)).Write(w)
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
	getPetsNames()
	Write(w http.ResponseWriter)
}

func NewGetPetsNamesResponse200JSON(body []string) GetPetsNamesResponse {
	var out GetPetsNamesResponse200JSON
	out.Body = body
	return out
}

type GetPetsNamesResponse200JSON struct {
	Body []string
}

func (r GetPetsNamesResponse200JSON) getPetsNames() {}

func (r GetPetsNamesResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsNamesResponse200JSON")
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
