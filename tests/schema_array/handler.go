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

type GetPetsHandlerFunc func(GetPetsParamsParser) GetPetsResponser

func (f GetPetsHandlerFunc) Handle(p GetPetsParamsParser) GetPetsResponser {
	return f(p)
}

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetsParams{Request: r}).writeGetPetsResponse(w)
}

type GetPetsParamsParser interface {
	Parse() GetPetsParams
}

type requestGetPetsParams struct {
	Request *http.Request
}

func (p requestGetPetsParams) Parse() GetPetsParams {
	return newGetPetsParams(p.Request)
}

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

type GetPetsNamesHandlerFunc func(GetPetsNamesParamsParser) GetPetsNamesResponser

func (f GetPetsNamesHandlerFunc) Handle(p GetPetsNamesParamsParser) GetPetsNamesResponser {
	return f(p)
}

func (f GetPetsNamesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetsNamesParams{Request: r}).writeGetPetsNamesResponse(w)
}

type GetPetsNamesParamsParser interface {
	Parse() GetPetsNamesParams
}

type requestGetPetsNamesParams struct {
	Request *http.Request
}

func (p requestGetPetsNamesParams) Parse() GetPetsNamesParams {
	return newGetPetsNamesParams(p.Request)
}

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
