package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPet -
// ---------------------------------------------

type GetPetHandlerFunc func(ctx context.Context, r GetPetRequest) GetPetResponse

func (f GetPetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetHTTPRequest(r)).writeGetPet(w)
}

type GetPetRequest interface {
	HTTP() *http.Request
	Parse() GetPetParams
}

func GetPetHTTPRequest(r *http.Request) GetPetRequest {
	return getPetHTTPRequest{r}
}

type getPetHTTPRequest struct {
	Request *http.Request
}

func (r getPetHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetHTTPRequest) Parse() GetPetParams {
	return newGetPetParams(r.Request)
}

type GetPetParams struct {
}

func newGetPetParams(r *http.Request) (zero GetPetParams) {
	var params GetPetParams

	return params
}

func (r GetPetParams) HTTP() *http.Request { return nil }

func (r GetPetParams) Parse() GetPetParams { return r }

type GetPetResponse interface {
	writeGetPet(http.ResponseWriter)
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

func (r GetPetResponse200JSON) writeGetPet(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
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

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
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
