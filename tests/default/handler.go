package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// Get -
// ---------------------------------------------

type GetHandlerFunc func(r GetRequest) GetResponse

func (f GetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetHTTPRequest(r)).Write(w)
}

type GetRequest interface {
	HTTP() *http.Request
	Parse() GetParams
}

func GetHTTPRequest(r *http.Request) GetRequest {
	return getHTTPRequest{r}
}

type getHTTPRequest struct {
	Request *http.Request
}

func (r getHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getHTTPRequest) Parse() GetParams {
	return newGetParams(r.Request)
}

type GetParams struct {
}

func newGetParams(r *http.Request) (zero GetParams) {
	var params GetParams

	return params
}

func (r GetParams) HTTP() *http.Request { return nil }

func (r GetParams) Parse() GetParams { return r }

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

type GetShopsHandlerFunc func(r GetShopsRequest) GetShopsResponse

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsHTTPRequest(r)).Write(w)
}

type GetShopsRequest interface {
	HTTP() *http.Request
	Parse() GetShopsParams
}

func GetShopsHTTPRequest(r *http.Request) GetShopsRequest {
	return getShopsHTTPRequest{r}
}

type getShopsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsHTTPRequest) Parse() GetShopsParams {
	return newGetShopsParams(r.Request)
}

type GetShopsParams struct {
}

func newGetShopsParams(r *http.Request) (zero GetShopsParams) {
	var params GetShopsParams

	return params
}

func (r GetShopsParams) HTTP() *http.Request { return nil }

func (r GetShopsParams) Parse() GetShopsParams { return r }

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

type GetShopsRTHandlerFunc func(r GetShopsRTRequest) GetShopsRTResponse

func (f GetShopsRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsRTHTTPRequest(r)).Write(w)
}

type GetShopsRTRequest interface {
	HTTP() *http.Request
	Parse() GetShopsRTParams
}

func GetShopsRTHTTPRequest(r *http.Request) GetShopsRTRequest {
	return getShopsRTHTTPRequest{r}
}

type getShopsRTHTTPRequest struct {
	Request *http.Request
}

func (r getShopsRTHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsRTHTTPRequest) Parse() GetShopsRTParams {
	return newGetShopsRTParams(r.Request)
}

type GetShopsRTParams struct {
}

func newGetShopsRTParams(r *http.Request) (zero GetShopsRTParams) {
	var params GetShopsRTParams

	return params
}

func (r GetShopsRTParams) HTTP() *http.Request { return nil }

func (r GetShopsRTParams) Parse() GetShopsRTParams { return r }

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

type GetShopsActivateHandlerFunc func(r GetShopsActivateRequest) GetShopsActivateResponse

func (f GetShopsActivateHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsActivateHTTPRequest(r)).Write(w)
}

type GetShopsActivateRequest interface {
	HTTP() *http.Request
	Parse() GetShopsActivateParams
}

func GetShopsActivateHTTPRequest(r *http.Request) GetShopsActivateRequest {
	return getShopsActivateHTTPRequest{r}
}

type getShopsActivateHTTPRequest struct {
	Request *http.Request
}

func (r getShopsActivateHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsActivateHTTPRequest) Parse() GetShopsActivateParams {
	return newGetShopsActivateParams(r.Request)
}

type GetShopsActivateParams struct {
}

func newGetShopsActivateParams(r *http.Request) (zero GetShopsActivateParams) {
	var params GetShopsActivateParams

	return params
}

func (r GetShopsActivateParams) HTTP() *http.Request { return nil }

func (r GetShopsActivateParams) Parse() GetShopsActivateParams { return r }

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
// GetShopsActivateRT -
// ---------------------------------------------

type GetShopsActivateRTHandlerFunc func(r GetShopsActivateRTRequest) GetShopsActivateRTResponse

func (f GetShopsActivateRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsActivateRTHTTPRequest(r)).Write(w)
}

type GetShopsActivateRTRequest interface {
	HTTP() *http.Request
	Parse() GetShopsActivateRTParams
}

func GetShopsActivateRTHTTPRequest(r *http.Request) GetShopsActivateRTRequest {
	return getShopsActivateRTHTTPRequest{r}
}

type getShopsActivateRTHTTPRequest struct {
	Request *http.Request
}

func (r getShopsActivateRTHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsActivateRTHTTPRequest) Parse() GetShopsActivateRTParams {
	return newGetShopsActivateRTParams(r.Request)
}

type GetShopsActivateRTParams struct {
}

func newGetShopsActivateRTParams(r *http.Request) (zero GetShopsActivateRTParams) {
	var params GetShopsActivateRTParams

	return params
}

func (r GetShopsActivateRTParams) HTTP() *http.Request { return nil }

func (r GetShopsActivateRTParams) Parse() GetShopsActivateRTParams { return r }

type GetShopsActivateRTResponse interface {
	getShopsActivateRT()
	Write(w http.ResponseWriter)
}

func NewGetShopsActivateRTResponseDefault(code int) GetShopsActivateRTResponse {
	var out GetShopsActivateRTResponseDefault
	out.Code = code
	return out
}

