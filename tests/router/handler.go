package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// Get -
// ---------------------------------------------

type GetHandlerFunc func(ctx context.Context, r GetRequest) GetResponse

func (f GetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetHTTPRequest(r)).writeGet(w)
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
	writeGet(http.ResponseWriter)
}

func NewGetResponseDefault(code int) GetResponse {
	var out GetResponseDefault
	out.Code = code
	return out
}

type GetResponseDefault struct {
	Code int
}

func (r GetResponseDefault) writeGet(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShops -
// ---------------------------------------------

type GetShopsHandlerFunc func(ctx context.Context, r GetShopsRequest) GetShopsResponse

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsHTTPRequest(r)).writeGetShops(w)
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
	writeGetShops(http.ResponseWriter)
}

func NewGetShopsResponseDefault(code int) GetShopsResponse {
	var out GetShopsResponseDefault
	out.Code = code
	return out
}

type GetShopsResponseDefault struct {
	Code int
}

func (r GetShopsResponseDefault) writeGetShops(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsRT -
// ---------------------------------------------

type GetShopsRTHandlerFunc func(ctx context.Context, r GetShopsRTRequest) GetShopsRTResponse

func (f GetShopsRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsRTHTTPRequest(r)).writeGetShopsRT(w)
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
	writeGetShopsRT(http.ResponseWriter)
}

func NewGetShopsRTResponseDefault(code int) GetShopsRTResponse {
	var out GetShopsRTResponseDefault
	out.Code = code
	return out
}

type GetShopsRTResponseDefault struct {
	Code int
}

func (r GetShopsRTResponseDefault) writeGetShopsRT(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsRTResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsActivate -
// ---------------------------------------------

type GetShopsActivateHandlerFunc func(ctx context.Context, r GetShopsActivateRequest) GetShopsActivateResponse

func (f GetShopsActivateHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsActivateHTTPRequest(r)).writeGetShopsActivate(w)
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
	writeGetShopsActivate(http.ResponseWriter)
}

func NewGetShopsActivateResponseDefault(code int) GetShopsActivateResponse {
	var out GetShopsActivateResponseDefault
	out.Code = code
	return out
}

type GetShopsActivateResponseDefault struct {
	Code int
}

func (r GetShopsActivateResponseDefault) writeGetShopsActivate(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsActivateResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsMinePetsMikeTails -
// ---------------------------------------------

type GetShopsMinePetsMikeTailsHandlerFunc func(ctx context.Context, r GetShopsMinePetsMikeTailsRequest) GetShopsMinePetsMikeTailsResponse

func (f GetShopsMinePetsMikeTailsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsMinePetsMikeTailsHTTPRequest(r)).writeGetShopsMinePetsMikeTails(w)
}

type GetShopsMinePetsMikeTailsRequest interface {
	HTTP() *http.Request
	Parse() GetShopsMinePetsMikeTailsParams
}

func GetShopsMinePetsMikeTailsHTTPRequest(r *http.Request) GetShopsMinePetsMikeTailsRequest {
	return getShopsMinePetsMikeTailsHTTPRequest{r}
}

type getShopsMinePetsMikeTailsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsMinePetsMikeTailsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsMinePetsMikeTailsHTTPRequest) Parse() GetShopsMinePetsMikeTailsParams {
	return newGetShopsMinePetsMikeTailsParams(r.Request)
}

type GetShopsMinePetsMikeTailsParams struct {
}

func newGetShopsMinePetsMikeTailsParams(r *http.Request) (zero GetShopsMinePetsMikeTailsParams) {
	var params GetShopsMinePetsMikeTailsParams

	return params
}

func (r GetShopsMinePetsMikeTailsParams) HTTP() *http.Request { return nil }

func (r GetShopsMinePetsMikeTailsParams) Parse() GetShopsMinePetsMikeTailsParams { return r }

type GetShopsMinePetsMikeTailsResponse interface {
	writeGetShopsMinePetsMikeTails(http.ResponseWriter)
}

func NewGetShopsMinePetsMikeTailsResponseDefault(code int) GetShopsMinePetsMikeTailsResponse {
	var out GetShopsMinePetsMikeTailsResponseDefault
	out.Code = code
	return out
}

type GetShopsMinePetsMikeTailsResponseDefault struct {
	Code int
}

func (r GetShopsMinePetsMikeTailsResponseDefault) writeGetShopsMinePetsMikeTails(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsMinePetsMikeTailsResponseDefault) Write(w http.ResponseWriter) {
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
	Path GetShopsShopParamsPath
}

