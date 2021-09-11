package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPetsIDs -
// ---------------------------------------------

type GetPetsIDsHandlerFunc func(GetPetsIDsParamsParser) GetPetsIDsResponser

func (f GetPetsIDsHandlerFunc) Handle(p GetPetsIDsParamsParser) GetPetsIDsResponser {
	return f(p)
}

func (f GetPetsIDsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetsIDsParams{Request: r}).writeGetPetsIDsResponse(w)
}

type GetPetsIDsParamsParser interface {
	Parse() GetPetsIDsParams
}

type requestGetPetsIDsParams struct {
	Request *http.Request
}

func (p requestGetPetsIDsParams) Parse() GetPetsIDsParams {
	return newGetPetsIDsParams(p.Request)
}

type GetPetsIDsParams struct {
	HTTPRequest *http.Request
}

func newGetPetsIDsParams(r *http.Request) (zero GetPetsIDsParams) {
	var params GetPetsIDsParams
	params.HTTPRequest = r

	return params
}

type GetPetsIDsResponser interface {
	writeGetPetsIDsResponse(w http.ResponseWriter)
}

func GetPetsIDsResponse200JSON(body []float64) GetPetsIDsResponser {
	var out getPetsIDsResponse200JSON
	out.Body = body
	return out
}

type getPetsIDsResponse200JSON struct {
	Body []float64
}

func (r getPetsIDsResponse200JSON) writeGetPetsIDsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsIDsResponse200JSON")
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
