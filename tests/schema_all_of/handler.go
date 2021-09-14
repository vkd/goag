package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPet -
// ---------------------------------------------

type GetPetHandlerFunc func(r GetPetRequester) GetPetResponder

func (f GetPetHandlerFunc) Handle(r GetPetRequester) GetPetResponder {
	return f(r)
}

func (f GetPetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetParams{Request: r}).writeGetPetResponse(w)
}

type GetPetRequester interface {
	Parse() GetPetRequest
}

type requestGetPetParams struct {
	Request *http.Request
}

func (r requestGetPetParams) Parse() GetPetRequest {
	return newGetPetParams(r.Request)
}

type GetPetRequest struct {
	HTTPRequest *http.Request
}

func newGetPetParams(r *http.Request) (zero GetPetRequest) {
	var params GetPetRequest
	params.HTTPRequest = r

	return params
}

type GetPetResponder interface {
	writeGetPetResponse(w http.ResponseWriter)
}

func GetPetResponse200JSON(body Pet) GetPetResponder {
	var out getPetResponse200JSON
	out.Body = body
	return out
}

type getPetResponse200JSON struct {
	Body Pet
}

func (r getPetResponse200JSON) writeGetPetResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetResponse200JSON")
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
