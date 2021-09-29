package test

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ---------------------------------------------
// GetPets -
// ---------------------------------------------

type GetPetsHandlerFunc func(r GetPetsRequester) GetPetsResponder

func (f GetPetsHandlerFunc) Handle(r GetPetsRequester) GetPetsResponder {
	return f(r)
}

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetPetsParams{Request: r}).writeGetPetsResponse(w)
}

type GetPetsRequester interface {
	Parse() (GetPetsRequest, error)
}

type requestGetPetsParams struct {
	Request *http.Request
}

func (r requestGetPetsParams) Parse() (GetPetsRequest, error) {
	return newGetPetsParams(r.Request)
}

type GetPetsRequest struct {
	HTTPRequest *http.Request

	Tag  []string
	Page []int64
}

func newGetPetsParams(r *http.Request) (zero GetPetsRequest, _ error) {
	var params GetPetsRequest
	params.HTTPRequest = r

	{
		query := r.URL.Query()
		{
			q, ok := query["tag"]
			if ok && len(q) > 0 {
				params.Tag = q
			}
		}
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				params.Page = make([]int64, len(q))
				for i := range q {
					vInt, err := strconv.ParseInt(q[i], 10, 64)
					if err != nil {
						return zero, ErrParseQueryParam{Name: "page", Err: fmt.Errorf("parse int64: %w", err)}
					}
					v1 := vInt
					params.Page[i] = v1
				}
			}
		}
	}

	return params, nil
}

type GetPetsResponder interface {
	writeGetPetsResponse(w http.ResponseWriter)
}

func GetPetsResponse200() GetPetsResponder {
	var out getPetsResponse200
	return out
}

type getPetsResponse200 struct{}

func (r getPetsResponse200) writeGetPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func GetPetsResponseDefault(code int) GetPetsResponder {
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
