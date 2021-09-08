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

func GetRTHandler(h GetRTHandlerer) http.Handler {
	return GetRTHandlerFunc(h.Handle)
}

func GetRTHandlerFunc(fn FuncGetRT) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetRTParams(r)

		fn(params).writeGetRTResponse(w)
	}
}

type GetRTHandlerer interface {
	Handle(GetRTParams) GetRTResponser
}

func NewGetRTHandlerer(fn FuncGetRT) GetRTHandlerer {
	return fn
}

type FuncGetRT func(GetRTParams) GetRTResponser

func (f FuncGetRT) Handle(params GetRTParams) GetRTResponser { return f(params) }

type GetRTParams struct {
	Request *http.Request
}

func newGetRTParams(r *http.Request) (zero GetRTParams) {
	var params GetRTParams
	params.Request = r

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

func GetShopsHandler(h GetShopsHandlerer) http.Handler {
	return GetShopsHandlerFunc(h.Handle)
}

func GetShopsHandlerFunc(fn FuncGetShops) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetShopsParams(r)

		fn(params).writeGetShopsResponse(w)
	}
}

type GetShopsHandlerer interface {
	Handle(GetShopsParams) GetShopsResponser
}

func NewGetShopsHandlerer(fn FuncGetShops) GetShopsHandlerer {
	return fn
}

type FuncGetShops func(GetShopsParams) GetShopsResponser

func (f FuncGetShops) Handle(params GetShopsParams) GetShopsResponser { return f(params) }

type GetShopsParams struct {
	Request *http.Request
}

func newGetShopsParams(r *http.Request) (zero GetShopsParams) {
	var params GetShopsParams
	params.Request = r

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

func GetShopsRTHandler(h GetShopsRTHandlerer) http.Handler {
	return GetShopsRTHandlerFunc(h.Handle)
}

func GetShopsRTHandlerFunc(fn FuncGetShopsRT) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetShopsRTParams(r)

		fn(params).writeGetShopsRTResponse(w)
	}
}

type GetShopsRTHandlerer interface {
	Handle(GetShopsRTParams) GetShopsRTResponser
}

func NewGetShopsRTHandlerer(fn FuncGetShopsRT) GetShopsRTHandlerer {
	return fn
}

type FuncGetShopsRT func(GetShopsRTParams) GetShopsRTResponser

func (f FuncGetShopsRT) Handle(params GetShopsRTParams) GetShopsRTResponser { return f(params) }

type GetShopsRTParams struct {
	Request *http.Request
}

func newGetShopsRTParams(r *http.Request) (zero GetShopsRTParams) {
	var params GetShopsRTParams
	params.Request = r

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

func GetShopsActivateHandler(h GetShopsActivateHandlerer) http.Handler {
	return GetShopsActivateHandlerFunc(h.Handle)
}

func GetShopsActivateHandlerFunc(fn FuncGetShopsActivate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetShopsActivateParams(r)

		fn(params).writeGetShopsActivateResponse(w)
	}
}

type GetShopsActivateHandlerer interface {
	Handle(GetShopsActivateParams) GetShopsActivateResponser
}

func NewGetShopsActivateHandlerer(fn FuncGetShopsActivate) GetShopsActivateHandlerer {
	return fn
}

type FuncGetShopsActivate func(GetShopsActivateParams) GetShopsActivateResponser

func (f FuncGetShopsActivate) Handle(params GetShopsActivateParams) GetShopsActivateResponser {
	return f(params)
}

type GetShopsActivateParams struct {
	Request *http.Request
}

func newGetShopsActivateParams(r *http.Request) (zero GetShopsActivateParams) {
	var params GetShopsActivateParams
	params.Request = r

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

func GetShopsShopHandler(h GetShopsShopHandlerer) http.Handler {
	return GetShopsShopHandlerFunc(h.Handle)
}

func GetShopsShopHandlerFunc(fn FuncGetShopsShop) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetShopsShopParams(r)

		fn(params).writeGetShopsShopResponse(w)
	}
}

type GetShopsShopHandlerer interface {
	Handle(GetShopsShopParams) GetShopsShopResponser
}

func NewGetShopsShopHandlerer(fn FuncGetShopsShop) GetShopsShopHandlerer {
	return fn
}

type FuncGetShopsShop func(GetShopsShopParams) GetShopsShopResponser

func (f FuncGetShopsShop) Handle(params GetShopsShopParams) GetShopsShopResponser { return f(params) }

type GetShopsShopParams struct {
	Request *http.Request
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopParams) {
	var params GetShopsShopParams
	params.Request = r

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

func GetShopsShopRTHandler(h GetShopsShopRTHandlerer) http.Handler {
	return GetShopsShopRTHandlerFunc(h.Handle)
}

func GetShopsShopRTHandlerFunc(fn FuncGetShopsShopRT) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetShopsShopRTParams(r)

		fn(params).writeGetShopsShopRTResponse(w)
	}
}

type GetShopsShopRTHandlerer interface {
	Handle(GetShopsShopRTParams) GetShopsShopRTResponser
}

func NewGetShopsShopRTHandlerer(fn FuncGetShopsShopRT) GetShopsShopRTHandlerer {
	return fn
}

type FuncGetShopsShopRT func(GetShopsShopRTParams) GetShopsShopRTResponser

func (f FuncGetShopsShopRT) Handle(params GetShopsShopRTParams) GetShopsShopRTResponser {
	return f(params)
}

type GetShopsShopRTParams struct {
	Request *http.Request
}

func newGetShopsShopRTParams(r *http.Request) (zero GetShopsShopRTParams) {
	var params GetShopsShopRTParams
	params.Request = r

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

func GetShopsShopPetsHandler(h GetShopsShopPetsHandlerer) http.Handler {
	return GetShopsShopPetsHandlerFunc(h.Handle)
}

func GetShopsShopPetsHandlerFunc(fn FuncGetShopsShopPets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := newGetShopsShopPetsParams(r)

		fn(params).writeGetShopsShopPetsResponse(w)
	}
}

type GetShopsShopPetsHandlerer interface {
	Handle(GetShopsShopPetsParams) GetShopsShopPetsResponser
}

func NewGetShopsShopPetsHandlerer(fn FuncGetShopsShopPets) GetShopsShopPetsHandlerer {
	return fn
}

type FuncGetShopsShopPets func(GetShopsShopPetsParams) GetShopsShopPetsResponser

func (f FuncGetShopsShopPets) Handle(params GetShopsShopPetsParams) GetShopsShopPetsResponser {
	return f(params)
}

type GetShopsShopPetsParams struct {
	Request *http.Request
}

func newGetShopsShopPetsParams(r *http.Request) (zero GetShopsShopPetsParams) {
	var params GetShopsShopPetsParams
	params.Request = r

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
