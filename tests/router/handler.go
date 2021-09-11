package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetRT -
// ---------------------------------------------

type GetRTHandlerFunc func(r GetRTRequester) GetRTResponser

func (f GetRTHandlerFunc) Handle(r GetRTRequester) GetRTResponser {
	return f(r)
}

func (f GetRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetRTParams{Request: r}).writeGetRTResponse(w)
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

type GetRTResponser interface {
	writeGetRTResponse(w http.ResponseWriter)
}

func GetRTResponseDefault(code int) GetRTResponser {
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

type GetShopsHandlerFunc func(r GetShopsRequester) GetShopsResponser

func (f GetShopsHandlerFunc) Handle(r GetShopsRequester) GetShopsResponser {
	return f(r)
}

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsParams{Request: r}).writeGetShopsResponse(w)
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

type GetShopsResponser interface {
	writeGetShopsResponse(w http.ResponseWriter)
}

func GetShopsResponseDefault(code int) GetShopsResponser {
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

type GetShopsRTHandlerFunc func(r GetShopsRTRequester) GetShopsRTResponser

func (f GetShopsRTHandlerFunc) Handle(r GetShopsRTRequester) GetShopsRTResponser {
	return f(r)
}

func (f GetShopsRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsRTParams{Request: r}).writeGetShopsRTResponse(w)
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

type GetShopsRTResponser interface {
	writeGetShopsRTResponse(w http.ResponseWriter)
}

func GetShopsRTResponseDefault(code int) GetShopsRTResponser {
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

type GetShopsActivateHandlerFunc func(r GetShopsActivateRequester) GetShopsActivateResponser

func (f GetShopsActivateHandlerFunc) Handle(r GetShopsActivateRequester) GetShopsActivateResponser {
	return f(r)
}

func (f GetShopsActivateHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsActivateParams{Request: r}).writeGetShopsActivateResponse(w)
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

type GetShopsActivateResponser interface {
	writeGetShopsActivateResponse(w http.ResponseWriter)
}

func GetShopsActivateResponseDefault(code int) GetShopsActivateResponser {
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

type GetShopsShopHandlerFunc func(r GetShopsShopRequester) GetShopsShopResponser

func (f GetShopsShopHandlerFunc) Handle(r GetShopsShopRequester) GetShopsShopResponser {
	return f(r)
}

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopParams{Request: r}).writeGetShopsShopResponse(w)
}

type GetShopsShopRequester interface {
	Parse() GetShopsShopRequest
}

type requestGetShopsShopParams struct {
	Request *http.Request
}

func (r requestGetShopsShopParams) Parse() GetShopsShopRequest {
	return newGetShopsShopParams(r.Request)
}

type GetShopsShopRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopRequest) {
	var params GetShopsShopRequest
	params.HTTPRequest = r

	return params
}

type GetShopsShopResponser interface {
	writeGetShopsShopResponse(w http.ResponseWriter)
}

func GetShopsShopResponseDefault(code int) GetShopsShopResponser {
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

type GetShopsShopRTHandlerFunc func(r GetShopsShopRTRequester) GetShopsShopRTResponser

func (f GetShopsShopRTHandlerFunc) Handle(r GetShopsShopRTRequester) GetShopsShopRTResponser {
	return f(r)
}

func (f GetShopsShopRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopRTParams{Request: r}).writeGetShopsShopRTResponse(w)
}

type GetShopsShopRTRequester interface {
	Parse() GetShopsShopRTRequest
}

type requestGetShopsShopRTParams struct {
	Request *http.Request
}

func (r requestGetShopsShopRTParams) Parse() GetShopsShopRTRequest {
	return newGetShopsShopRTParams(r.Request)
}

type GetShopsShopRTRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsShopRTParams(r *http.Request) (zero GetShopsShopRTRequest) {
	var params GetShopsShopRTRequest
	params.HTTPRequest = r

	return params
}

type GetShopsShopRTResponser interface {
	writeGetShopsShopRTResponse(w http.ResponseWriter)
}

func GetShopsShopRTResponseDefault(code int) GetShopsShopRTResponser {
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

type GetShopsShopPetsHandlerFunc func(r GetShopsShopPetsRequester) GetShopsShopPetsResponser

func (f GetShopsShopPetsHandlerFunc) Handle(r GetShopsShopPetsRequester) GetShopsShopPetsResponser {
	return f(r)
}

func (f GetShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopPetsParams{Request: r}).writeGetShopsShopPetsResponse(w)
}

type GetShopsShopPetsRequester interface {
	Parse() GetShopsShopPetsRequest
}

type requestGetShopsShopPetsParams struct {
	Request *http.Request
}

func (r requestGetShopsShopPetsParams) Parse() GetShopsShopPetsRequest {
	return newGetShopsShopPetsParams(r.Request)
}

type GetShopsShopPetsRequest struct {
	HTTPRequest *http.Request
}

func newGetShopsShopPetsParams(r *http.Request) (zero GetShopsShopPetsRequest) {
	var params GetShopsShopPetsRequest
	params.HTTPRequest = r

	return params
}

type GetShopsShopPetsResponser interface {
	writeGetShopsShopPetsResponse(w http.ResponseWriter)
}

func GetShopsShopPetsResponseDefault(code int) GetShopsShopPetsResponser {
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

func writeJSON(w io.Writer, v interface{}, name string) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		LogError(fmt.Errorf("write json response %q: %w", name, err))
	}
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
