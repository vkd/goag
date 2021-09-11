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

type GetPetHandlerFunc func(GetPetParamsParser) GetPetResponser

func (f GetPetHandlerFunc) Handle(p GetPetParamsParser) GetPetResponser {
	return f(p)
}

func (f GetPetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetParams{Request: r}).writeGetPetResponse(w)
}

type GetPetParamsParser interface {
	Parse() GetPetParams
}

type requestGetPetParams struct {
	Request *http.Request
}

func (p requestGetPetParams) Parse() GetPetParams {
	return newGetPetParams(p.Request)
}

type GetPetParams struct {
	HTTPRequest *http.Request
}

func newGetPetParams(r *http.Request) (zero GetPetParams) {
	var params GetPetParams
	params.HTTPRequest = r

	return params
}

type GetPetResponser interface {
	writeGetPetResponse(w http.ResponseWriter)
}

func GetPetResponse200JSON(body GetPetResponse200JSONBody) GetPetResponser {
	var out getPetResponse200JSON
	out.Body = body
	return out
}

type GetPetResponse200JSONBody struct {
	Length               int             `json:"length"`
	AdditionalProperties map[string]Pets `json:"-"`
}

var _ json.Marshaler = (*GetPetResponse200JSONBody)(nil)

func (b GetPetResponse200JSONBody) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range b.AdditionalProperties {
		m[k] = v
	}
	m["length"] = b.Length
	return json.Marshal(m)

}

type getPetResponse200JSON struct {
	Body GetPetResponse200JSONBody
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
