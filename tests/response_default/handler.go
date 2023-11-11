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

type GetPetsHandlerFunc func(r GetPetsRequestParser) GetPetsResponse

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetsHTTPRequest(r)).Write(w)
}

type GetPetsRequestParser interface {
	Parse() GetPetsRequest
}

func GetPetsHTTPRequest(r *http.Request) GetPetsRequestParser {
	return getPetsHTTPRequest{r}
}

type getPetsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsHTTPRequest) Parse() GetPetsRequest {
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

func (r GetPetsRequest) Parse() GetPetsRequest { return r }

type GetPetsResponse interface {
	getPets()
	Write(w http.ResponseWriter)
}

func NewGetPetsResponse200() GetPetsResponse {
	var out GetPetsResponse200
	return out
}

type GetPetsResponse200 struct{}

func (r GetPetsResponse200) getPets() {}

func (r GetPetsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewGetPetsResponseDefaultJSON(code int, body Error) GetPetsResponse {
	var out GetPetsResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

type GetPetsResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r GetPetsResponseDefaultJSON) getPets() {}

func (r GetPetsResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "GetPetsResponseDefaultJSON")
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
