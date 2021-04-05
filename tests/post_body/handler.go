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
	return PostPetsHandlerFunc(h.Handler, h.InvalidResponce)
}

func PostPetsHandlerFunc(fn FuncPostPets, invalidFn FuncPostPetsInvalidResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := newPostPetsParams(r)
		if err != nil {
			invalidFn(err).writePostPetsResponse(w)
			return
		}

		fn(params).writePostPetsResponse(w)
	}
}

type PostPetsHandlerer interface {
	Handler(PostPetsParams) PostPetsResponser
	InvalidResponce(error) PostPetsResponser
}

func NewPostPetsHandlerer(fn FuncPostPets, invalidFn FuncPostPetsInvalidResponse) PostPetsHandlerer {
	return privatePostPetsHandlerer{
		FuncPostPets:                fn,
		FuncPostPetsInvalidResponse: invalidFn,
	}
}

type privatePostPetsHandlerer struct {
	FuncPostPets
	FuncPostPetsInvalidResponse
}

type FuncPostPets func(PostPetsParams) PostPetsResponser

func (f FuncPostPets) Handler(params PostPetsParams) PostPetsResponser { return f(params) }

type FuncPostPetsInvalidResponse func(error) PostPetsResponser

func (f FuncPostPetsInvalidResponse) InvalidResponce(err error) PostPetsResponser { return f(err) }

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
