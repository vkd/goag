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

// ---------------------------------------------
// GetV2Pet -
// ---------------------------------------------

type GetV2PetHandlerFunc func(ctx context.Context, r GetV2PetRequest) GetV2PetResponse

func (f GetV2PetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetV2PetHTTPRequest(r)).writeGetV2Pet(w)
}

type GetV2PetRequest interface {
	HTTP() *http.Request
	Parse() GetV2PetParams
}

func GetV2PetHTTPRequest(r *http.Request) GetV2PetRequest {
	return getV2PetHTTPRequest{r}
}

type getV2PetHTTPRequest struct {
	Request *http.Request
}

func (r getV2PetHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getV2PetHTTPRequest) Parse() GetV2PetParams {
	return newGetV2PetParams(r.Request)
}

type GetV2PetParams struct {
}

func newGetV2PetParams(r *http.Request) (zero GetV2PetParams) {
	var params GetV2PetParams

	return params
}

func (r GetV2PetParams) HTTP() *http.Request { return nil }

func (r GetV2PetParams) Parse() GetV2PetParams { return r }

type GetV2PetResponse interface {
	writeGetV2Pet(http.ResponseWriter)
}

// ---------------------------------------------
// GetV3Pet -
// ---------------------------------------------

type GetV3PetHandlerFunc func(ctx context.Context, r GetV3PetRequest) GetV3PetResponse

func (f GetV3PetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetV3PetHTTPRequest(r)).writeGetV3Pet(w)
}

type GetV3PetRequest interface {
	HTTP() *http.Request
	Parse() GetV3PetParams
}

func GetV3PetHTTPRequest(r *http.Request) GetV3PetRequest {
	return getV3PetHTTPRequest{r}
}

type getV3PetHTTPRequest struct {
	Request *http.Request
}

func (r getV3PetHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getV3PetHTTPRequest) Parse() GetV3PetParams {
	return newGetV3PetParams(r.Request)
}

type GetV3PetParams struct {
}

func newGetV3PetParams(r *http.Request) (zero GetV3PetParams) {
	var params GetV3PetParams

	return params
}

func (r GetV3PetParams) HTTP() *http.Request { return nil }

func (r GetV3PetParams) Parse() GetV3PetParams { return r }

type GetV3PetResponse interface {
	writeGetV3Pet(http.ResponseWriter)
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
