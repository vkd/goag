package test

import (
	"context"
	"encoding/json"
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
	Parse() (PostPetsParams, error)
}

func PostPetsHTTPRequest(r *http.Request) PostPetsRequest {
	return postPetsHTTPRequest{r}
}

type postPetsHTTPRequest struct {
	Request *http.Request
}

func (r postPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postPetsHTTPRequest) Parse() (PostPetsParams, error) {
	return newPostPetsParams(r.Request)
}

type PostPetsParams struct {
	Body NewPet
}

func newPostPetsParams(r *http.Request) (zero PostPetsParams, _ error) {
	var params PostPetsParams

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r PostPetsParams) HTTP() *http.Request { return nil }

func (r PostPetsParams) Parse() (PostPetsParams, error) { return r, nil }

type PostPetsResponse interface {
	postPets()
	Write(w http.ResponseWriter)
}

func NewPostPetsResponse201() PostPetsResponse {
	var out PostPetsResponse201
	return out
}

type PostPetsResponse201 struct{}

func (r PostPetsResponse201) postPets() {}

func (r PostPetsResponse201) Write(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func NewPostPetsResponseDefault(code int) PostPetsResponse {
	var out PostPetsResponseDefault
	out.Code = code
	return out
}

type PostPetsResponseDefault struct {
	Code int
}

func (r PostPetsResponseDefault) postPets() {}

func (r PostPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
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
