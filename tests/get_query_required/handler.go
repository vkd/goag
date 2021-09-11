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

type GetPetsHandlerFunc func(r GetPetsRequester) GetPetsResponser

func (f GetPetsHandlerFunc) Handle(r GetPetsRequester) GetPetsResponser {
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

	Limit string
}

func newGetPetsParams(r *http.Request) (zero GetPetsRequest, _ error) {
	var params GetPetsRequest
	params.HTTPRequest = r

	{
		query := r.URL.Query()
		{
			q, ok := query["limit"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'limit': is required")
			}
			if ok && len(q) > 0 {
				v := q[0]
				params.Limit = v
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
