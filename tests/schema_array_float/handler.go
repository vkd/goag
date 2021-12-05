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

type GetPetsIDsHandlerFunc func(r GetPetsIDsRequester) GetPetsIDsResponder

func (f GetPetsIDsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetsIDsParams{Request: r}).writeGetPetsIDsResponse(w)
}

type GetPetsIDsRequester interface {
	Parse() GetPetsIDsRequest
}

type requestGetPetsIDsParams struct {
	Request *http.Request
}

func (r requestGetPetsIDsParams) Parse() GetPetsIDsRequest {
	return newGetPetsIDsParams(r.Request)
}

type GetPetsIDsRequest struct {
	HTTPRequest *http.Request
}

func newGetPetsIDsParams(r *http.Request) (zero GetPetsIDsRequest) {
	var params GetPetsIDsRequest
	params.HTTPRequest = r

	return params
}

type GetPetsIDsResponder interface {
	writeGetPetsIDsResponse(w http.ResponseWriter)
}

func GetPetsIDsResponse200JSON(body []float64) GetPetsIDsResponder {
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