type GetShopsActivateRTResponseDefault struct {
	Code int
}

func (r GetShopsActivateRTResponseDefault) getShopsActivateRT() {}

func (r GetShopsActivateRTResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsActivateTag -
// ---------------------------------------------

type GetShopsActivateTagHandlerFunc func(r GetShopsActivateTagRequest) GetShopsActivateTagResponse

func (f GetShopsActivateTagHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsActivateTagHTTPRequest(r)).Write(w)
}

type GetShopsActivateTagRequest interface {
	HTTP() *http.Request
	Parse() GetShopsActivateTagParams
}

func GetShopsActivateTagHTTPRequest(r *http.Request) GetShopsActivateTagRequest {
	return getShopsActivateTagHTTPRequest{r}
}

type getShopsActivateTagHTTPRequest struct {
	Request *http.Request
}

func (r getShopsActivateTagHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsActivateTagHTTPRequest) Parse() GetShopsActivateTagParams {
	return newGetShopsActivateTagParams(r.Request)
}

type GetShopsActivateTagParams struct {
}

func newGetShopsActivateTagParams(r *http.Request) (zero GetShopsActivateTagParams) {
	var params GetShopsActivateTagParams

	return params
}

func (r GetShopsActivateTagParams) HTTP() *http.Request { return nil }

func (r GetShopsActivateTagParams) Parse() GetShopsActivateTagParams { return r }

type GetShopsActivateTagResponse interface {
	getShopsActivateTag()
	Write(w http.ResponseWriter)
}

func NewGetShopsActivateTagResponseDefault(code int) GetShopsActivateTagResponse {
	var out GetShopsActivateTagResponseDefault
	out.Code = code
	return out
}

type GetShopsActivateTagResponseDefault struct {
	Code int
}

func (r GetShopsActivateTagResponseDefault) getShopsActivateTag() {}

func (r GetShopsActivateTagResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsShop -
// ---------------------------------------------

type GetShopsShopHandlerFunc func(r GetShopsShopRequest) GetShopsShopResponse

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsShopHTTPRequest(r)).Write(w)
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
	Path struct {
		Shop int32
	}
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopParams, _ error) {
	var params GetShopsShopParams

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

			vInt, err := strconv.ParseInt(vPath, 10, 32)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse int32", Err: err}
			}
			v := int32(vInt)
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

type GetShopsShopRTHandlerFunc func(r GetShopsShopRTRequest) GetShopsShopRTResponse

func (f GetShopsShopRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsShopRTHTTPRequest(r)).Write(w)
}

type GetShopsShopRTRequest interface {
	HTTP() *http.Request
	Parse() (GetShopsShopRTParams, error)
}

func GetShopsShopRTHTTPRequest(r *http.Request) GetShopsShopRTRequest {
	return getShopsShopRTHTTPRequest{r}
}

type getShopsShopRTHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopRTHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsShopRTHTTPRequest) Parse() (GetShopsShopRTParams, error) {
	return newGetShopsShopRTParams(r.Request)
}

type GetShopsShopRTParams struct {
	Path struct {
		Shop int32
	}
}

func newGetShopsShopRTParams(r *http.Request) (zero GetShopsShopRTParams, _ error) {
	var params GetShopsShopRTParams

	// Path parameters
	{
		p := r.URL.Path

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

			vInt, err := strconv.ParseInt(vPath, 10, 32)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse int32", Err: err}
			}
			v := int32(vInt)
			params.Path.Shop = v
		}

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/'")
		}
		p = p[1:] // "/"
	}

	return params, nil
}

func (r GetShopsShopRTParams) HTTP() *http.Request { return nil }

func (r GetShopsShopRTParams) Parse() (GetShopsShopRTParams, error) { return r, nil }

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

type GetShopsShopPetsHandlerFunc func(r GetShopsShopPetsRequest) GetShopsShopPetsResponse

func (f GetShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsShopPetsHTTPRequest(r)).Write(w)
}

type GetShopsShopPetsRequest interface {
	HTTP() *http.Request
	Parse() (GetShopsShopPetsParams, error)
}

func GetShopsShopPetsHTTPRequest(r *http.Request) GetShopsShopPetsRequest {
	return getShopsShopPetsHTTPRequest{r}
}

type getShopsShopPetsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsShopPetsHTTPRequest) Parse() (GetShopsShopPetsParams, error) {
	return newGetShopsShopPetsParams(r.Request)
}

type GetShopsShopPetsParams struct {
	Query struct {
		Page *int32

		PageSize int32
	}

	Path struct {
		Shop int32
	}
}

func newGetShopsShopPetsParams(r *http.Request) (zero GetShopsShopPetsParams, _ error) {
	var params GetShopsShopPetsParams

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
		{
			q, ok := query["page_size"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'page_size': is required")
			}
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_size", Reason: "parse int32", Err: err}
				}
				v := int32(vInt)
				params.Query.PageSize = v
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

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

			vInt, err := strconv.ParseInt(vPath, 10, 32)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse int32", Err: err}
			}
			v := int32(vInt)
			params.Path.Shop = v
		}

		if !strings.HasPrefix(p, "/pets") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets'")
		}
		p = p[5:] // "/pets"
	}

	return params, nil
}

