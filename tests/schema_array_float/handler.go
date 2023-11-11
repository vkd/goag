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

type GetPetsIDsHandlerFunc func(r GetPetsIDsRequestParser) GetPetsIDsResponse

func (f GetPetsIDsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetsIDsHTTPRequest(r)).Write(w)
}

type GetPetsIDsRequestParser interface {
	Parse() GetPetsIDsRequest
}

func GetPetsIDsHTTPRequest(r *http.Request) GetPetsIDsRequestParser {
	return getPetsIDsHTTPRequest{r}
}

type getPetsIDsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsIDsHTTPRequest) Parse() GetPetsIDsRequest {
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

func (r GetPetsIDsRequest) Parse() GetPetsIDsRequest { return r }

type GetPetsIDsResponse interface {
	getPetsIDs()
	Write(w http.ResponseWriter)
}

func NewGetPetsIDsResponse200JSON(body []float64) GetPetsIDsResponse {
	var out GetPetsIDsResponse200JSON
	out.Body = body
	return out
}

type GetPetsIDsResponse200JSON struct {
	Body []float64
}

func (r GetPetsIDsResponse200JSON) getPetsIDs() {}

func (r GetPetsIDsResponse200JSON) Write(w http.ResponseWriter) {
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
