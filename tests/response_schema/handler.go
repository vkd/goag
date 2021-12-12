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

func (f GetPetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetParams{Request: r}).writeGetPetResponse(w)
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