func (r GetShopsShopPetsParams) HTTP() *http.Request { return nil }

func (r GetShopsShopPetsParams) Parse() (GetShopsShopPetsParams, error) { return r, nil }

type GetShopsShopPetsResponse interface {
	getShopsShopPets()
	Write(w http.ResponseWriter)
}

func NewGetShopsShopPetsResponse200JSON(body GetShopsShopPetsResponse200JSONBody, xNext string) GetShopsShopPetsResponse {
	var out GetShopsShopPetsResponse200JSON
	out.Body = body
	out.Headers.XNext = xNext
	return out
}

type GetShopsShopPetsResponse200JSONBody struct {
	Groups map[string]Pets `json:"groups"`
}

type GetShopsShopPetsResponse200JSON struct {
	Body    GetShopsShopPetsResponse200JSONBody
	Headers struct {
		Body  GetShopsShopPetsResponse200JSONBody
		XNext string
	}
}

func (r GetShopsShopPetsResponse200JSON) getShopsShopPets() {}

func (r GetShopsShopPetsResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("x-next", r.Headers.XNext)
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetShopsShopPetsResponse200JSON")
}

func NewGetShopsShopPetsResponseDefaultJSON(code int, body Error) GetShopsShopPetsResponse {
	var out GetShopsShopPetsResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

type GetShopsShopPetsResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r GetShopsShopPetsResponseDefaultJSON) getShopsShopPets() {}

func (r GetShopsShopPetsResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "GetShopsShopPetsResponseDefaultJSON")
}

// ---------------------------------------------
// ReviewShop -
// ---------------------------------------------

type ReviewShopHandlerFunc func(r ReviewShopRequest) ReviewShopResponse

func (f ReviewShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(ReviewShopHTTPRequest(r)).Write(w)
}

type ReviewShopRequest interface {
	HTTP() *http.Request
	Parse() (ReviewShopParams, error)
}

func ReviewShopHTTPRequest(r *http.Request) ReviewShopRequest {
	return reviewShopHTTPRequest{r}
}

type reviewShopHTTPRequest struct {
	Request *http.Request
}

func (r reviewShopHTTPRequest) HTTP() *http.Request { return r.Request }

func (r reviewShopHTTPRequest) Parse() (ReviewShopParams, error) {
	return newReviewShopParams(r.Request)
}

type ReviewShopParams struct {
	Query struct {
		Page *int32

		PageSize int32

		Tag []string

		Filter []int32
	}

	Path struct {
		Shop int32
	}

	Headers struct {
		RequestID *string

		UserID string
	}

	Body NewPet
}

func newReviewShopParams(r *http.Request) (zero ReviewShopParams, _ error) {
	var params ReviewShopParams

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
		{
			q, ok := query["page_size"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'page_size': is required")
			}
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_size", Reason: "parse int32", Err: err}
				}
				v := int32(vInt)
				params.Query.PageSize = v
			}
		}
		{
			q, ok := query["tag"]
			if ok && len(q) > 0 {
				params.Query.Tag = q
			}
		}
		{
			q, ok := query["filter"]
			if ok && len(q) > 0 {
				params.Query.Filter = make([]int32, len(q))
				for i := range q {
					vInt, err := strconv.ParseInt(q[i], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "filter", Reason: "parse int32", Err: err}
					}
					v1 := int32(vInt)
					params.Query.Filter[i] = v1
				}
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
				params.Headers.RequestID = &v
			}
		}
		{
			hs := header.Values("user-id")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'user-id': is required")
			}
			if len(hs) > 0 {
				v := hs[0]
				params.Headers.UserID = v
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/review'")
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

			vInt, err := strconv.ParseInt(vPath, 10, 32)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse int32", Err: err}
			}
			v := int32(vInt)
			params.Path.Shop = v
		}

		if !strings.HasPrefix(p, "/review") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/review'")
		}
		p = p[7:] // "/review"
	}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r ReviewShopParams) HTTP() *http.Request { return nil }

func (r ReviewShopParams) Parse() (ReviewShopParams, error) { return r, nil }

type ReviewShopResponse interface {
	reviewShop()
	Write(w http.ResponseWriter)
}

func NewReviewShopResponse200JSON(body Pet, xNext string) ReviewShopResponse {
	var out ReviewShopResponse200JSON
	out.Body = body
	out.Headers.XNext = xNext
	return out
}

type ReviewShopResponse200JSON struct {
	Body    Pet
	Headers struct {
		Body  Pet
		XNext string
	}
}

func (r ReviewShopResponse200JSON) reviewShop() {}

func (r ReviewShopResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("x-next", r.Headers.XNext)
	w.WriteHeader(200)
	writeJSON(w, r.Body, "ReviewShopResponse200JSON")
}

func NewReviewShopResponseDefaultJSON(code int, body Error) ReviewShopResponse {
	var out ReviewShopResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

type ReviewShopResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r ReviewShopResponseDefaultJSON) reviewShop() {}

func (r ReviewShopResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "ReviewShopResponseDefaultJSON")
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
