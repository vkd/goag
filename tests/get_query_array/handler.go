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

type GetPetsHandlerFunc func(r GetPetsRequestParser) GetPetsResponse

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetsHTTPRequest(r)).Write(w)
}

type GetPetsRequestParser interface {
	Parse() (GetPetsRequest, error)
}

func GetPetsHTTPRequest(r *http.Request) GetPetsRequestParser {
	return getPetsHTTPRequest{r}
}

type getPetsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsHTTPRequest) Parse() (GetPetsRequest, error) {
	return newGetPetsParams(r.Request)
}

type GetPetsRequest struct {
	HTTPRequest *http.Request

	Query struct {
		Tag []string

		Page []int64
	}
}

func newGetPetsParams(r *http.Request) (zero GetPetsRequest, _ error) {
	var params GetPetsRequest
	params.HTTPRequest = r

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
					vInt, err := strconv.ParseInt(q[i], 10, 64)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int64", Err: err}
					}
					v1 := int64(vInt)
					params.Query.Page[i] = v1
				}
			}
		}
	}

	return params, nil
}

func (r GetPetsRequest) Parse() (GetPetsRequest, error) { return r, nil }

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
