package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// PostShopsNew -
// ---------------------------------------------

type PostShopsNewHandlerFunc func(ctx context.Context, r PostShopsNewRequest) PostShopsNewResponse

func (f PostShopsNewHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsNewHTTPRequest(r)).writePostShopsNew(w)
}

type PostShopsNewRequest interface {
	HTTP() *http.Request
	Parse() (PostShopsNewParams, error)
}

func PostShopsNewHTTPRequest(r *http.Request) PostShopsNewRequest {
	return postShopsNewHTTPRequest{r}
}

type postShopsNewHTTPRequest struct {
	Request *http.Request
}

func (r postShopsNewHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postShopsNewHTTPRequest) Parse() (PostShopsNewParams, error) {
	return newPostShopsNewParams(r.Request)
}

type PostShopsNewParams struct {
	Query PostShopsNewParamsQuery
}

type PostShopsNewParamsQuery struct {
	Page Maybe[int32]
}

func newPostShopsNewParams(r *http.Request) (zero PostShopsNewParams, _ error) {
	var params PostShopsNewParams

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
				vOpt := int32(vInt64)
				params.Query.Page.Set(vOpt)
			}
		}
	}

	return params, nil
}

func (r PostShopsNewParams) HTTP() *http.Request { return nil }

func (r PostShopsNewParams) Parse() (PostShopsNewParams, error) { return r, nil }

type PostShopsNewResponse interface {
	writePostShopsNew(http.ResponseWriter)
}

func NewPostShopsNewResponse200() PostShopsNewResponse {
	var out PostShopsNewResponse200
	return out
}

type PostShopsNewResponse200 struct{}

func (r PostShopsNewResponse200) writePostShopsNew(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsNewResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewPostShopsNewResponseDefault(code int) PostShopsNewResponse {
	var out PostShopsNewResponseDefault
	out.Code = code
	return out
}

type PostShopsNewResponseDefault struct {
	Code int
}

func (r PostShopsNewResponseDefault) writePostShopsNew(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsNewResponseDefault) Write(w http.ResponseWriter) {
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
	Query GetShopsShopParamsQuery

	Path GetShopsShopParamsPath
}

type GetShopsShopParamsQuery struct {
	Page Maybe[int32]
}

type GetShopsShopParamsPath struct {
	Shop string
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
				vOpt := int32(vInt64)
				params.Query.Page.Set(vOpt)
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
// GetShopsShopReviews -
// ---------------------------------------------

type GetShopsShopReviewsHandlerFunc func(ctx context.Context, r GetShopsShopReviewsRequest) GetShopsShopReviewsResponse

func (f GetShopsShopReviewsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopReviewsHTTPRequest(r)).writeGetShopsShopReviews(w)
}

type GetShopsShopReviewsRequest interface {
	HTTP() *http.Request
	Parse() (GetShopsShopReviewsParams, error)
}

func GetShopsShopReviewsHTTPRequest(r *http.Request) GetShopsShopReviewsRequest {
	return getShopsShopReviewsHTTPRequest{r}
}

type getShopsShopReviewsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopReviewsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsShopReviewsHTTPRequest) Parse() (GetShopsShopReviewsParams, error) {
	return newGetShopsShopReviewsParams(r.Request)
}

type GetShopsShopReviewsParams struct {
	Query GetShopsShopReviewsParamsQuery

	Path GetShopsShopReviewsParamsPath
}

type GetShopsShopReviewsParamsQuery struct {
	Page Maybe[int32]
}

type GetShopsShopReviewsParamsPath struct {
	Shop string
}

func newGetShopsShopReviewsParams(r *http.Request) (zero GetShopsShopReviewsParams, _ error) {
	var params GetShopsShopReviewsParams

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
				vOpt := int32(vInt64)
				params.Query.Page.Set(vOpt)
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/reviews'")
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

		if !strings.HasPrefix(p, "/reviews") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/reviews'")
		}
		p = p[8:] // "/reviews"
	}

	return params, nil
}

func (r GetShopsShopReviewsParams) HTTP() *http.Request { return nil }

func (r GetShopsShopReviewsParams) Parse() (GetShopsShopReviewsParams, error) { return r, nil }

type GetShopsShopReviewsResponse interface {
	writeGetShopsShopReviews(http.ResponseWriter)
}

func NewGetShopsShopReviewsResponse200() GetShopsShopReviewsResponse {
	var out GetShopsShopReviewsResponse200
	return out
}

type GetShopsShopReviewsResponse200 struct{}

func (r GetShopsShopReviewsResponse200) writeGetShopsShopReviews(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopReviewsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func NewGetShopsShopReviewsResponseDefault(code int) GetShopsShopReviewsResponse {
	var out GetShopsShopReviewsResponseDefault
	out.Code = code
	return out
}

type GetShopsShopReviewsResponseDefault struct {
	Code int
}

func (r GetShopsShopReviewsResponseDefault) writeGetShopsShopReviews(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopReviewsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

func write(w io.Writer, r io.Reader, name string) {
	_, err := io.Copy(w, r)
	if err != nil {
		LogError(fmt.Errorf("write response %q: %w", name, err))
	}
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
