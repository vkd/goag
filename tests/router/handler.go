package test

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// GetRT -
// ---------------------------------------------

type GetRTHandlerFunc func(r GetRTRequester) GetRTResponder

func (f GetRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetRTParams{Request: r}).writeGetRTResponse(w)
}

type GetRTRequester interface {
	Parse() GetRTRequest
}

type requestGetRTParams struct {
	Request *http.Request
}

func (r requestGetRTParams) Parse() GetRTRequest {
	return newGetRTParams(r.Request)
}

type GetRTRequest struct {
	HTTPRequest *http.Request
}

func newGetRTParams(r *http.Request) (zero GetRTRequest) {
	var params GetRTRequest
	params.HTTPRequest = r

	return params
}

type GetRTResponder interface {
	writeGetRTResponse(w http.ResponseWriter)
}

func GetRTResponseDefault(code int) GetRTResponder {
	var out getRTResponseDefault
	out.Code = code
	return out
}

type getRTResponseDefault struct {
	Code int
}

func (r getRTResponseDefault) writeGetRTResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShops -
// ---------------------------------------------

type GetShopsHandlerFunc func(r GetShopsRequester) GetShopsResponder

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetShopsParams{Request: r}).writeGetShopsResponse(w)
}

type GetShopsRequester interface {
	Parse() GetShopsRequest
}

type requestGetShopsParams struct {
	Request *http.Request
}

func (r requestGetShopsParams) Parse() GetShopsRequest {
	return newGetShopsParams(r.Request)
}

type GetShopsRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsParams(r *http.Request) (zero GetShopsRequest) {
	var params GetShopsRequest
	params.HTTPRequest = r

	return params
}

type GetShopsResponder interface {
	writeGetShopsResponse(w http.ResponseWriter)
}

func GetShopsResponseDefault(code int) GetShopsResponder {
	var out getShopsResponseDefault
	out.Code = code
	return out
}

type getShopsResponseDefault struct {
	Code int
}

func (r getShopsResponseDefault) writeGetShopsResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsRT -
// ---------------------------------------------

type GetShopsRTHandlerFunc func(r GetShopsRTRequester) GetShopsRTResponder

func (f GetShopsRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetShopsRTParams{Request: r}).writeGetShopsRTResponse(w)
}

type GetShopsRTRequester interface {
	Parse() GetShopsRTRequest
}

type requestGetShopsRTParams struct {
	Request *http.Request
}

func (r requestGetShopsRTParams) Parse() GetShopsRTRequest {
	return newGetShopsRTParams(r.Request)
}

type GetShopsRTRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsRTParams(r *http.Request) (zero GetShopsRTRequest) {
	var params GetShopsRTRequest
	params.HTTPRequest = r

	return params
}

type GetShopsRTResponder interface {
	writeGetShopsRTResponse(w http.ResponseWriter)
}

func GetShopsRTResponseDefault(code int) GetShopsRTResponder {
	var out getShopsRTResponseDefault
	out.Code = code
	return out
}

type getShopsRTResponseDefault struct {
	Code int
}

func (r getShopsRTResponseDefault) writeGetShopsRTResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsActivate -
// ---------------------------------------------

type GetShopsActivateHandlerFunc func(r GetShopsActivateRequester) GetShopsActivateResponder

func (f GetShopsActivateHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetShopsActivateParams{Request: r}).writeGetShopsActivateResponse(w)
}

type GetShopsActivateRequester interface {
	Parse() GetShopsActivateRequest
}

type requestGetShopsActivateParams struct {
	Request *http.Request
}

func (r requestGetShopsActivateParams) Parse() GetShopsActivateRequest {
	return newGetShopsActivateParams(r.Request)
}

type GetShopsActivateRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsActivateParams(r *http.Request) (zero GetShopsActivateRequest) {
	var params GetShopsActivateRequest
	params.HTTPRequest = r

	return params
}

type GetShopsActivateResponder interface {
	writeGetShopsActivateResponse(w http.ResponseWriter)
}

func GetShopsActivateResponseDefault(code int) GetShopsActivateResponder {
	var out getShopsActivateResponseDefault
	out.Code = code
	return out
}

type getShopsActivateResponseDefault struct {
	Code int
}

