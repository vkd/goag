package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPets -
// ---------------------------------------------

func GetPetsHandler(h GetPetsHandlerer) http.Handler {
	return GetPetsHandlerFunc(h.Handle)
}

func GetPetsHandlerFunc(fn FuncGetPets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetPetsParams(r)

		fn(params).writeGetPetsResponse(w)
	}
}

type GetPetsHandlerer interface {
	Handle(GetPetsParams) GetPetsResponser
}

func NewGetPetsHandlerer(fn FuncGetPets) GetPetsHandlerer {
	return fn
}

type FuncGetPets func(GetPetsParams) GetPetsResponser

func (f FuncGetPets) Handle(params GetPetsParams) GetPetsResponser { return f(params) }

type GetPetsParams struct {
	Request *http.Request
}

func newGetPetsParams(r *http.Request) (zero GetPetsParams) {
	var params GetPetsParams
	params.Request = r

	return params
}

type GetPetsResponser interface {
	writeGetPetsResponse(w http.ResponseWriter)
}

func GetPetsResponse200() GetPetsResponser {
	var out getPetsResponse200
	return out
}

type getPetsResponse200 struct{}

func (r getPetsResponse200) writeGetPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func GetPetsResponseDefaultJSON(code int, body Error) GetPetsResponser {
	var out getPetsResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

type getPetsResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r getPetsResponseDefaultJSON) writeGetPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "GetPetsResponseDefaultJSON")
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
