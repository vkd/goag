package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPets -
// ---------------------------------------------

func GetPetsHandler(h GetPetsHandlerer) http.Handler {
	return GetPetsHandlerFunc(h.Handle, h.InvalidResponce)
}

func GetPetsHandlerFunc(fn FuncGetPets, invalidFn FuncGetPetsInvalidResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := newGetPetsParams(r)
		if err != nil {
			invalidFn(err).writeGetPetsResponse(w)
			return
		}

		fn(params).writeGetPetsResponse(w)
	}
}

type GetPetsHandlerer interface {
	Handle(GetPetsParams) GetPetsResponser
	InvalidResponce(error) GetPetsResponser
}

func NewGetPetsHandlerer(fn FuncGetPets, invalidFn FuncGetPetsInvalidResponse) GetPetsHandlerer {
	return privateGetPetsHandlerer{
		FuncGetPets:                fn,
		FuncGetPetsInvalidResponse: invalidFn,
	}
}

type privateGetPetsHandlerer struct {
	FuncGetPets
	FuncGetPetsInvalidResponse
}

type FuncGetPets func(GetPetsParams) GetPetsResponser

func (f FuncGetPets) Handle(params GetPetsParams) GetPetsResponser { return f(params) }

type FuncGetPetsInvalidResponse func(error) GetPetsResponser

func (f FuncGetPetsInvalidResponse) InvalidResponce(err error) GetPetsResponser { return f(err) }

type GetPetsParams struct {
	Request *http.Request

	Limit string
}

func newGetPetsParams(r *http.Request) (zero GetPetsParams, _ error) {
	var params GetPetsParams
	params.Request = r

	{
		query := r.URL.Query()
		{
			q, ok := query["limit"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'limit': is required")
			}
			if ok && len(q) > 0 {
				params.Limit = q[0]
			}
		}
	}

	return params, nil
}

type GetPetsResponser interface {
	writeGetPetsResponse(w http.ResponseWriter)
}

func GetPetsResponse200() GetPetsResponser {
	var out getPetsResponse200
	return out
}

type getPetsResponse200 struct{}

func (r getPetsResponse200) writeGetPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func GetPetsResponseDefault(code int) GetPetsResponser {
	var out getPetsResponseDefault
	out.Code = code
	return out
}

type getPetsResponseDefault struct {
	Code int
}

func (r getPetsResponseDefault) writeGetPetsResponse(w http.ResponseWriter) {
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
