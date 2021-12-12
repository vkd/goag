package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPets -
// ---------------------------------------------

type GetPetsHandlerFunc func(r GetPetsRequester) GetPetsResponder

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetsParams{Request: r}).writeGetPetsResponse(w)
}

type GetPetsRequester interface {
	Parse() GetPetsRequest
}

type requestGetPetsParams struct {
	Request *http.Request
}

func (r requestGetPetsParams) Parse() GetPetsRequest {
	return newGetPetsParams(r.Request)
}

type GetPetsRequest struct {
	HTTPRequest *http.Request
}

func newGetPetsParams(r *http.Request) (zero GetPetsRequest) {
	var params GetPetsRequest
	params.HTTPRequest = r

	return params
}

type GetPetsResponder interface {
	writeGetPetsResponse(w http.ResponseWriter)
}

func GetPetsResponse200JSON(body []Pet) GetPetsResponder {
	var out getPetsResponse200JSON
	out.Body = body
	return out
}

type getPetsResponse200JSON struct {
	Body []Pet
}

func (r getPetsResponse200JSON) writeGetPetsResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsResponse200JSON")
}

// ---------------------------------------------
// GetPetsNames -
// ---------------------------------------------

type GetPetsNamesHandlerFunc func(r GetPetsNamesRequester) GetPetsNamesResponder

func (f GetPetsNamesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(requestGetPetsNamesParams{Request: r}).writeGetPetsNamesResponse(w)
}

type GetPetsNamesRequester interface {
	Parse() GetPetsNamesRequest
}

type requestGetPetsNamesParams struct {
	Request *http.Request
}

func (r requestGetPetsNamesParams) Parse() GetPetsNamesRequest {
	return newGetPetsNamesParams(r.Request)
}

type GetPetsNamesRequest struct {
	HTTPRequest *http.Request
}

func newGetPetsNamesParams(r *http.Request) (zero GetPetsNamesRequest) {
	var params GetPetsNamesRequest
	params.HTTPRequest = r

	return params
}

type GetPetsNamesResponder interface {
	writeGetPetsNamesResponse(w http.ResponseWriter)
}

func GetPetsNamesResponse200JSON(body []string) GetPetsNamesResponder {
	var out getPetsNamesResponse200JSON
	out.Body = body
	return out
}

type getPetsNamesResponse200JSON struct {
	Body []string
}

func (r getPetsNamesResponse200JSON) writeGetPetsNamesResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsNamesResponse200JSON")
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
