package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// GetShopsShop -
// ---------------------------------------------

type GetShopsShopHandlerFunc func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopHTTPRequest(r)).writeGetShopsShop(w)
}

type GetShopsShopRequest interface {
	HTTP() *http.Request
	Parse() (GetShopsShopParams, error)
}

func GetShopsShopHTTPRequest(r *http.Request) GetShopsShopRequest {
	return getShopsShopHTTPRequest{r}
}

type getShopsShopHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsShopHTTPRequest) Parse() (GetShopsShopParams, error) {
	return newGetShopsShopParams(r.Request)
}

type GetShopsShopParams struct {
	Query GetShopsShopParamsQuery

	Path GetShopsShopParamsPath

	Headers GetShopsShopParamsHeaders
}

type GetShopsShopParamsQuery struct {
	Page Maybe[int32]
}

type GetShopsShopParamsPath struct {
	Shop string
}

type GetShopsShopParamsHeaders struct {
	RequestID Maybe[string]
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopParams, _ error) {
	var params GetShopsShopParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				vInt64, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int32", Err: err}
				}
				v := int32(vInt64)
				params.Query.Page.Set(v)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("request-id")
			if len(hs) > 0 {
				v := hs[0]
				params.Headers.RequestID.Set(v)
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}'")
		}
		p = p[7:] // "/shops/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "required"}
			}

			params.Path.Shop = vPath
		}
	}

	return params, nil
}

func (r GetShopsShopParams) HTTP() *http.Request { return nil }

func (r GetShopsShopParams) Parse() (GetShopsShopParams, error) { return r, nil }

type GetShopsShopResponse interface {
	writeGetShopsShop(http.ResponseWriter)
}

func NewGetShopsShopResponse200() GetShopsShopResponse {
	var out GetShopsShopResponse200
	return out
}

type GetShopsShopResponse200 struct{}

func (r GetShopsShopResponse200) writeGetShopsShop(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewGetShopsShopResponseDefault(code int) GetShopsShopResponse {
	var out GetShopsShopResponseDefault
	out.Code = code
	return out
}

type GetShopsShopResponseDefault struct {
	Code int
}

func (r GetShopsShopResponseDefault) writeGetShopsShop(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopResponseDefault) Write(w http.ResponseWriter) {
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
