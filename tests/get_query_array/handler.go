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

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetsParams{Request: r}).writeGetPetsResponse(w)
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

	// Query parameters
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
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int64", Err: err}
					}
					v1 := int64(vInt)
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