type GetShopsShopParamsPath struct {
	Shop string
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopParams, _ error) {
	var params GetShopsShopParams

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
// GetShopsShopRT -
// ---------------------------------------------

type GetShopsShopRTHandlerFunc func(ctx context.Context, r GetShopsShopRTRequest) GetShopsShopRTResponse

func (f GetShopsShopRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopRTHTTPRequest(r)).writeGetShopsShopRT(w)
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
	Path GetShopsShopRTParamsPath
}

type GetShopsShopRTParamsPath struct {
	Shop string
}

func newGetShopsShopRTParams(r *http.Request) (zero GetShopsShopRTParams, _ error) {
	var params GetShopsShopRTParams

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

			params.Path.Shop = vPath
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
	writeGetShopsShopRT(http.ResponseWriter)
}

func NewGetShopsShopRTResponseDefault(code int) GetShopsShopRTResponse {
	var out GetShopsShopRTResponseDefault
	out.Code = code
	return out
}

type GetShopsShopRTResponseDefault struct {
	Code int
}

func (r GetShopsShopRTResponseDefault) writeGetShopsShopRT(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopRTResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsShopPets -
// ---------------------------------------------

type GetShopsShopPetsHandlerFunc func(ctx context.Context, r GetShopsShopPetsRequest) GetShopsShopPetsResponse

func (f GetShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopPetsHTTPRequest(r)).writeGetShopsShopPets(w)
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
	Path GetShopsShopPetsParamsPath
}

type GetShopsShopPetsParamsPath struct {
	Shop string
}

func newGetShopsShopPetsParams(r *http.Request) (zero GetShopsShopPetsParams, _ error) {
	var params GetShopsShopPetsParams

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

			params.Path.Shop = vPath
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
	writeGetShopsShopPets(http.ResponseWriter)
}

func NewGetShopsShopPetsResponseDefault(code int) GetShopsShopPetsResponse {
	var out GetShopsShopPetsResponseDefault
	out.Code = code
	return out
}

type GetShopsShopPetsResponseDefault struct {
	Code int
}

func (r GetShopsShopPetsResponseDefault) writeGetShopsShopPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

// ---------------------------------------------
// GetShopsShopPetsMikePaws -
// ---------------------------------------------

type GetShopsShopPetsMikePawsHandlerFunc func(ctx context.Context, r GetShopsShopPetsMikePawsRequest) GetShopsShopPetsMikePawsResponse

func (f GetShopsShopPetsMikePawsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopPetsMikePawsHTTPRequest(r)).writeGetShopsShopPetsMikePaws(w)
}

type GetShopsShopPetsMikePawsRequest interface {
	HTTP() *http.Request
	Parse() (GetShopsShopPetsMikePawsParams, error)
}

func GetShopsShopPetsMikePawsHTTPRequest(r *http.Request) GetShopsShopPetsMikePawsRequest {
	return getShopsShopPetsMikePawsHTTPRequest{r}
}

type getShopsShopPetsMikePawsHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopPetsMikePawsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsShopPetsMikePawsHTTPRequest) Parse() (GetShopsShopPetsMikePawsParams, error) {
	return newGetShopsShopPetsMikePawsParams(r.Request)
}

type GetShopsShopPetsMikePawsParams struct {
	Path GetShopsShopPetsMikePawsParamsPath
}

type GetShopsShopPetsMikePawsParamsPath struct {
	Shop string
}

func newGetShopsShopPetsMikePawsParams(r *http.Request) (zero GetShopsShopPetsMikePawsParams, _ error) {
	var params GetShopsShopPetsMikePawsParams

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
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets/mike/paws'")
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

		if !strings.HasPrefix(p, "/pets/mike/paws") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets/mike/paws'")
		}
		p = p[15:] // "/pets/mike/paws"
	}

	return params, nil
}

func (r GetShopsShopPetsMikePawsParams) HTTP() *http.Request { return nil }

func (r GetShopsShopPetsMikePawsParams) Parse() (GetShopsShopPetsMikePawsParams, error) {
	return r, nil
}

type GetShopsShopPetsMikePawsResponse interface {
	writeGetShopsShopPetsMikePaws(http.ResponseWriter)
}

func NewGetShopsShopPetsMikePawsResponseDefault(code int) GetShopsShopPetsMikePawsResponse {
	var out GetShopsShopPetsMikePawsResponseDefault
	out.Code = code
	return out
}

type GetShopsShopPetsMikePawsResponseDefault struct {
	Code int
}

func (r GetShopsShopPetsMikePawsResponseDefault) writeGetShopsShopPetsMikePaws(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopPetsMikePawsResponseDefault) Write(w http.ResponseWriter) {
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
