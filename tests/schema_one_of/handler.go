package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ---------------------------------------------
// GetPet -
// ---------------------------------------------

type GetPetHandlerFunc func(ctx context.Context, r GetPetRequest) GetPetResponse

func (f GetPetHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPetHTTPRequest(r)).writeGetPet(w)
}

type GetPetRequest interface {
	HTTP() *http.Request
	Parse() (GetPetParams, error)
}

func GetPetHTTPRequest(r *http.Request) GetPetRequest {
	return getPetHTTPRequest{r}
}

type getPetHTTPRequest struct {
	Request *http.Request
}

func (r getPetHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPetHTTPRequest) Parse() (GetPetParams, error) {
	return newGetPetParams(r.Request)
}

type GetPetParams struct {
	Body Pet
}

func newGetPetParams(r *http.Request) (zero GetPetParams, _ error) {
	var params GetPetParams

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r GetPetParams) HTTP() *http.Request { return nil }

func (r GetPetParams) Parse() (GetPetParams, error) { return r, nil }

type GetPetResponse interface {
	writeGetPet(http.ResponseWriter)
}

func NewGetPetResponse200JSON(body Resp) GetPetResponse {
	var out GetPetResponse200JSON
	out.Body = body
	return out
}

// GetPetResponse200JSON - OK
type GetPetResponse200JSON struct {
	Body Resp
}

func (r GetPetResponse200JSON) writeGetPet(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPetResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetResponse200JSON")
}

// ---------------------------------------------
// GetPet2 -
// ---------------------------------------------

type GetPet2HandlerFunc func(ctx context.Context, r GetPet2Request) GetPet2Response

func (f GetPet2HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetPet2HTTPRequest(r)).writeGetPet2(w)
}

type GetPet2Request interface {
	HTTP() *http.Request
	Parse() (GetPet2Params, error)
}

func GetPet2HTTPRequest(r *http.Request) GetPet2Request {
	return getPet2HTTPRequest{r}
}

type getPet2HTTPRequest struct {
	Request *http.Request
}

func (r getPet2HTTPRequest) HTTP() *http.Request { return r.Request }

func (r getPet2HTTPRequest) Parse() (GetPet2Params, error) {
	return newGetPet2Params(r.Request)
}

type GetPet2Params struct {
	Body Pet2
}

func newGetPet2Params(r *http.Request) (zero GetPet2Params, _ error) {
	var params GetPet2Params

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r GetPet2Params) HTTP() *http.Request { return nil }

func (r GetPet2Params) Parse() (GetPet2Params, error) { return r, nil }

type GetPet2Response interface {
	writeGetPet2(http.ResponseWriter)
}

func NewGetPet2Response200JSON(body Resp2) GetPet2Response {
	var out GetPet2Response200JSON
	out.Body = body
	return out
}

// GetPet2Response200JSON - OK
type GetPet2Response200JSON struct {
	Body Resp2
}

func (r GetPet2Response200JSON) writeGetPet2(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetPet2Response200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPet2Response200JSON")
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
