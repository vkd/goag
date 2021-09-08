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

func GetPetHandler(h GetPetHandlerer) http.Handler {
	return GetPetHandlerFunc(h.Handle)
}

func GetPetHandlerFunc(fn FuncGetPet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetPetParams(r)

		fn(params).writeGetPetResponse(w)
	}
}

type GetPetHandlerer interface {
	Handle(GetPetParams) GetPetResponser
}

func NewGetPetHandlerer(fn FuncGetPet) GetPetHandlerer {
	return fn
}

type FuncGetPet func(GetPetParams) GetPetResponser

func (f FuncGetPet) Handle(params GetPetParams) GetPetResponser { return f(params) }

type GetPetParams struct {
	Request *http.Request
}

func newGetPetParams(r *http.Request) (zero GetPetParams) {
	var params GetPetParams
	params.Request = r

	return params
}

type GetPetResponser interface {
	writeGetPetResponse(w http.ResponseWriter)
}

func GetPetResponse200JSON(body Pet) GetPetResponser {
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
