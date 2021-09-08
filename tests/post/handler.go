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

func PostPetsHandler(h PostPetsHandlerer) http.Handler {
	return PostPetsHandlerFunc(h.Handle)
}

func PostPetsHandlerFunc(fn FuncPostPets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newPostPetsParams(r)

		fn(params).writePostPetsResponse(w)
	}
}

type PostPetsHandlerer interface {
	Handle(PostPetsParams) PostPetsResponser
}

func NewPostPetsHandlerer(fn FuncPostPets) PostPetsHandlerer {
	return fn
}

type FuncPostPets func(PostPetsParams) PostPetsResponser

func (f FuncPostPets) Handle(params PostPetsParams) PostPetsResponser { return f(params) }

type PostPetsParams struct {
	Request *http.Request
}

func newPostPetsParams(r *http.Request) (zero PostPetsParams) {
	var params PostPetsParams
	params.Request = r

	return params
}

type PostPetsResponser interface {
	writePostPetsResponse(w http.ResponseWriter)
}

func PostPetsResponse200() PostPetsResponser {
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
