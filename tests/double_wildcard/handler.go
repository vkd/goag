package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// GetPetsPetIDNames -
// ---------------------------------------------

type GetPetsPetIDNamesHandlerFunc func(ctx context.Context, r GetPetsPetIDNamesRequest) GetPetsPetIDNamesResponse

func (f GetPetsPetIDNamesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsPetIDNamesHTTPRequest(r)).Write(w)
}

type GetPetsPetIDNamesRequest interface {
	HTTP() *http.Request
	Parse() (GetPetsPetIDNamesParams, error)
}

func GetPetsPetIDNamesHTTPRequest(r *http.Request) GetPetsPetIDNamesRequest {
	return getPetsPetIDNamesHTTPRequest{r}
}

type getPetsPetIDNamesHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDNamesHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsPetIDNamesHTTPRequest) Parse() (GetPetsPetIDNamesParams, error) {
	return newGetPetsPetIDNamesParams(r.Request)
}

type GetPetsPetIDNamesParams struct {
	Path struct {
		PetID string
	}
}

func newGetPetsPetIDNamesParams(r *http.Request) (zero GetPetsPetIDNamesParams, _ error) {
	var params GetPetsPetIDNamesParams

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/names'")
		}
		p = p[6:] // "/pets/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "pet_id", Reason: "required"}
			}

			v := vPath
			params.Path.PetID = v
		}

		if !strings.HasPrefix(p, "/names") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/names'")
		}
		p = p[6:] // "/names"
	}

	return params, nil
}

func (r GetPetsPetIDNamesParams) HTTP() *http.Request { return nil }

func (r GetPetsPetIDNamesParams) Parse() (GetPetsPetIDNamesParams, error) { return r, nil }

type GetPetsPetIDNamesResponse interface {
	getPetsPetIDNames()
	Write(w http.ResponseWriter)
}

func NewGetPetsPetIDNamesResponse200() GetPetsPetIDNamesResponse {
	var out GetPetsPetIDNamesResponse200
	return out
}

type GetPetsPetIDNamesResponse200 struct{}

func (r GetPetsPetIDNamesResponse200) getPetsPetIDNames() {}

func (r GetPetsPetIDNamesResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

// ---------------------------------------------
// GetPetsPetIDShops -
// ---------------------------------------------

type GetPetsPetIDShopsHandlerFunc func(ctx context.Context, r GetPetsPetIDShopsRequest) GetPetsPetIDShopsResponse

func (f GetPetsPetIDShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsPetIDShopsHTTPRequest(r)).Write(w)
}

type GetPetsPetIDShopsRequest interface {
	HTTP() *http.Request
	Parse() (GetPetsPetIDShopsParams, error)
}

func GetPetsPetIDShopsHTTPRequest(r *http.Request) GetPetsPetIDShopsRequest {
	return getPetsPetIDShopsHTTPRequest{r}
}

type getPetsPetIDShopsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDShopsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsPetIDShopsHTTPRequest) Parse() (GetPetsPetIDShopsParams, error) {
	return newGetPetsPetIDShopsParams(r.Request)
}

type GetPetsPetIDShopsParams struct {
	Path struct {
		PetID string
	}
}

func newGetPetsPetIDShopsParams(r *http.Request) (zero GetPetsPetIDShopsParams, _ error) {
	var params GetPetsPetIDShopsParams

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/shops'")
		}
		p = p[6:] // "/pets/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "pet_id", Reason: "required"}
			}

			v := vPath
			params.Path.PetID = v
		}

		if !strings.HasPrefix(p, "/shops") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}/shops'")
		}
		p = p[6:] // "/shops"
	}

	return params, nil
}

func (r GetPetsPetIDShopsParams) HTTP() *http.Request { return nil }

func (r GetPetsPetIDShopsParams) Parse() (GetPetsPetIDShopsParams, error) { return r, nil }

type GetPetsPetIDShopsResponse interface {
	getPetsPetIDShops()
	Write(w http.ResponseWriter)
}

func NewGetPetsPetIDShopsResponse200() GetPetsPetIDShopsResponse {
	var out GetPetsPetIDShopsResponse200
	return out
}

type GetPetsPetIDShopsResponse200 struct{}

func (r GetPetsPetIDShopsResponse200) getPetsPetIDShops() {}

func (r GetPetsPetIDShopsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
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
