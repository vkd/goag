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
// GetShops -
// ---------------------------------------------

type GetShopsHandlerFunc func(ctx context.Context, r GetShopsRequest) GetShopsResponse

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsHTTPRequest(r)).writeGetShops(w)
}

type GetShopsRequest interface {
	HTTP() *http.Request
	Parse() (GetShopsParams, error)
}

func GetShopsHTTPRequest(r *http.Request) GetShopsRequest {
	return getShopsHTTPRequest{r}
}

type getShopsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsHTTPRequest) Parse() (GetShopsParams, error) {
	return newGetShopsParams(r.Request)
}

type GetShopsParams struct {
	Query struct {
		Page Maybe[int32]
	}

	Headers struct {
		AccessKey Maybe[string]
	}
}

func newGetShopsParams(r *http.Request) (zero GetShopsParams, _ error) {
	var params GetShopsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int32", Err: err}
				}
				v := int32(vInt)
				params.Query.Page.Set(v)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("access-key")
			if len(hs) > 0 {
				v := hs[0]
				params.Headers.AccessKey.Set(v)
			}
		}
	}

	return params, nil
}

func (r GetShopsParams) HTTP() *http.Request { return nil }

func (r GetShopsParams) Parse() (GetShopsParams, error) { return r, nil }

type GetShopsResponse interface {
	writeGetShops(http.ResponseWriter)
}

func NewGetShopsResponse200() GetShopsResponse {
	var out GetShopsResponse200
	return out
}

type GetShopsResponse200 struct{}

func (r GetShopsResponse200) writeGetShops(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewGetShopsResponseDefault(code int) GetShopsResponse {
	var out GetShopsResponseDefault
	out.Code = code
	return out
}

type GetShopsResponseDefault struct {
	Code int
}

func (r GetShopsResponseDefault) writeGetShops(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

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
	Query struct {
		Page Maybe[int32]
	}

	Path struct {
		Shop string
	}

	Headers struct {
		RequestID Maybe[string]
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
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int32", Err: err}
				}
				v := int32(vInt)
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

// ---------------------------------------------
// PostShopsShop -
// ---------------------------------------------

type PostShopsShopHandlerFunc func(ctx context.Context, r PostShopsShopRequest) PostShopsShopResponse

func (f PostShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsShopHTTPRequest(r)).writePostShopsShop(w)
}

type PostShopsShopRequest interface {
	HTTP() *http.Request
	Parse() (PostShopsShopParams, error)
}

func PostShopsShopHTTPRequest(r *http.Request) PostShopsShopRequest {
	return postShopsShopHTTPRequest{r}
}

type postShopsShopHTTPRequest struct {
	Request *http.Request
}

func (r postShopsShopHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postShopsShopHTTPRequest) Parse() (PostShopsShopParams, error) {
	return newPostShopsShopParams(r.Request)
}

type PostShopsShopParams struct {
	Query struct {
		Page Maybe[int32]
	}

	Path struct {
		Shop string
	}

	Headers struct {
		QueryID Maybe[string]
	}
}

func newPostShopsShopParams(r *http.Request) (zero PostShopsShopParams, _ error) {
	var params PostShopsShopParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int32", Err: err}
				}
				v := int32(vInt)
				params.Query.Page.Set(v)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("query-id")
			if len(hs) > 0 {
				v := hs[0]
				params.Headers.QueryID.Set(v)
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

func (r PostShopsShopParams) HTTP() *http.Request { return nil }

func (r PostShopsShopParams) Parse() (PostShopsShopParams, error) { return r, nil }

type PostShopsShopResponse interface {
	writePostShopsShop(http.ResponseWriter)
}

func NewPostShopsShopResponse200() PostShopsShopResponse {
	var out PostShopsShopResponse200
	return out
}

type PostShopsShopResponse200 struct{}

func (r PostShopsShopResponse200) writePostShopsShop(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsShopResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewPostShopsShopResponseDefault(code int) PostShopsShopResponse {
	var out PostShopsShopResponseDefault
	out.Code = code
	return out
}

type PostShopsShopResponseDefault struct {
	Code int
}

func (r PostShopsShopResponseDefault) writePostShopsShop(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsShopResponseDefault) Write(w http.ResponseWriter) {
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
