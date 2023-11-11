package test

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// Get -
// ---------------------------------------------

type GetHandlerFunc func(r GetRequestParser) GetResponse

func (f GetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetHTTPRequest(r)).Write(w)
}

type GetRequestParser interface {
	Parse() GetRequest
}

func GetHTTPRequest(r *http.Request) GetRequestParser {
	return getHTTPRequest{r}
}

type getHTTPRequest struct {
	Request *http.Request
}

func (r getHTTPRequest) Parse() GetRequest {
	return newGetParams(r.Request)
}

type GetRequest struct {
	HTTPRequest *http.Request
}

func newGetParams(r *http.Request) (zero GetRequest) {
	var params GetRequest
	params.HTTPRequest = r

	return params
}

func (r GetRequest) Parse() GetRequest { return r }

type GetResponse interface {
	get()
	Write(w http.ResponseWriter)
}

func NewGetResponseDefault(code int) GetResponse {
	var out GetResponseDefault
	out.Code = code
	return out
}

type GetResponseDefault struct {
	Code int
}

func (r GetResponseDefault) get() {}

func (r GetResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShops -
// ---------------------------------------------

type GetShopsHandlerFunc func(r GetShopsRequestParser) GetShopsResponse

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsHTTPRequest(r)).Write(w)
}

type GetShopsRequestParser interface {
	Parse() GetShopsRequest
}

func GetShopsHTTPRequest(r *http.Request) GetShopsRequestParser {
	return getShopsHTTPRequest{r}
}

type getShopsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsHTTPRequest) Parse() GetShopsRequest {
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

func (r GetShopsRequest) Parse() GetShopsRequest { return r }

type GetShopsResponse interface {
	getShops()
	Write(w http.ResponseWriter)
}

func NewGetShopsResponseDefault(code int) GetShopsResponse {
	var out GetShopsResponseDefault
	out.Code = code
	return out
}

type GetShopsResponseDefault struct {
	Code int
}

func (r GetShopsResponseDefault) getShops() {}

func (r GetShopsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsRT -
// ---------------------------------------------

type GetShopsRTHandlerFunc func(r GetShopsRTRequestParser) GetShopsRTResponse

func (f GetShopsRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsRTHTTPRequest(r)).Write(w)
}

type GetShopsRTRequestParser interface {
	Parse() GetShopsRTRequest
}

func GetShopsRTHTTPRequest(r *http.Request) GetShopsRTRequestParser {
	return getShopsRTHTTPRequest{r}
}

type getShopsRTHTTPRequest struct {
	Request *http.Request
}

func (r getShopsRTHTTPRequest) Parse() GetShopsRTRequest {
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

func (r GetShopsRTRequest) Parse() GetShopsRTRequest { return r }

type GetShopsRTResponse interface {
	getShopsRT()
	Write(w http.ResponseWriter)
}

func NewGetShopsRTResponseDefault(code int) GetShopsRTResponse {
	var out GetShopsRTResponseDefault
	out.Code = code
	return out
}

type GetShopsRTResponseDefault struct {
	Code int
}

func (r GetShopsRTResponseDefault) getShopsRT() {}

func (r GetShopsRTResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsActivate -
// ---------------------------------------------

type GetShopsActivateHandlerFunc func(r GetShopsActivateRequestParser) GetShopsActivateResponse

func (f GetShopsActivateHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsActivateHTTPRequest(r)).Write(w)
}

type GetShopsActivateRequestParser interface {
	Parse() GetShopsActivateRequest
}

func GetShopsActivateHTTPRequest(r *http.Request) GetShopsActivateRequestParser {
	return getShopsActivateHTTPRequest{r}
}

type getShopsActivateHTTPRequest struct {
	Request *http.Request
}

func (r getShopsActivateHTTPRequest) Parse() GetShopsActivateRequest {
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

func (r GetShopsActivateRequest) Parse() GetShopsActivateRequest { return r }

type GetShopsActivateResponse interface {
	getShopsActivate()
	Write(w http.ResponseWriter)
}

func NewGetShopsActivateResponseDefault(code int) GetShopsActivateResponse {
	var out GetShopsActivateResponseDefault
	out.Code = code
	return out
}

type GetShopsActivateResponseDefault struct {
	Code int
}

func (r GetShopsActivateResponseDefault) getShopsActivate() {}

func (r GetShopsActivateResponseDefault) Write(w http.ResponseWriter) {
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

func (r GetShopsShopRequest) Parse() (GetShopsShopRequest, error) { return r, nil }

type GetShopsShopResponse interface {
	getShopsShop()
	Write(w http.ResponseWriter)
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
// GetShopsShopRT -
// ---------------------------------------------

type GetShopsShopRTHandlerFunc func(r GetShopsShopRTRequestParser) GetShopsShopRTResponse

func (f GetShopsShopRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsShopRTHTTPRequest(r)).Write(w)
}

type GetShopsShopRTRequestParser interface {
	Parse() (GetShopsShopRTRequest, error)
}

func GetShopsShopRTHTTPRequest(r *http.Request) GetShopsShopRTRequestParser {
	return getShopsShopRTHTTPRequest{r}
}

type getShopsShopRTHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopRTHTTPRequest) Parse() (GetShopsShopRTRequest, error) {
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

func (r GetShopsShopRTRequest) Parse() (GetShopsShopRTRequest, error) { return r, nil }

type GetShopsShopRTResponse interface {
	getShopsShopRT()
	Write(w http.ResponseWriter)
}

func NewGetShopsShopRTResponseDefault(code int) GetShopsShopRTResponse {
	var out GetShopsShopRTResponseDefault
	out.Code = code
	return out
}

type GetShopsShopRTResponseDefault struct {
	Code int
}

func (r GetShopsShopRTResponseDefault) getShopsShopRT() {}

func (r GetShopsShopRTResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsShopPets -
// ---------------------------------------------

type GetShopsShopPetsHandlerFunc func(r GetShopsShopPetsRequestParser) GetShopsShopPetsResponse

func (f GetShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsShopPetsHTTPRequest(r)).Write(w)
}

type GetShopsShopPetsRequestParser interface {
	Parse() (GetShopsShopPetsRequest, error)
}

func GetShopsShopPetsHTTPRequest(r *http.Request) GetShopsShopPetsRequestParser {
	return getShopsShopPetsHTTPRequest{r}
}

type getShopsShopPetsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopPetsHTTPRequest) Parse() (GetShopsShopPetsRequest, error) {
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

func (r GetShopsShopPetsRequest) Parse() (GetShopsShopPetsRequest, error) { return r, nil }

type GetShopsShopPetsResponse interface {
	getShopsShopPets()
	Write(w http.ResponseWriter)
}

func NewGetShopsShopPetsResponseDefault(code int) GetShopsShopPetsResponse {
	var out GetShopsShopPetsResponseDefault
	out.Code = code
	return out
}

type GetShopsShopPetsResponseDefault struct {
	Code int
}

func (r GetShopsShopPetsResponseDefault) getShopsShopPets() {}

func (r GetShopsShopPetsResponseDefault) Write(w http.ResponseWriter) {
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
