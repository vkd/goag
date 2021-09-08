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

func GetPetsPetIDHandler(h GetPetsPetIDHandlerer) http.Handler {
	return GetPetsPetIDHandlerFunc(h.Handle, h.InvalidResponce)
}

func GetPetsPetIDHandlerFunc(fn FuncGetPetsPetID, invalidFn FuncGetPetsPetIDInvalidResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := newGetPetsPetIDParams(r)
		if err != nil {
			invalidFn(err).writeGetPetsPetIDResponse(w)
			return
		}

		fn(params).writeGetPetsPetIDResponse(w)
	}
}

type GetPetsPetIDHandlerer interface {
	Handle(GetPetsPetIDParams) GetPetsPetIDResponser
	InvalidResponce(error) GetPetsPetIDResponser
}

func NewGetPetsPetIDHandlerer(fn FuncGetPetsPetID, invalidFn FuncGetPetsPetIDInvalidResponse) GetPetsPetIDHandlerer {
	return privateGetPetsPetIDHandlerer{
		FuncGetPetsPetID:                fn,
		FuncGetPetsPetIDInvalidResponse: invalidFn,
	}
}

type privateGetPetsPetIDHandlerer struct {
	FuncGetPetsPetID
	FuncGetPetsPetIDInvalidResponse
}

type FuncGetPetsPetID func(GetPetsPetIDParams) GetPetsPetIDResponser

func (f FuncGetPetsPetID) Handle(params GetPetsPetIDParams) GetPetsPetIDResponser { return f(params) }

type FuncGetPetsPetIDInvalidResponse func(error) GetPetsPetIDResponser

func (f FuncGetPetsPetIDInvalidResponse) InvalidResponce(err error) GetPetsPetIDResponser {
	return f(err)
}

type GetPetsPetIDParams struct {
	Request *http.Request

	PetID int
}

func newGetPetsPetIDParams(r *http.Request) (zero GetPetsPetIDParams, _ error) {
	var params GetPetsPetIDParams
	params.Request = r

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
			v := p[:idx]
			p = p[idx:]

			vInt, err := strconv.Atoi(v)
			if err != nil {
				return zero, ErrParsePathParam{Name: "petId", Err: fmt.Errorf("parse int: %w", err)}
			}
			params.PetID = vInt
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
