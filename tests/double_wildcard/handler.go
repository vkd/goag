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

type GetPetsPetIDNamesHandlerFunc func(r GetPetsPetIDNamesRequester) GetPetsPetIDNamesResponder

func (f GetPetsPetIDNamesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetsPetIDNamesParams{Request: r}).writeGetPetsPetIDNamesResponse(w)
}

type GetPetsPetIDNamesRequester interface {
	Parse() (GetPetsPetIDNamesRequest, error)
}

type requestGetPetsPetIDNamesParams struct {
	Request *http.Request
}

func (r requestGetPetsPetIDNamesParams) Parse() (GetPetsPetIDNamesRequest, error) {
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

type GetPetsPetIDNamesResponder interface {
	writeGetPetsPetIDNamesResponse(w http.ResponseWriter)
}

func GetPetsPetIDNamesResponse200() GetPetsPetIDNamesResponder {
	var out getPetsPetIDNamesResponse200
	return out
}

type getPetsPetIDNamesResponse200 struct{}

func (r getPetsPetIDNamesResponse200) writeGetPetsPetIDNamesResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

// ---------------------------------------------
// GetPetsPetIDShops -
// ---------------------------------------------

type GetPetsPetIDShopsHandlerFunc func(r GetPetsPetIDShopsRequester) GetPetsPetIDShopsResponder

func (f GetPetsPetIDShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetsPetIDShopsParams{Request: r}).writeGetPetsPetIDShopsResponse(w)
}

type GetPetsPetIDShopsRequester interface {
	Parse() (GetPetsPetIDShopsRequest, error)
}

type requestGetPetsPetIDShopsParams struct {
	Request *http.Request
}

func (r requestGetPetsPetIDShopsParams) Parse() (GetPetsPetIDShopsRequest, error) {
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

type GetPetsPetIDShopsResponder interface {
	writeGetPetsPetIDShopsResponse(w http.ResponseWriter)
}

func GetPetsPetIDShopsResponse200() GetPetsPetIDShopsResponder {
	var out getPetsPetIDShopsResponse200
	return out
}

type getPetsPetIDShopsResponse200 struct{}

func (r getPetsPetIDShopsResponse200) writeGetPetsPetIDShopsResponse(w http.ResponseWriter) {
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
