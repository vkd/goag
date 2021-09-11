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

type GetPetsHandlerFunc func(GetPetsParamsParser) GetPetsResponser

func (f GetPetsHandlerFunc) Handle(p GetPetsParamsParser) GetPetsResponser {
	return f(p)
}

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetsParams{Request: r}).writeGetPetsResponse(w)
}

type GetPetsParamsParser interface {
	Parse() GetPetsParams
}

type requestGetPetsParams struct {
	Request *http.Request
}

func (p requestGetPetsParams) Parse() GetPetsParams {
	return newGetPetsParams(p.Request)
}

type GetPetsParams struct {
	HTTPRequest *http.Request
}

func newGetPetsParams(r *http.Request) (zero GetPetsParams) {
	var params GetPetsParams
	params.HTTPRequest = r

	return params
}

type GetPetsResponser interface {
	writeGetPetsResponse(w http.ResponseWriter)
}

func GetPetsResponse200(xNext string) GetPetsResponser {
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
