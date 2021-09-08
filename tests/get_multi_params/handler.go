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
// GetShopsShopPetsPetID -
// ---------------------------------------------

func GetShopsShopPetsPetIDHandler(h GetShopsShopPetsPetIDHandlerer) http.Handler {
	return GetShopsShopPetsPetIDHandlerFunc(h.Handle, h.InvalidResponce)
}

func GetShopsShopPetsPetIDHandlerFunc(fn FuncGetShopsShopPetsPetID, invalidFn FuncGetShopsShopPetsPetIDInvalidResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := newGetShopsShopPetsPetIDParams(r)
		if err != nil {
			invalidFn(err).writeGetShopsShopPetsPetIDResponse(w)
			return
		}

		fn(params).writeGetShopsShopPetsPetIDResponse(w)
	}
}

type GetShopsShopPetsPetIDHandlerer interface {
	Handle(GetShopsShopPetsPetIDParams) GetShopsShopPetsPetIDResponser
	InvalidResponce(error) GetShopsShopPetsPetIDResponser
}

func NewGetShopsShopPetsPetIDHandlerer(fn FuncGetShopsShopPetsPetID, invalidFn FuncGetShopsShopPetsPetIDInvalidResponse) GetShopsShopPetsPetIDHandlerer {
	return privateGetShopsShopPetsPetIDHandlerer{
		FuncGetShopsShopPetsPetID:                fn,
		FuncGetShopsShopPetsPetIDInvalidResponse: invalidFn,
	}
}

type privateGetShopsShopPetsPetIDHandlerer struct {
	FuncGetShopsShopPetsPetID
	FuncGetShopsShopPetsPetIDInvalidResponse
}

type FuncGetShopsShopPetsPetID func(GetShopsShopPetsPetIDParams) GetShopsShopPetsPetIDResponser

func (f FuncGetShopsShopPetsPetID) Handle(params GetShopsShopPetsPetIDParams) GetShopsShopPetsPetIDResponser {
	return f(params)
}

type FuncGetShopsShopPetsPetIDInvalidResponse func(error) GetShopsShopPetsPetIDResponser

func (f FuncGetShopsShopPetsPetIDInvalidResponse) InvalidResponce(err error) GetShopsShopPetsPetIDResponser {
	return f(err)
}

type GetShopsShopPetsPetIDParams struct {
	Request *http.Request

	Color string
	Page  int32
	Shop  string
	PetID int64
}

func newGetShopsShopPetsPetIDParams(r *http.Request) (zero GetShopsShopPetsPetIDParams, _ error) {
	var params GetShopsShopPetsPetIDParams
	params.Request = r

	{
		query := r.URL.Query()
		{
			q, ok := query["color"]
			if ok && len(q) > 0 {
				params.Color = q[0]
			}
		}
		{
			q, ok := query["page"]
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseQueryParam{Name: "page", Err: fmt.Errorf("parse int32: %w", err)}
				}
				params.Page = int32(vInt)
			}
		}
	}

	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets/{petId}'")
		}
		p = p[7:]

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			v := p[:idx]
			p = p[idx:]

			params.Shop = v
		}

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets/{petId}'")
		}
		p = p[6:]

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			v := p[:idx]
			p = p[idx:]

			vInt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return zero, ErrParsePathParam{Name: "petId", Err: fmt.Errorf("parse int64: %w", err)}
			}
			params.PetID = vInt
		}
	}

	return params, nil
}

type GetShopsShopPetsPetIDResponser interface {
	writeGetShopsShopPetsPetIDResponse(w http.ResponseWriter)
}

func GetShopsShopPetsPetIDResponse200() GetShopsShopPetsPetIDResponser {
	var out getShopsShopPetsPetIDResponse200
	return out
}

type getShopsShopPetsPetIDResponse200 struct{}

func (r getShopsShopPetsPetIDResponse200) writeGetShopsShopPetsPetIDResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func GetShopsShopPetsPetIDResponseDefault(code int) GetShopsShopPetsPetIDResponser {
	var out getShopsShopPetsPetIDResponseDefault
	out.Code = code
	return out
}

type getShopsShopPetsPetIDResponseDefault struct {
	Code int
}

func (r getShopsShopPetsPetIDResponseDefault) writeGetShopsShopPetsPetIDResponse(w http.ResponseWriter) {
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
