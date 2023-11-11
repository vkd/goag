package test

import (
	"fmt"
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

func NewGetPetsResponse200(xNext string) GetPetsResponse {
	var out GetPetsResponse200
	out.Headers.XNext = xNext
	return out
}

type GetPetsResponse200 struct {
	Headers struct {
		XNext string
	}
}

func (r GetPetsResponse200) getPets() {}

func (r GetPetsResponse200) Write(w http.ResponseWriter) {
	w.Header().Set("x-next", r.Headers.XNext)
	w.WriteHeader(200)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
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
