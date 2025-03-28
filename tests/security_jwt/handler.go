package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostLogin -
// POST /login
// ---------------------------------------------

type PostLoginHandlerFunc func(ctx context.Context, r PostLoginRequest) PostLoginResponse

func (f PostLoginHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostLoginHTTPRequest(r)).writePostLogin(w)
}

func (PostLoginHandlerFunc) Path() string { return "/login" }

func (PostLoginHandlerFunc) Method() string { return http.MethodPost }

type PostLoginRequest interface {
	HTTP() *http.Request
	Parse() PostLoginParams
}

func PostLoginHTTPRequest(r *http.Request) PostLoginRequest {
	return postLoginHTTPRequest{r}
}

type postLoginHTTPRequest struct {
	Request *http.Request
}

func (r postLoginHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postLoginHTTPRequest) Parse() PostLoginParams {
	return newPostLoginParams(r.Request)
}

type PostLoginParams struct {
}

func newPostLoginParams(r *http.Request) (zero PostLoginParams) {
	var params PostLoginParams

	return params
}

func (r PostLoginParams) HTTP() *http.Request { return nil }

func (r PostLoginParams) Parse() PostLoginParams { return r }

type PostLoginResponse interface {
	writePostLogin(http.ResponseWriter)
}

func NewPostLoginResponse200() PostLoginResponse {
	var out PostLoginResponse200
	return out
}

type PostLoginResponse200 struct{}

func (r PostLoginResponse200) writePostLogin(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostLoginResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewPostLoginResponse401() PostLoginResponse {
	var out PostLoginResponse401
	return out
}

type PostLoginResponse401 struct{}

func (r PostLoginResponse401) writePostLogin(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostLoginResponse401) Write(w http.ResponseWriter) {
	w.WriteHeader(401)
}

// ---------------------------------------------
// PostShops -
// POST /shops
// ---------------------------------------------

type PostShopsHandlerFunc func(ctx context.Context, r PostShopsRequest) PostShopsResponse

func (f PostShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsHTTPRequest(r)).writePostShops(w)
}

func (PostShopsHandlerFunc) Path() string { return "/shops" }

func (PostShopsHandlerFunc) Method() string { return http.MethodPost }

type PostShopsRequest interface {
	HTTP() *http.Request
	Parse() (PostShopsParams, error)
}

func PostShopsHTTPRequest(r *http.Request) PostShopsRequest {
	return postShopsHTTPRequest{r}
}

type postShopsHTTPRequest struct {
	Request *http.Request
}

func (r postShopsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postShopsHTTPRequest) Parse() (PostShopsParams, error) {
	return newPostShopsParams(r.Request)
}

type PostShopsParams struct {
	Headers PostShopsParamsHeaders
}

type PostShopsParamsHeaders struct {

	// Authorization - JWT
	Authorization string
}

func newPostShopsParams(r *http.Request) (zero PostShopsParams, _ error) {
	var params PostShopsParams

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("Authorization")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'Authorization': is required")
			}
			if len(hs) > 0 {
				if len(hs) == 1 {
					params.Headers.Authorization = hs[0]
				} else {
					return zero, ErrParseParam{In: "header", Parameter: "Authorization", Reason: "multiple values found: single value expected"}
				}
			}
		}
	}

	return params, nil
}

func (r PostShopsParams) HTTP() *http.Request { return nil }

func (r PostShopsParams) Parse() (PostShopsParams, error) { return r, nil }

type PostShopsResponse interface {
	writePostShops(http.ResponseWriter)
}

func NewPostShopsResponse200() PostShopsResponse {
	var out PostShopsResponse200
	return out
}

type PostShopsResponse200 struct{}

func (r PostShopsResponse200) writePostShops(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewPostShopsResponse401() PostShopsResponse {
	var out PostShopsResponse401
	return out
}

type PostShopsResponse401 struct{}

func (r PostShopsResponse401) writePostShops(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsResponse401) Write(w http.ResponseWriter) {
	w.WriteHeader(401)
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
