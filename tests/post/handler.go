package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostPets -
// ---------------------------------------------

type PostPetsHandlerFunc func(ctx context.Context, r PostPetsRequest) PostPetsResponse

func (f PostPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostPetsHTTPRequest(r)).Write(w)
}

type PostPetsRequest interface {
	HTTP() *http.Request
	Parse() PostPetsParams
}

func PostPetsHTTPRequest(r *http.Request) PostPetsRequest {
	return postPetsHTTPRequest{r}
}

type postPetsHTTPRequest struct {
	Request *http.Request
}

func (r postPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postPetsHTTPRequest) Parse() PostPetsParams {
	return newPostPetsParams(r.Request)
}

type PostPetsParams struct {
}

func newPostPetsParams(r *http.Request) (zero PostPetsParams) {
	var params PostPetsParams

	return params
}

func (r PostPetsParams) HTTP() *http.Request { return nil }

func (r PostPetsParams) Parse() PostPetsParams { return r }

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
