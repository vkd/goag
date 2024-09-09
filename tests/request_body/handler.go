package test

import (
	"context"
	"encoding/json"
	"fmt"
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
	Body NewPet
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

func NewPostPetsResponse201() PostPetsResponse {
	var out PostPetsResponse201
	return out
}

type PostPetsResponse201 struct{}

func (r PostPetsResponse201) writePostPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostPetsResponse201) Write(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func NewPostPetsResponseDefault(code int) PostPetsResponse {
	var out PostPetsResponseDefault
	out.Code = code
	return out
}

type PostPetsResponseDefault struct {
	Code int
}

func (r PostPetsResponseDefault) writePostPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// PostPets2 -
// ---------------------------------------------

type PostPets2HandlerFunc func(ctx context.Context, r PostPets2Request) PostPets2Response

func (f PostPets2HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostPets2HTTPRequest(r)).writePostPets2(w)
}

type PostPets2Request interface {
	HTTP() *http.Request
	Parse() (PostPets2Params, error)
}

func PostPets2HTTPRequest(r *http.Request) PostPets2Request {
	return postPets2HTTPRequest{r}
}

type postPets2HTTPRequest struct {
	Request *http.Request
}

func (r postPets2HTTPRequest) HTTP() *http.Request { return r.Request }

func (r postPets2HTTPRequest) Parse() (PostPets2Params, error) {
	return newPostPets2Params(r.Request)
}

type PostPets2Params struct {
	Body Pets2JSON
}

func newPostPets2Params(r *http.Request) (zero PostPets2Params, _ error) {
	var params PostPets2Params

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r PostPets2Params) HTTP() *http.Request { return nil }

func (r PostPets2Params) Parse() (PostPets2Params, error) { return r, nil }

type PostPets2Response interface {
	writePostPets2(http.ResponseWriter)
}

func NewPostPets2Response201() PostPets2Response {
	var out PostPets2Response201
	return out
}

type PostPets2Response201 struct{}

func (r PostPets2Response201) writePostPets2(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostPets2Response201) Write(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func NewPostPets2ResponseDefault(code int) PostPets2Response {
	var out PostPets2ResponseDefault
	out.Code = code
	return out
}

type PostPets2ResponseDefault struct {
	Code int
}

func (r PostPets2ResponseDefault) writePostPets2(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostPets2ResponseDefault) Write(w http.ResponseWriter) {
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
