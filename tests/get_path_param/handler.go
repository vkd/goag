package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// GetPetsPetID -
// ---------------------------------------------

type GetPetsPetIDHandlerFunc func(GetPetsPetIDParamsParser) GetPetsPetIDResponser

func (f GetPetsPetIDHandlerFunc) Handle(p GetPetsPetIDParamsParser) GetPetsPetIDResponser {
	return f(p)
}

func (f GetPetsPetIDHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetsPetIDParams{Request: r}).writeGetPetsPetIDResponse(w)
}

type GetPetsPetIDParamsParser interface {
	Parse() (GetPetsPetIDParams, error)
}

type requestGetPetsPetIDParams struct {
	Request *http.Request
}

func (p requestGetPetsPetIDParams) Parse() (GetPetsPetIDParams, error) {
	return newGetPetsPetIDParams(p.Request)
}

type GetPetsPetIDParams struct {
	HTTPRequest *http.Request

	PetID int
}

func newGetPetsPetIDParams(r *http.Request) (zero GetPetsPetIDParams, _ error) {
	var params GetPetsPetIDParams
	params.HTTPRequest = r

	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{petId}'")
		}
		p = p[6:]

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			vInt, err := strconv.Atoi(vPath)
			if err != nil {
				return zero, ErrParsePathParam{Name: "petId", Err: fmt.Errorf("parse int: %w", err)}
			}
			v := vInt
			params.PetID = v
		}
	}

	return params, nil
}

type GetPetsPetIDResponser interface {
	writeGetPetsPetIDResponse(w http.ResponseWriter)
}

func GetPetsPetIDResponse200() GetPetsPetIDResponser {
	var out getPetsPetIDResponse200
	return out
}

type getPetsPetIDResponse200 struct{}

func (r getPetsPetIDResponse200) writeGetPetsPetIDResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func GetPetsPetIDResponseDefault(code int) GetPetsPetIDResponser {
	var out getPetsPetIDResponseDefault
	out.Code = code
	return out
}

type getPetsPetIDResponseDefault struct {
	Code int
}

func (r getPetsPetIDResponseDefault) writeGetPetsPetIDResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
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
