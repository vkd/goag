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

type GetPetHandlerFunc func(r GetPetRequestParser) GetPetResponse

func (f GetPetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetHTTPRequest(r)).Write(w)
}

type GetPetRequestParser interface {
	Parse() GetPetRequest
}

func GetPetHTTPRequest(r *http.Request) GetPetRequestParser {
	return getPetHTTPRequest{r}
}

type getPetHTTPRequest struct {
	Request *http.Request
}

func (r getPetHTTPRequest) Parse() GetPetRequest {
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

func (r GetPetRequest) Parse() GetPetRequest { return r }

type GetPetResponse interface {
	getPet()
	Write(w http.ResponseWriter)
}

func NewGetPetResponse200JSON(body GetPetResponse200JSONBody) GetPetResponse {
	var out GetPetResponse200JSON
	out.Body = body
	return out
}

type GetPetResponse200JSONBody struct {
	Groups map[string]Pets `json:"groups"`
}

type GetPetResponse200JSON struct {
	Body GetPetResponse200JSONBody
}

func (r GetPetResponse200JSON) getPet() {}

func (r GetPetResponse200JSON) Write(w http.ResponseWriter) {
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
