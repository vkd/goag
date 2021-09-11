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

type GetShopsShopPetsPetIDHandlerFunc func(r GetShopsShopPetsPetIDRequester) GetShopsShopPetsPetIDResponser

func (f GetShopsShopPetsPetIDHandlerFunc) Handle(r GetShopsShopPetsPetIDRequester) GetShopsShopPetsPetIDResponser {
	return f(r)
}

func (f GetShopsShopPetsPetIDHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handle(requestGetShopsShopPetsPetIDParams{Request: r}).writeGetShopsShopPetsPetIDResponse(w)
}

type GetShopsShopPetsPetIDRequester interface {
	Parse() (GetShopsShopPetsPetIDRequest, error)
}

type requestGetShopsShopPetsPetIDParams struct {
	Request *http.Request
}

func (r requestGetShopsShopPetsPetIDParams) Parse() (GetShopsShopPetsPetIDRequest, error) {
	return newGetShopsShopPetsPetIDParams(r.Request)
}

type GetShopsShopPetsPetIDRequest struct {
	HTTPRequest *http.Request

	Color *string
	Page  *int32
	Shop  string
	PetID int64
}

func newGetShopsShopPetsPetIDParams(r *http.Request) (zero GetShopsShopPetsPetIDRequest, _ error) {
	var params GetShopsShopPetsPetIDRequest
	params.HTTPRequest = r

	{
		query := r.URL.Query()
		{
			q, ok := query["color"]
			if ok && len(q) > 0 {
				v := q[0]
				params.Color = &v
			}
		}
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

	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets/{petId}'")
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

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets/{petId}'")
		}
		p = p[6:] // "/pets/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParsePathParam{Name: "petId", Err: fmt.Errorf("is required")}
			}

			vInt, err := strconv.ParseInt(vPath, 10, 64)
			if err != nil {
				return zero, ErrParsePathParam{Name: "petId", Err: fmt.Errorf("parse int64: %w", err)}
			}
			v := vInt
			params.PetID = v
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
