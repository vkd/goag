package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostLogin -
// ---------------------------------------------

type PostLoginHandlerFunc func(ctx context.Context, r PostLoginRequest) PostLoginResponse

func (f PostLoginHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostLoginHTTPRequest(r)).Write(w)
}

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
	postLogin()
	Write(w http.ResponseWriter)
}

func NewPostLoginResponse200() PostLoginResponse {
	var out PostLoginResponse200
	return out
}

type PostLoginResponse200 struct{}

func (r PostLoginResponse200) postLogin() {}

func (r PostLoginResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewPostLoginResponse401() PostLoginResponse {
	var out PostLoginResponse401
	return out
}

type PostLoginResponse401 struct{}

func (r PostLoginResponse401) postLogin() {}

func (r PostLoginResponse401) Write(w http.ResponseWriter) {
	w.WriteHeader(401)
}

// ---------------------------------------------
// PostShops -
// ---------------------------------------------

type PostShopsHandlerFunc func(ctx context.Context, r PostShopsRequest) PostShopsResponse

func (f PostShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsHTTPRequest(r)).Write(w)
}

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
	Headers struct {

		// Authorization - JWT
		Authorization string
	}
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
				params.Headers.Authorization = hs[0]
			}
		}
	}

	return params, nil
}

func (r PostShopsParams) HTTP() *http.Request { return nil }

func (r PostShopsParams) Parse() (PostShopsParams, error) { return r, nil }

type PostShopsResponse interface {
	postShops()
	Write(w http.ResponseWriter)
}

func NewPostShopsResponse200() PostShopsResponse {
	var out PostShopsResponse200
	return out
}

type PostShopsResponse200 struct{}

func (r PostShopsResponse200) postShops() {}

func (r PostShopsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewPostShopsResponse401() PostShopsResponse {
	var out PostShopsResponse401
	return out
}

type PostShopsResponse401 struct{}

func (r PostShopsResponse401) postShops() {}

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
