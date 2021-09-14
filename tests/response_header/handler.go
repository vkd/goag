package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPets -
// ---------------------------------------------

type GetPetsHandlerFunc func(r GetPetsRequester) GetPetsResponder

func (f GetPetsHandlerFunc) Handle(r GetPetsRequester) GetPetsResponder {
	return f(r)
}

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetsParams{Request: r}).writeGetPetsResponse(w)
}

type GetPetsRequester interface {
	Parse() GetPetsRequest
}

type requestGetPetsParams struct {
	Request *http.Request
}

func (r requestGetPetsParams) Parse() GetPetsRequest {
	return newGetPetsParams(r.Request)
}

type GetPetsRequest struct {
	HTTPRequest *http.Request
}

func newGetPetsParams(r *http.Request) (zero GetPetsRequest) {
	var params GetPetsRequest
	params.HTTPRequest = r

	return params
}

type GetPetsResponder interface {
	writeGetPetsResponse(w http.ResponseWriter)
}

func GetPetsResponse200(xNext string) GetPetsResponder {
	var out getPetsResponse200
	out.Headers.XNext = xNext
	return out
}

type getPetsResponse200 struct {
	Headers struct {
		XNext string
	}
}

func (r getPetsResponse200) writeGetPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Header().Set("x-next", r.Headers.XNext)
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

type ErrParseQueryParam struct {
	Name string
	Err  error
}

func (e ErrParseQueryParam) Error() string {
	return fmt.Sprintf("query parameter '%s': %e", e.Name, e.Err)
}

type ErrParsePathParam struct {
	Name string
	Err  error
}

func (e ErrParsePathParam) Error() string {
	return fmt.Sprintf("path parameter '%s': %e", e.Name, e.Err)
}
