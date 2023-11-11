package test

import (
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostPets -
// ---------------------------------------------

type PostPetsHandlerFunc func(r PostPetsRequestParser) PostPetsResponse

func (f PostPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(PostPetsHTTPRequest(r)).Write(w)
}

type PostPetsRequestParser interface {
	Parse() PostPetsRequest
}

func PostPetsHTTPRequest(r *http.Request) PostPetsRequestParser {
	return postPetsHTTPRequest{r}
}

type postPetsHTTPRequest struct {
	Request *http.Request
}

func (r postPetsHTTPRequest) Parse() PostPetsRequest {
	return newPostPetsParams(r.Request)
}

type PostPetsRequest struct {
	HTTPRequest *http.Request
}

func newPostPetsParams(r *http.Request) (zero PostPetsRequest) {
	var params PostPetsRequest
	params.HTTPRequest = r

	return params
}

func (r PostPetsRequest) Parse() PostPetsRequest { return r }

type PostPetsResponse interface {
	postPets()
	Write(w http.ResponseWriter)
}

func NewPostPetsResponse200() PostPetsResponse {
	var out PostPetsResponse200
	return out
}

type PostPetsResponse200 struct{}

func (r PostPetsResponse200) postPets() {}

func (r PostPetsResponse200) Write(w http.ResponseWriter) {
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
