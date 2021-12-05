package test

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// GetShopsShop -
// ---------------------------------------------

type GetShopsShopHandlerFunc func(r GetShopsShopRequester) GetShopsShopResponder

func (f GetShopsShopHandlerFunc) Handle(r GetShopsShopRequester) GetShopsShopResponder {
	return f(r)
}

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopParams{Request: r}).writeGetShopsShopResponse(w)
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

	Page      *int32
	Shop      string
	RequestID *string
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
					return zero, ErrParseQueryParam{Name: "page", Err: fmt.Errorf("parse int32: %w", err)}
				}
				v := int32(vInt)
				params.Page = &v
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("request-id")
			if len(hs) > 0 {
				v := hs[0]
				params.RequestID = &v
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
				return zero, ErrParsePathParam{Name: "shop", Err: fmt.Errorf("is required")}
			}

			v := vPath
			params.Shop = v
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

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

type ErrParseQueryParam struct {
	Name string
	Err  error
}

func (e ErrParseQueryParam) Error() string {
	return fmt.Sprintf("query parameter '%s': %e", e.Name, e.Err)
}

type ErrParsePathParam struct {
	Name string
	Err  error
}

func (e ErrParsePathParam) Error() string {
	return fmt.Sprintf("path parameter '%s': %e", e.Name, e.Err)
}

type ErrParseHeaderParam struct {
	Name string
	Err  error
}

func (e ErrParseHeaderParam) Error() string {
	return fmt.Sprintf("header parameter '%s': %e", e.Name, e.Err)
}
