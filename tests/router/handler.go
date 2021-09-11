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

type GetRTHandlerFunc func(GetRTParamsParser) GetRTResponser

func (f GetRTHandlerFunc) Handle(p GetRTParamsParser) GetRTResponser {
	return f(p)
}

func (f GetRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetRTParams{Request: r}).writeGetRTResponse(w)
}

type GetRTParamsParser interface {
	Parse() GetRTParams
}

type requestGetRTParams struct {
	Request *http.Request
}

func (p requestGetRTParams) Parse() GetRTParams {
	return newGetRTParams(p.Request)
}

type GetRTParams struct {
	HTTPRequest *http.Request
}

func newGetRTParams(r *http.Request) (zero GetRTParams) {
	var params GetRTParams
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

type GetShopsHandlerFunc func(GetShopsParamsParser) GetShopsResponser

func (f GetShopsHandlerFunc) Handle(p GetShopsParamsParser) GetShopsResponser {
	return f(p)
}

func (f GetShopsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsParams{Request: r}).writeGetShopsResponse(w)
}

type GetShopsParamsParser interface {
	Parse() GetShopsParams
}

type requestGetShopsParams struct {
	Request *http.Request
}

func (p requestGetShopsParams) Parse() GetShopsParams {
	return newGetShopsParams(p.Request)
}

type GetShopsParams struct {
	HTTPRequest *http.Request
}

func newGetShopsParams(r *http.Request) (zero GetShopsParams) {
	var params GetShopsParams
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

type GetShopsRTHandlerFunc func(GetShopsRTParamsParser) GetShopsRTResponser

func (f GetShopsRTHandlerFunc) Handle(p GetShopsRTParamsParser) GetShopsRTResponser {
	return f(p)
}

func (f GetShopsRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsRTParams{Request: r}).writeGetShopsRTResponse(w)
}

type GetShopsRTParamsParser interface {
	Parse() GetShopsRTParams
}

type requestGetShopsRTParams struct {
	Request *http.Request
}

func (p requestGetShopsRTParams) Parse() GetShopsRTParams {
	return newGetShopsRTParams(p.Request)
}

type GetShopsRTParams struct {
	HTTPRequest *http.Request
}

func newGetShopsRTParams(r *http.Request) (zero GetShopsRTParams) {
	var params GetShopsRTParams
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

type GetShopsActivateHandlerFunc func(GetShopsActivateParamsParser) GetShopsActivateResponser

func (f GetShopsActivateHandlerFunc) Handle(p GetShopsActivateParamsParser) GetShopsActivateResponser {
	return f(p)
}

func (f GetShopsActivateHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsActivateParams{Request: r}).writeGetShopsActivateResponse(w)
}

type GetShopsActivateParamsParser interface {
	Parse() GetShopsActivateParams
}

type requestGetShopsActivateParams struct {
	Request *http.Request
}

func (p requestGetShopsActivateParams) Parse() GetShopsActivateParams {
	return newGetShopsActivateParams(p.Request)
}

type GetShopsActivateParams struct {
	HTTPRequest *http.Request
}

func newGetShopsActivateParams(r *http.Request) (zero GetShopsActivateParams) {
	var params GetShopsActivateParams
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

type GetShopsShopHandlerFunc func(GetShopsShopParamsParser) GetShopsShopResponser

func (f GetShopsShopHandlerFunc) Handle(p GetShopsShopParamsParser) GetShopsShopResponser {
	return f(p)
}

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopParams{Request: r}).writeGetShopsShopResponse(w)
}

type GetShopsShopParamsParser interface {
	Parse() GetShopsShopParams
}

type requestGetShopsShopParams struct {
	Request *http.Request
}

func (p requestGetShopsShopParams) Parse() GetShopsShopParams {
	return newGetShopsShopParams(p.Request)
}

type GetShopsShopParams struct {
	HTTPRequest *http.Request
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopParams) {
	var params GetShopsShopParams
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

type GetShopsShopRTHandlerFunc func(GetShopsShopRTParamsParser) GetShopsShopRTResponser

func (f GetShopsShopRTHandlerFunc) Handle(p GetShopsShopRTParamsParser) GetShopsShopRTResponser {
	return f(p)
}

func (f GetShopsShopRTHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopRTParams{Request: r}).writeGetShopsShopRTResponse(w)
}

type GetShopsShopRTParamsParser interface {
	Parse() GetShopsShopRTParams
}

type requestGetShopsShopRTParams struct {
	Request *http.Request
}

func (p requestGetShopsShopRTParams) Parse() GetShopsShopRTParams {
	return newGetShopsShopRTParams(p.Request)
}

type GetShopsShopRTParams struct {
	HTTPRequest *http.Request
}

func newGetShopsShopRTParams(r *http.Request) (zero GetShopsShopRTParams) {
	var params GetShopsShopRTParams
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

type GetShopsShopPetsHandlerFunc func(GetShopsShopPetsParamsParser) GetShopsShopPetsResponser

func (f GetShopsShopPetsHandlerFunc) Handle(p GetShopsShopPetsParamsParser) GetShopsShopPetsResponser {
	return f(p)
}

func (f GetShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopPetsParams{Request: r}).writeGetShopsShopPetsResponse(w)
}

type GetShopsShopPetsParamsParser interface {
	Parse() GetShopsShopPetsParams
}

type requestGetShopsShopPetsParams struct {
	Request *http.Request
}

func (p requestGetShopsShopPetsParams) Parse() GetShopsShopPetsParams {
	return newGetShopsShopPetsParams(p.Request)
}

type GetShopsShopPetsParams struct {
	HTTPRequest *http.Request
}

func newGetShopsShopPetsParams(r *http.Request) (zero GetShopsShopPetsParams) {
	var params GetShopsShopPetsParams
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
