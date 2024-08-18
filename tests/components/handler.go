package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// PostPets -
// ---------------------------------------------

type PostPetsHandlerFunc func(ctx context.Context, r PostPetsRequest) PostPetsResponse

func (f PostPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostPetsHTTPRequest(r)).writePostPets(w)
}

type PostPetsRequest interface {
	HTTP() *http.Request
	Parse() (PostPetsParams, error)
}

func PostPetsHTTPRequest(r *http.Request) PostPetsRequest {
	return postPetsHTTPRequest{r}
}

type postPetsHTTPRequest struct {
	Request *http.Request
}

func (r postPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postPetsHTTPRequest) Parse() (PostPetsParams, error) {
	return newPostPetsParams(r.Request)
}

type PostPetsParams struct {
	Body CreatePetJSON
}

func newPostPetsParams(r *http.Request) (zero PostPetsParams, _ error) {
	var params PostPetsParams

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r PostPetsParams) HTTP() *http.Request { return nil }

func (r PostPetsParams) Parse() (PostPetsParams, error) { return r, nil }

type PostPetsResponse interface {
	writePostPets(http.ResponseWriter)
}

func NewPostPetsResponse200JSON(body Pets) PostPetsResponse {
	var out PostPetsResponse200JSON
	out.Body = body
	return out
}

// PostPetsResponse200JSON - Pets response
type PostPetsResponse200JSON struct {
	Body Pets
}

func (r PostPetsResponse200JSON) writePostPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostPetsResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "PostPetsResponse200JSON")
}

// ---------------------------------------------
// PostShops -
// ---------------------------------------------

type PostShopsHandlerFunc func(ctx context.Context, r PostShopsRequest) PostShopsResponse

func (f PostShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsHTTPRequest(r)).writePostShops(w)
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
	Body CreatePetJSON
}

func newPostShopsParams(r *http.Request) (zero PostShopsParams, _ error) {
	var params PostShopsParams

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r PostShopsParams) HTTP() *http.Request { return nil }

func (r PostShopsParams) Parse() (PostShopsParams, error) { return r, nil }

type PostShopsResponse interface {
	writePostShops(http.ResponseWriter)
}

func NewPostShopsResponse200JSON(body Pets) PostShopsResponse {
	var out PostShopsResponse200JSON
	out.Body = body
	return out
}

// PostShopsResponse200JSON - Pets response
type PostShopsResponse200JSON struct {
	Body Pets
}

func (r PostShopsResponse200JSON) writePostShops(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "PostShopsResponse200JSON")
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

func writeJSON(w io.Writer, v interface{}, name string) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		LogError(fmt.Errorf("write json response %q: %w", name, err))
	}
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
