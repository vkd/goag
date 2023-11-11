package test

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// GetPetsPetIDNames -
// ---------------------------------------------

type GetPetsPetIDNamesHandlerFunc func(r GetPetsPetIDNamesRequestParser) GetPetsPetIDNamesResponse

func (f GetPetsPetIDNamesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetsPetIDNamesHTTPRequest(r)).Write(w)
}

type GetPetsPetIDNamesRequestParser interface {
	Parse() (GetPetsPetIDNamesRequest, error)
}

func GetPetsPetIDNamesHTTPRequest(r *http.Request) GetPetsPetIDNamesRequestParser {
	return getPetsPetIDNamesHTTPRequest{r}
}

type getPetsPetIDNamesHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDNamesHTTPRequest) Parse() (GetPetsPetIDNamesRequest, error) {
	return newGetPetsPetIDNamesParams(r.Request)
}

type GetPetsPetIDNamesRequest struct {
	HTTPRequest *http.Request

	Path struct {
		PetID string
	}
}

func newGetPetsPetIDNamesParams(r *http.Request) (zero GetPetsPetIDNamesRequest, _ error) {
	var params GetPetsPetIDNamesRequest
	params.HTTPRequest = r

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

func (r GetPetsPetIDNamesRequest) Parse() (GetPetsPetIDNamesRequest, error) { return r, nil }

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

type GetPetsPetIDShopsHandlerFunc func(r GetPetsPetIDShopsRequestParser) GetPetsPetIDShopsResponse

func (f GetPetsPetIDShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetsPetIDShopsHTTPRequest(r)).Write(w)
}

type GetPetsPetIDShopsRequestParser interface {
	Parse() (GetPetsPetIDShopsRequest, error)
}

func GetPetsPetIDShopsHTTPRequest(r *http.Request) GetPetsPetIDShopsRequestParser {
	return getPetsPetIDShopsHTTPRequest{r}
}

type getPetsPetIDShopsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDShopsHTTPRequest) Parse() (GetPetsPetIDShopsRequest, error) {
	return newGetPetsPetIDShopsParams(r.Request)
}

type GetPetsPetIDShopsRequest struct {
	HTTPRequest *http.Request

	Path struct {
		PetID string
	}
}

func newGetPetsPetIDShopsParams(r *http.Request) (zero GetPetsPetIDShopsRequest, _ error) {
	var params GetPetsPetIDShopsRequest
	params.HTTPRequest = r

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

func (r GetPetsPetIDShopsRequest) Parse() (GetPetsPetIDShopsRequest, error) { return r, nil }

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
