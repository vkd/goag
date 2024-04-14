package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// GetShopsShop -
// ---------------------------------------------

type GetShopsShopHandlerFunc func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopHTTPRequest(r)).Write(w)
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
	Query struct {
		Page Maybe[Page]

		PageReq Page

		Pages Maybe[[]Page]
	}

	Path struct {
		Shop Shop
	}

	Headers struct {
		RequestID Maybe[RequestID]
	}
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopParams, _ error) {
	var params GetShopsShopParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				var v Page
				err := v.Parse(q[0])
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse custom type", Err: err}
				}
				params.Query.Page = Just(v)
			}
		}
		{
			q, ok := query["page_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'page_req': is required")
			}
			if ok && len(q) > 0 {
				err := params.Query.PageReq.Parse(q[0])
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_req", Reason: "parse custom type", Err: err}
				}
			}
		}
		{
			q, ok := query["pages"]
			if ok && len(q) > 0 {
				v := make([]Page, len(q))
				for i := range q {
					err := v[i].Parse(q[i])
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "pages", Reason: "parse custom type", Err: err}
					}
				}
				params.Query.Pages = Just(v)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("request-id")
			if len(hs) > 0 {
				var v RequestID
				err := v.Parse(hs[0])
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "request-id", Reason: "parse custom type", Err: err}
				}
				params.Headers.RequestID = Just(v)
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

			err := params.Path.Shop.Parse(vPath)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse custom type", Err: err}
			}
		}
	}

	return params, nil
}

func (r GetShopsShopParams) HTTP() *http.Request { return nil }

func (r GetShopsShopParams) Parse() (GetShopsShopParams, error) { return r, nil }

type GetShopsShopResponse interface {
	getShopsShop()
	Write(w http.ResponseWriter)
}

func NewGetShopsShopResponse200() GetShopsShopResponse {
	var out GetShopsShopResponse200
	return out
}

type GetShopsShopResponse200 struct{}

func (r GetShopsShopResponse200) getShopsShop() {}

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

func (r GetShopsShopResponseDefault) getShopsShop() {}

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
		Value: v,
		IsSet: true,
	}
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
