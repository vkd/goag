package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// GetShops -
// GET /shops
// ---------------------------------------------

type GetShopsHandlerFunc func(ctx context.Context, r GetShopsRequest) GetShopsResponse

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsHTTPRequest(r)).writeGetShops(w)
}

func (GetShopsHandlerFunc) Path() string { return "/shops" }

func (GetShopsHandlerFunc) Method() string { return http.MethodGet }

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
	Query GetShopsParamsQuery

	Headers GetShopsParamsHeaders
}

type GetShopsParamsQuery struct {
	Page Maybe[int32]
}

type GetShopsParamsHeaders struct {
	AccessKey Maybe[string]
}

func newGetShopsParams(r *http.Request) (zero GetShopsParams, _ error) {
	var params GetShopsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				var vOpt int32
				if len(q) == 1 {
					vInt64, err := strconv.ParseInt(q[0], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int32", Err: err}
					}
					vOpt = int32(vInt64)
				} else {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "multiple values found: single value expected"}
				}
				params.Query.Page.Set(vOpt)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("access-key")
			if len(hs) > 0 {
				var vOpt string
				if len(hs) == 1 {
					vOpt = hs[0]
				} else {
					return zero, ErrParseParam{In: "header", Parameter: "access-key", Reason: "multiple values found: single value expected"}
				}
				params.Headers.AccessKey.Set(vOpt)
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
// GET /shops/{shop}
// ---------------------------------------------

type GetShopsShopHandlerFunc func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopHTTPRequest(r)).writeGetShopsShop(w)
}

func (GetShopsShopHandlerFunc) Path() string { return "/shops/{shop}" }

func (GetShopsShopHandlerFunc) Method() string { return http.MethodGet }

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
				var vOpt int32
				if len(q) == 1 {
					vInt64, err := strconv.ParseInt(q[0], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int32", Err: err}
					}
					vOpt = int32(vInt64)
				} else {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "multiple values found: single value expected"}
				}
				params.Query.Page.Set(vOpt)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("request-id")
			if len(hs) > 0 {
				var vOpt string
				if len(hs) == 1 {
					vOpt = hs[0]
				} else {
					return zero, ErrParseParam{In: "header", Parameter: "request-id", Reason: "multiple values found: single value expected"}
				}
				params.Headers.RequestID.Set(vOpt)
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
// POST /shops/{shop}
// ---------------------------------------------

type PostShopsShopHandlerFunc func(ctx context.Context, r PostShopsShopRequest) PostShopsShopResponse

func (f PostShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsShopHTTPRequest(r)).writePostShopsShop(w)
}

func (PostShopsShopHandlerFunc) Path() string { return "/shops/{shop}" }

func (PostShopsShopHandlerFunc) Method() string { return http.MethodPost }

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
	Query PostShopsShopParamsQuery

	Path PostShopsShopParamsPath

	Headers PostShopsShopParamsHeaders
}

type PostShopsShopParamsQuery struct {
	Page Maybe[int32]
}

type PostShopsShopParamsPath struct {
	Shop string
}

type PostShopsShopParamsHeaders struct {
	QueryID Maybe[string]
}

func newPostShopsShopParams(r *http.Request) (zero PostShopsShopParams, _ error) {
	var params PostShopsShopParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				var vOpt int32
				if len(q) == 1 {
					vInt64, err := strconv.ParseInt(q[0], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "parse int32", Err: err}
					}
					vOpt = int32(vInt64)
				} else {
					return zero, ErrParseParam{In: "query", Parameter: "page", Reason: "multiple values found: single value expected"}
				}
				params.Query.Page.Set(vOpt)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("query-id")
			if len(hs) > 0 {
				var vOpt string
				if len(hs) == 1 {
					vOpt = hs[0]
				} else {
					return zero, ErrParseParam{In: "header", Parameter: "query-id", Reason: "multiple values found: single value expected"}
				}
				params.Headers.QueryID.Set(vOpt)
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

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

type Nullable[T any] struct {
	IsSet bool
	Value T
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{}
}

func Pointer[T any](v T) Nullable[T] {
	return Nullable[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Nullable[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Nullable[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

var _ json.Marshaler = (*Nullable[any])(nil)

func (m Nullable[T]) MarshalJSON() ([]byte, error) {
	if m.IsSet {
		return json.Marshal(&m.Value)
	}
	return []byte(nullValue), nil
}

var _ json.Unmarshaler = (*Nullable[any])(nil)

const nullValue = "null"

var nullValueBs = []byte(nullValue)

func (m *Nullable[T]) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, nullValueBs) {
		m.IsSet = false
		return nil
	}
	m.IsSet = true
	return json.Unmarshal(bs, &m.Value)
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
