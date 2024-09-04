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
	f(r.Context(), GetPetsHTTPRequest(r)).writeGetPets(w)
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
	Query GetPetsParamsQuery
}

type GetPetsParamsQuery struct {
	Tag Maybe[[]string]

	Page Maybe[[]int64]
}

func newGetPetsParams(r *http.Request) (zero GetPetsParams, _ error) {
	var params GetPetsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["tag"]
			if ok && len(q) > 0 {
				v := q
				params.Query.Tag.Set(v)
			}
		}
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				v := make([]int64, len(q))
				for i := range q {
					var err error
					v[i], err = strconv.ParseInt(q[i], 10, 64)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int64", Err: err}
					}
				}
				params.Query.Page.Set(v)
			}
		}
	}

	return params, nil
}

func (r GetPetsParams) HTTP() *http.Request { return nil }

func (r GetPetsParams) Parse() (GetPetsParams, error) { return r, nil }

type GetPetsResponse interface {
	writeGetPets(http.ResponseWriter)
}

func NewGetPetsResponse200() GetPetsResponse {
	var out GetPetsResponse200
	return out
}

type GetPetsResponse200 struct{}

func (r GetPetsResponse200) writeGetPets(w http.ResponseWriter) {
	r.Write(w)
}

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

func (r GetPetsResponseDefault) writeGetPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

type Nullable[T any] struct {
	IsSet bool
	Value T
}

func Ptr[T any](v T) Nullable[T] {
	return Nullable[T]{
		IsSet: true,
		Value: v,
	}
}

func (m *Nullable[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
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
