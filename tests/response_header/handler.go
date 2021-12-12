package test

import (
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPets -
// ---------------------------------------------

type GetPetsHandlerFunc func(r GetPetsRequester) GetPetsResponder

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetsParams{Request: r}).writeGetPetsResponse(w)
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
