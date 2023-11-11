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

type PostShopsNewHandlerFunc func(r PostShopsNewRequestParser) PostShopsNewResponse

func (f PostShopsNewHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(PostShopsNewHTTPRequest(r)).Write(w)
}

type PostShopsNewRequestParser interface {
	Parse() (PostShopsNewRequest, error)
}

func PostShopsNewHTTPRequest(r *http.Request) PostShopsNewRequestParser {
	return postShopsNewHTTPRequest{r}
}

type postShopsNewHTTPRequest struct {
	Request *http.Request
}

func (r postShopsNewHTTPRequest) Parse() (PostShopsNewRequest, error) {
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

func (r PostShopsNewRequest) Parse() (PostShopsNewRequest, error) { return r, nil }

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

type GetShopsShopHandlerFunc func(r GetShopsShopRequestParser) GetShopsShopResponse

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsShopHTTPRequest(r)).Write(w)
}

type GetShopsShopRequestParser interface {
	Parse() (GetShopsShopRequest, error)
}

func GetShopsShopHTTPRequest(r *http.Request) GetShopsShopRequestParser {
	return getShopsShopHTTPRequest{r}
}

type getShopsShopHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopHTTPRequest) Parse() (GetShopsShopRequest, error) {
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

func (r GetShopsShopRequest) Parse() (GetShopsShopRequest, error) { return r, nil }

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

type GetShopsShopReviewsHandlerFunc func(r GetShopsShopReviewsRequestParser) GetShopsShopReviewsResponse

func (f GetShopsShopReviewsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsShopReviewsHTTPRequest(r)).Write(w)
}

type GetShopsShopReviewsRequestParser interface {
	Parse() (GetShopsShopReviewsRequest, error)
}

func GetShopsShopReviewsHTTPRequest(r *http.Request) GetShopsShopReviewsRequestParser {
	return getShopsShopReviewsHTTPRequest{r}
}

type getShopsShopReviewsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopReviewsHTTPRequest) Parse() (GetShopsShopReviewsRequest, error) {
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

func (r GetShopsShopReviewsRequest) Parse() (GetShopsShopReviewsRequest, error) { return r, nil }

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