func (r getShopsActivateResponseDefault) writeGetShopsActivateResponse(w http.ResponseWriter) {
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

	Path struct {
		Shop string
	}
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopRequest, _ error) {
	var params GetShopsShopRequest
	params.HTTPRequest = r

	// Path parameters
	{
		p := r.URL.Path
		if !strings.HasPrefix(p, "/api/v1") {
			return zero, fmt.Errorf("wrong path: expected '/api/v1...'")
		}
		p = p[7:] // "/api/v1"

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '/api/v1/...'")
		}

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
// GetShopsShopRT -
// ---------------------------------------------

type GetShopsShopRTHandlerFunc func(r GetShopsShopRTRequester) GetShopsShopRTResponder

func (f GetShopsShopRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetShopsShopRTParams{Request: r}).writeGetShopsShopRTResponse(w)
}

type GetShopsShopRTRequester interface {
	Parse() (GetShopsShopRTRequest, error)
}

type requestGetShopsShopRTParams struct {
	Request *http.Request
}

func (r requestGetShopsShopRTParams) Parse() (GetShopsShopRTRequest, error) {
	return newGetShopsShopRTParams(r.Request)
}

type GetShopsShopRTRequest struct {
	HTTPRequest *http.Request

	Path struct {
		Shop string
	}
}

func newGetShopsShopRTParams(r *http.Request) (zero GetShopsShopRTRequest, _ error) {
	var params GetShopsShopRTRequest
	params.HTTPRequest = r

	// Path parameters
	{
		p := r.URL.Path
		if !strings.HasPrefix(p, "/api/v1") {
			return zero, fmt.Errorf("wrong path: expected '/api/v1...'")
		}
		p = p[7:] // "/api/v1"

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '/api/v1/...'")
		}

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/'")
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

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/'")
		}
		p = p[1:] // "/"
	}

	return params, nil
}

type GetShopsShopRTResponder interface {
	writeGetShopsShopRTResponse(w http.ResponseWriter)
}

func GetShopsShopRTResponseDefault(code int) GetShopsShopRTResponder {
	var out getShopsShopRTResponseDefault
	out.Code = code
	return out
}

type getShopsShopRTResponseDefault struct {
	Code int
}

func (r getShopsShopRTResponseDefault) writeGetShopsShopRTResponse(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsShopPets -
// ---------------------------------------------

type GetShopsShopPetsHandlerFunc func(r GetShopsShopPetsRequester) GetShopsShopPetsResponder

func (f GetShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetShopsShopPetsParams{Request: r}).writeGetShopsShopPetsResponse(w)
}

type GetShopsShopPetsRequester interface {
	Parse() (GetShopsShopPetsRequest, error)
}

type requestGetShopsShopPetsParams struct {
	Request *http.Request
}

func (r requestGetShopsShopPetsParams) Parse() (GetShopsShopPetsRequest, error) {
	return newGetShopsShopPetsParams(r.Request)
}

type GetShopsShopPetsRequest struct {
	HTTPRequest *http.Request

	Path struct {
		Shop string
	}
}

func newGetShopsShopPetsParams(r *http.Request) (zero GetShopsShopPetsRequest, _ error) {
	var params GetShopsShopPetsRequest
	params.HTTPRequest = r

	// Path parameters
	{
		p := r.URL.Path
		if !strings.HasPrefix(p, "/api/v1") {
			return zero, fmt.Errorf("wrong path: expected '/api/v1...'")
		}
		p = p[7:] // "/api/v1"

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '/api/v1/...'")
		}

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets'")
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

		if !strings.HasPrefix(p, "/pets") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets'")
		}
		p = p[5:] // "/pets"
	}

	return params, nil
}

type GetShopsShopPetsResponder interface {
	writeGetShopsShopPetsResponse(w http.ResponseWriter)
}

func GetShopsShopPetsResponseDefault(code int) GetShopsShopPetsResponder {
	var out getShopsShopPetsResponseDefault
	out.Code = code
	return out
}

type getShopsShopPetsResponseDefault struct {
	Code int
}

func (r getShopsShopPetsResponseDefault) writeGetShopsShopPetsResponse(w http.ResponseWriter) {
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
