package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// PostShopsShopStringShopSchemaPets -
// POST /shops/{shop_string}/{shop_schema}/pets
// ---------------------------------------------

type PostShopsShopStringShopSchemaPetsHandlerFunc func(ctx context.Context, r PostShopsShopStringShopSchemaPetsRequest) PostShopsShopStringShopSchemaPetsResponse

func (f PostShopsShopStringShopSchemaPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsShopStringShopSchemaPetsHTTPRequest(r)).writePostShopsShopStringShopSchemaPets(w)
}

func (PostShopsShopStringShopSchemaPetsHandlerFunc) Path() string {
	return "/shops/{shop_string}/{shop_schema}/pets"
}

func (PostShopsShopStringShopSchemaPetsHandlerFunc) Method() string { return http.MethodPost }

type PostShopsShopStringShopSchemaPetsRequest interface {
	HTTP() *http.Request
	Parse() (PostShopsShopStringShopSchemaPetsParams, error)
}

func PostShopsShopStringShopSchemaPetsHTTPRequest(r *http.Request) PostShopsShopStringShopSchemaPetsRequest {
	return postShopsShopStringShopSchemaPetsHTTPRequest{r}
}

type postShopsShopStringShopSchemaPetsHTTPRequest struct {
	Request *http.Request
}

func (r postShopsShopStringShopSchemaPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postShopsShopStringShopSchemaPetsHTTPRequest) Parse() (PostShopsShopStringShopSchemaPetsParams, error) {
	return newPostShopsShopStringShopSchemaPetsParams(r.Request)
}

type PostShopsShopStringShopSchemaPetsParams struct {
	Path PostShopsShopStringShopSchemaPetsParamsPath
}

type PostShopsShopStringShopSchemaPetsParamsPath struct {
	ShopString string

	ShopSchema Shop
}

func newPostShopsShopStringShopSchemaPetsParams(r *http.Request) (zero PostShopsShopStringShopSchemaPetsParams, _ error) {
	var params PostShopsShopStringShopSchemaPetsParams

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop_string}/{shop_schema}/pets'")
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
				return zero, ErrParseParam{In: "path", Parameter: "shop_string", Reason: "required"}
			}

			params.Path.ShopString = vPath
		}

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop_string}/{shop_schema}/pets'")
		}
		p = p[1:] // "/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "shop_schema", Reason: "required"}
			}

			vString := vPath
			params.Path.ShopSchema = NewShop(vString)
		}

		if !strings.HasPrefix(p, "/pets") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop_string}/{shop_schema}/pets'")
		}
		p = p[5:] // "/pets"
	}

	return params, nil
}

func (r PostShopsShopStringShopSchemaPetsParams) HTTP() *http.Request { return nil }

func (r PostShopsShopStringShopSchemaPetsParams) Parse() (PostShopsShopStringShopSchemaPetsParams, error) {
	return r, nil
}

type PostShopsShopStringShopSchemaPetsResponse interface {
	writePostShopsShopStringShopSchemaPets(http.ResponseWriter)
}

func NewPostShopsShopStringShopSchemaPetsResponse200() PostShopsShopStringShopSchemaPetsResponse {
	var out PostShopsShopStringShopSchemaPetsResponse200
	return out
}

// PostShopsShopStringShopSchemaPetsResponse200 - OK response
type PostShopsShopStringShopSchemaPetsResponse200 struct{}

func (r PostShopsShopStringShopSchemaPetsResponse200) writePostShopsShopStringShopSchemaPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsShopStringShopSchemaPetsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
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
