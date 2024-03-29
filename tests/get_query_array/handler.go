package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ---------------------------------------------
// GetPets -
// ---------------------------------------------

type GetPetsHandlerFunc func(ctx context.Context, r GetPetsRequest) GetPetsResponse

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetsHTTPRequest(r)).Write(w)
}

type GetPetsRequest interface {
	HTTP() *http.Request
	Parse() (GetPetsParams, error)
}

func GetPetsHTTPRequest(r *http.Request) GetPetsRequest {
	return getPetsHTTPRequest{r}
}

type getPetsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetsHTTPRequest) Parse() (GetPetsParams, error) {
	return newGetPetsParams(r.Request)
}

type GetPetsParams struct {
	Query struct {
		Tag []string

		Page []int64
	}
}

func newGetPetsParams(r *http.Request) (zero GetPetsParams, _ error) {
	var params GetPetsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["tag"]
			if ok && len(q) > 0 {
				params.Query.Tag = q
			}
		}
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				params.Query.Page = make([]int64, len(q))
				for i := range q {
					var err error
					params.Query.Page[i], err = strconv.ParseInt(q[i], 10, 64)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int64", Err: err}
					}
				}
			}
		}
	}

	return params, nil
}

func (r GetPetsParams) HTTP() *http.Request { return nil }

func (r GetPetsParams) Parse() (GetPetsParams, error) { return r, nil }

type GetPetsResponse interface {
	getPets()
	Write(w http.ResponseWriter)
}

func NewGetPetsResponse200() GetPetsResponse {
	var out GetPetsResponse200
	return out
}

type GetPetsResponse200 struct{}

func (r GetPetsResponse200) getPets() {}

func (r GetPetsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewGetPetsResponseDefault(code int) GetPetsResponse {
	var out GetPetsResponseDefault
	out.Code = code
	return out
}

type GetPetsResponseDefault struct {
	Code int
}

func (r GetPetsResponseDefault) getPets() {}

func (r GetPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
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
