package test

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// PostShopsNew -
// ---------------------------------------------

type PostShopsNewHandlerFunc func(r PostShopsNewRequester) PostShopsNewResponder

func (f PostShopsNewHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestPostShopsNewParams{Request: r}).writePostShopsNewResponse(w)
}

type PostShopsNewRequester interface {
	Parse() (PostShopsNewRequest, error)
}

type requestPostShopsNewParams struct {
	Request *http.Request
}

func (r requestPostShopsNewParams) Parse() (PostShopsNewRequest, error) {
	return newPostShopsNewParams(r.Request)
}

type PostShopsNewRequest struct {
	HTTPRequest *http.Request

	Query struct {
		Page *int32
	}
}

func newPostShopsNewParams(r *http.Request) (zero PostShopsNewRequest, _ error) {
	var params PostShopsNewRequest
	params.HTTPRequest = r

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

type PostShopsNewResponder interface {
	writePostShopsNewResponse(w http.ResponseWriter)
}

func PostShopsNewResponse200() PostShopsNewResponder {
	var out postShopsNewResponse200
	return out
}

type postShopsNewResponse200 struct{}

func (r postShopsNewResponse200) writePostShopsNewResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func PostShopsNewResponseDefault(code int) PostShopsNewResponder {
	var out postShopsNewResponseDefault
	out.Code = code
	return out
}

type postShopsNewResponseDefault struct {
	Code int
}

func (r postShopsNewResponseDefault) writePostShopsNewResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsShop -
// ---------------------------------------------

type GetShopsShopHandlerFunc func(r GetShopsShopRequester) GetShopsShopResponder

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetShopsShopParams{Request: r}).writeGetShopsShopResponse(w)
}

type GetShopsShopRequester interface {
	Parse() (GetShopsShopRequest, error)
}

type requestGetShopsShopParams struct {
	Request *http.Request
}

func (r requestGetShopsShopParams) Parse() (GetShopsShopRequest, error) {
	return newGetShopsShopParams(r.Request)
}

type GetShopsShopRequest struct {
	HTTPRequest *http.Request

	Query struct {
		Page *int32
	}

	Path struct {
		Shop string
	}
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopRequest, _ error) {
	var params GetShopsShopRequest
	params.HTTPRequest = r

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

type GetShopsShopResponder interface {
	writeGetShopsShopResponse(w http.ResponseWriter)
}

func GetShopsShopResponse200() GetShopsShopResponder {
	var out getShopsShopResponse200
	return out
}

type getShopsShopResponse200 struct{}

func (r getShopsShopResponse200) writeGetShopsShopResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func GetShopsShopResponseDefault(code int) GetShopsShopResponder {
	var out getShopsShopResponseDefault
	out.Code = code
	return out
}

type getShopsShopResponseDefault struct {
	Code int
}

func (r getShopsShopResponseDefault) writeGetShopsShopResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsShopReviews -
// ---------------------------------------------

type GetShopsShopReviewsHandlerFunc func(r GetShopsShopReviewsRequester) GetShopsShopReviewsResponder

func (f GetShopsShopReviewsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetShopsShopReviewsParams{Request: r}).writeGetShopsShopReviewsResponse(w)
}

type GetShopsShopReviewsRequester interface {
	Parse() (GetShopsShopReviewsRequest, error)
}

type requestGetShopsShopReviewsParams struct {
	Request *http.Request
}

func (r requestGetShopsShopReviewsParams) Parse() (GetShopsShopReviewsRequest, error) {
	return newGetShopsShopReviewsParams(r.Request)
}

type GetShopsShopReviewsRequest struct {
	HTTPRequest *http.Request

	Query struct {
		Page *int32
	}

	Path struct {
		Shop string
	}
}

func newGetShopsShopReviewsParams(r *http.Request) (zero GetShopsShopReviewsRequest, _ error) {
	var params GetShopsShopReviewsRequest
	params.HTTPRequest = r

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

type GetShopsShopReviewsResponder interface {
	writeGetShopsShopReviewsResponse(w http.ResponseWriter)
}

func GetShopsShopReviewsResponse200() GetShopsShopReviewsResponder {
	var out getShopsShopReviewsResponse200
	return out
}

type getShopsShopReviewsResponse200 struct{}

func (r getShopsShopReviewsResponse200) writeGetShopsShopReviewsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func GetShopsShopReviewsResponseDefault(code int) GetShopsShopReviewsResponder {
	var out getShopsShopReviewsResponseDefault
	out.Code = code
	return out
}

type getShopsShopReviewsResponseDefault struct {
	Code int
}

func (r getShopsShopReviewsResponseDefault) writeGetShopsShopReviewsResponse(w http.ResponseWriter) {
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
