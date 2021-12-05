package test

import (
	"encoding/json"
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
	Parse() (PostPetsRequest, error)
}

type requestPostPetsParams struct {
	Request *http.Request
}

func (r requestPostPetsParams) Parse() (PostPetsRequest, error) {
	return newPostPetsParams(r.Request)
}

type PostPetsRequest struct {
	HTTPRequest *http.Request

	Body NewPet
}

func newPostPetsParams(r *http.Request) (zero PostPetsRequest, _ error) {
	var params PostPetsRequest
	params.HTTPRequest = r

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

type PostPetsResponder interface {
	writePostPetsResponse(w http.ResponseWriter)
}

func PostPetsResponse201() PostPetsResponder {
	var out postPetsResponse201
	return out
}

type postPetsResponse201 struct{}

func (r postPetsResponse201) writePostPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func PostPetsResponseDefault(code int) PostPetsResponder {
	var out postPetsResponseDefault
	out.Code = code
	return out
}

type postPetsResponseDefault struct {
	Code int
}

func (r postPetsResponseDefault) writePostPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
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
