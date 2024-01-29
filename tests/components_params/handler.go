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
// PostShopsNew -
// ---------------------------------------------

type PostShopsNewHandlerFunc func(ctx context.Context, r PostShopsNewRequest) PostShopsNewResponse

func (f PostShopsNewHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsNewHTTPRequest(r)).Write(w)
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
	Query struct {
		Page *int32
	}
}

func newPostShopsNewParams(r *http.Request) (zero PostShopsNewParams, _ error) {
	var params PostShopsNewParams

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
				params.Query.Page = &v
			}
		}
	}

	return params, nil
}

func (r PostShopsNewParams) HTTP() *http.Request { return nil }

func (r PostShopsNewParams) Parse() (PostShopsNewParams, error) { return r, nil }

type PostShopsNewResponse interface {
	postShopsNew()
	Write(w http.ResponseWriter)
}

func NewPostShopsNewResponse200() PostShopsNewResponse {
	var out PostShopsNewResponse200
	return out
}

type PostShopsNewResponse200 struct{}

func (r PostShopsNewResponse200) postShopsNew() {}

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

func (r PostShopsNewResponseDefault) postShopsNew() {}

func (r PostShopsNewResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

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
		Page *int32
	}

	Path struct {
		Shop string
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
				params.Query.Page = &v
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

			v := vPath
			params.Path.Shop = v
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

// ---------------------------------------------
// GetShopsShopReviews -
// ---------------------------------------------

type GetShopsShopReviewsHandlerFunc func(ctx context.Context, r GetShopsShopReviewsRequest) GetShopsShopReviewsResponse

func (f GetShopsShopReviewsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopReviewsHTTPRequest(r)).Write(w)
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
	Query struct {
		Page *int32
	}

	Path struct {
		Shop string
	}
}

func newGetShopsShopReviewsParams(r *http.Request) (zero GetShopsShopReviewsParams, _ error) {
	var params GetShopsShopReviewsParams

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
				params.Query.Page = &v
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

			v := vPath
			params.Path.Shop = v
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
	getShopsShopReviews()
	Write(w http.ResponseWriter)
}

func NewGetShopsShopReviewsResponse200() GetShopsShopReviewsResponse {
	var out GetShopsShopReviewsResponse200
	return out
}

type GetShopsShopReviewsResponse200 struct{}

func (r GetShopsShopReviewsResponse200) getShopsShopReviews() {}

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

func (r GetShopsShopReviewsResponseDefault) getShopsShopReviews() {}

func (r GetShopsShopReviewsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
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
