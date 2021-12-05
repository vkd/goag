package test

import (
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostPets -
// ---------------------------------------------

type PostPetsHandlerFunc func(r PostPetsRequester) PostPetsResponder

func (f PostPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestPostPetsParams{Request: r}).writePostPetsResponse(w)
}

type PostPetsRequester interface {
	Parse() PostPetsRequest
}

type requestPostPetsParams struct {
	Request *http.Request
}

func (r requestPostPetsParams) Parse() PostPetsRequest {
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

type PostPetsResponder interface {
	writePostPetsResponse(w http.ResponseWriter)
}

func PostPetsResponse200() PostPetsResponder {
	var out postPetsResponse200
	return out
}

type postPetsResponse200 struct{}

func (r postPetsResponse200) writePostPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
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
