package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostPets -
// ---------------------------------------------

type PostPetsHandlerFunc func(PostPetsParamsParser) PostPetsResponser

func (f PostPetsHandlerFunc) Handle(p PostPetsParamsParser) PostPetsResponser {
	return f(p)
}

func (f PostPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestPostPetsParams{Request: r}).writePostPetsResponse(w)
}

type PostPetsParamsParser interface {
	Parse() (PostPetsParams, error)
}

type requestPostPetsParams struct {
	Request *http.Request
}

func (p requestPostPetsParams) Parse() (PostPetsParams, error) {
	return newPostPetsParams(p.Request)
}

type PostPetsParams struct {
	Request *http.Request

	Body NewPet
}

func newPostPetsParams(r *http.Request) (zero PostPetsParams, _ error) {
	var params PostPetsParams
	params.Request = r

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

type PostPetsResponser interface {
	writePostPetsResponse(w http.ResponseWriter)
}

func PostPetsResponse201() PostPetsResponser {
	var out postPetsResponse201
	return out
}

type postPetsResponse201 struct{}

func (r postPetsResponse201) writePostPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func PostPetsResponseDefault(code int) PostPetsResponser {
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
