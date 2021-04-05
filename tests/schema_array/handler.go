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
	return GetPetsHandlerFunc(h.Handler)
}

func GetPetsHandlerFunc(fn FuncGetPets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetPetsParams(r)

		fn(params).writeGetPetsResponse(w)
	}
}

type GetPetsHandlerer interface {
	Handler(GetPetsParams) GetPetsResponser
}

func NewGetPetsHandlerer(fn FuncGetPets) GetPetsHandlerer {
	return fn
}

type FuncGetPets func(GetPetsParams) GetPetsResponser

func (f FuncGetPets) Handler(params GetPetsParams) GetPetsResponser { return f(params) }

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

func GetPetsResponse200JSON(body []Pet) GetPetsResponser {
	var out getPetsResponse200JSON
	out.Body = body
	return out
}

type getPetsResponse200JSON struct {
	Body []Pet
}

func (r getPetsResponse200JSON) writeGetPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsResponse200JSON")
}

// ---------------------------------------------
// GetPetsNames -
// ---------------------------------------------

func GetPetsNamesHandler(h GetPetsNamesHandlerer) http.Handler {
	return GetPetsNamesHandlerFunc(h.Handler)
}

func GetPetsNamesHandlerFunc(fn FuncGetPetsNames) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetPetsNamesParams(r)

		fn(params).writeGetPetsNamesResponse(w)
	}
}

type GetPetsNamesHandlerer interface {
	Handler(GetPetsNamesParams) GetPetsNamesResponser
}

func NewGetPetsNamesHandlerer(fn FuncGetPetsNames) GetPetsNamesHandlerer {
	return fn
}

type FuncGetPetsNames func(GetPetsNamesParams) GetPetsNamesResponser

func (f FuncGetPetsNames) Handler(params GetPetsNamesParams) GetPetsNamesResponser { return f(params) }

type GetPetsNamesParams struct {
	Request *http.Request
}

func newGetPetsNamesParams(r *http.Request) (zero GetPetsNamesParams) {
	var params GetPetsNamesParams
	params.Request = r

	return params
}

type GetPetsNamesResponser interface {
	writeGetPetsNamesResponse(w http.ResponseWriter)
}

func GetPetsNamesResponse200JSON(body []string) GetPetsNamesResponser {
	var out getPetsNamesResponse200JSON
	out.Body = body
	return out
}

type getPetsNamesResponse200JSON struct {
	Body []string
}

func (r getPetsNamesResponse200JSON) writeGetPetsNamesResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsNamesResponse200JSON")
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
