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
// GetShopsActivateRT -
// ---------------------------------------------

type GetShopsActivateRTHandlerFunc func(r GetShopsActivateRTRequestParser) GetShopsActivateRTResponse

func (f GetShopsActivateRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsActivateRTHTTPRequest(r)).Write(w)
}

type GetShopsActivateRTRequestParser interface {
	Parse() GetShopsActivateRTRequest
}

func GetShopsActivateRTHTTPRequest(r *http.Request) GetShopsActivateRTRequestParser {
	return getShopsActivateRTHTTPRequest{r}
}

type getShopsActivateRTHTTPRequest struct {
	Request *http.Request
}

func (r getShopsActivateRTHTTPRequest) Parse() GetShopsActivateRTRequest {
	return newGetShopsActivateRTParams(r.Request)
}

type GetShopsActivateRTRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsActivateRTParams(r *http.Request) (zero GetShopsActivateRTRequest) {
	var params GetShopsActivateRTRequest
	params.HTTPRequest = r

	return params
}

func (r GetShopsActivateRTRequest) Parse() GetShopsActivateRTRequest { return r }

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

type GetShopsActivateTagHandlerFunc func(r GetShopsActivateTagRequestParser) GetShopsActivateTagResponse

func (f GetShopsActivateTagHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetShopsActivateTagHTTPRequest(r)).Write(w)
}

type GetShopsActivateTagRequestParser interface {
	Parse() GetShopsActivateTagRequest
}

func GetShopsActivateTagHTTPRequest(r *http.Request) GetShopsActivateTagRequestParser {
	return getShopsActivateTagHTTPRequest{r}
}

type getShopsActivateTagHTTPRequest struct {
	Request *http.Request
}

func (r getShopsActivateTagHTTPRequest) Parse() GetShopsActivateTagRequest {
	return newGetShopsActivateTagParams(r.Request)
}

type GetShopsActivateTagRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsActivateTagParams(r *http.Request) (zero GetShopsActivateTagRequest) {
	var params GetShopsActivateTagRequest
	params.HTTPRequest = r

	return params
}

func (r GetShopsActivateTagRequest) Parse() GetShopsActivateTagRequest { return r }

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
		Shop int32
	}
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopRequest, _ error) {
	var params GetShopsShopRequest
	params.HTTPRequest = r

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
		Shop int32
	}
}

func newGetShopsShopRTParams(r *http.Request) (zero GetShopsShopRTRequest, _ error) {
	var params GetShopsShopRTRequest
	params.HTTPRequest = r

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

	Query struct {
		Page *int32

		PageSize int32
	}

	Path struct {
		Shop int32
	}
}

func newGetShopsShopPetsParams(r *http.Request) (zero GetShopsShopPetsRequest, _ error) {
	var params GetShopsShopPetsRequest
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

func (r GetShopsShopPetsRequest) Parse() (GetShopsShopPetsRequest, error) { return r, nil }

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
// PostShopsShopReview -
// ---------------------------------------------

type PostShopsShopReviewHandlerFunc func(r PostShopsShopReviewRequestParser) PostShopsShopReviewResponse

func (f PostShopsShopReviewHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(PostShopsShopReviewHTTPRequest(r)).Write(w)
}

type PostShopsShopReviewRequestParser interface {
	Parse() (PostShopsShopReviewRequest, error)
}

func PostShopsShopReviewHTTPRequest(r *http.Request) PostShopsShopReviewRequestParser {
	return postShopsShopReviewHTTPRequest{r}
}

type postShopsShopReviewHTTPRequest struct {
	Request *http.Request
}

func (r postShopsShopReviewHTTPRequest) Parse() (PostShopsShopReviewRequest, error) {
	return newPostShopsShopReviewParams(r.Request)
}

type PostShopsShopReviewRequest struct {
	HTTPRequest *http.Request

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

func newPostShopsShopReviewParams(r *http.Request) (zero PostShopsShopReviewRequest, _ error) {
	var params PostShopsShopReviewRequest
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

func (r PostShopsShopReviewRequest) Parse() (PostShopsShopReviewRequest, error) { return r, nil }

type PostShopsShopReviewResponse interface {
	postShopsShopReview()
	Write(w http.ResponseWriter)
}

func NewPostShopsShopReviewResponse200JSON(body Pet, xNext string) PostShopsShopReviewResponse {
	var out PostShopsShopReviewResponse200JSON
	out.Body = body
	out.Headers.XNext = xNext
	return out
}

type PostShopsShopReviewResponse200JSON struct {
	Body    Pet
	Headers struct {
		Body  Pet
		XNext string
	}
}

func (r PostShopsShopReviewResponse200JSON) postShopsShopReview() {}

func (r PostShopsShopReviewResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("x-next", r.Headers.XNext)
	w.WriteHeader(200)
	writeJSON(w, r.Body, "PostShopsShopReviewResponse200JSON")
}

func NewPostShopsShopReviewResponseDefaultJSON(code int, body Error) PostShopsShopReviewResponse {
	var out PostShopsShopReviewResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

type PostShopsShopReviewResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r PostShopsShopReviewResponseDefaultJSON) postShopsShopReview() {}

func (r PostShopsShopReviewResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "PostShopsShopReviewResponseDefaultJSON")
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
