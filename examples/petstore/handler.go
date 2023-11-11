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
// GetPets -
// ---------------------------------------------

// GetPetsHandlerFunc - List all pets
type GetPetsHandlerFunc func(r GetPetsRequestParser) GetPetsResponse

func (f GetPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetsHTTPRequest(r)).Write(w)
}

type GetPetsRequestParser interface {
	Parse() (GetPetsRequest, error)
}

func GetPetsHTTPRequest(r *http.Request) GetPetsRequestParser {
	return getPetsHTTPRequest{r}
}

type getPetsHTTPRequest struct {
	Request *http.Request
}

func (r getPetsHTTPRequest) Parse() (GetPetsRequest, error) {
	return newGetPetsParams(r.Request)
}

type GetPetsRequest struct {
	HTTPRequest *http.Request

	Query struct {

		// Limit - How many items to return at one time (max 100)
		Limit *int32
	}
}

func newGetPetsParams(r *http.Request) (zero GetPetsRequest, _ error) {
	var params GetPetsRequest
	params.HTTPRequest = r

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["limit"]
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "limit", Reason: "parse int32", Err: err}
				}
				v := int32(vInt)
				params.Query.Limit = &v
			}
		}
	}

	return params, nil
}

func (r GetPetsRequest) Parse() (GetPetsRequest, error) { return r, nil }

type GetPetsResponse interface {
	getPets()
	Write(w http.ResponseWriter)
}

func NewGetPetsResponse200JSON(body Pets, xNext string) GetPetsResponse {
	var out GetPetsResponse200JSON
	out.Body = body
	out.Headers.XNext = xNext
	return out
}

// GetPetsResponse200JSON - A paged array of pets
type GetPetsResponse200JSON struct {
	Body    Pets
	Headers struct {
		Body  Pets
		XNext string
	}
}

func (r GetPetsResponse200JSON) getPets() {}

func (r GetPetsResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("x-next", r.Headers.XNext)
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsResponse200JSON")
}

func NewGetPetsResponseDefaultJSON(code int, body Error) GetPetsResponse {
	var out GetPetsResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

// GetPetsResponseDefaultJSON - unexpected error
type GetPetsResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r GetPetsResponseDefaultJSON) getPets() {}

func (r GetPetsResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "GetPetsResponseDefaultJSON")
}

// ---------------------------------------------
// PostPets -
// ---------------------------------------------

// PostPetsHandlerFunc - Create a pet
type PostPetsHandlerFunc func(r PostPetsRequestParser) PostPetsResponse

func (f PostPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(PostPetsHTTPRequest(r)).Write(w)
}

type PostPetsRequestParser interface {
	Parse() PostPetsRequest
}

func PostPetsHTTPRequest(r *http.Request) PostPetsRequestParser {
	return postPetsHTTPRequest{r}
}

type postPetsHTTPRequest struct {
	Request *http.Request
}

func (r postPetsHTTPRequest) Parse() PostPetsRequest {
	return newPostPetsParams(r.Request)
}

type PostPetsRequest struct {
	HTTPRequest *http.Request
}

func newPostPetsParams(r *http.Request) (zero PostPetsRequest) {
	var params PostPetsRequest
	params.HTTPRequest = r

	return params
}

func (r PostPetsRequest) Parse() PostPetsRequest { return r }

type PostPetsResponse interface {
	postPets()
	Write(w http.ResponseWriter)
}

func NewPostPetsResponse201() PostPetsResponse {
	var out PostPetsResponse201
	return out
}

// PostPetsResponse201 - Null response
type PostPetsResponse201 struct{}

func (r PostPetsResponse201) postPets() {}

func (r PostPetsResponse201) Write(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func NewPostPetsResponseDefaultJSON(code int, body Error) PostPetsResponse {
	var out PostPetsResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

// PostPetsResponseDefaultJSON - unexpected error
type PostPetsResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r PostPetsResponseDefaultJSON) postPets() {}

func (r PostPetsResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "PostPetsResponseDefaultJSON")
}

// ---------------------------------------------
// GetPetsPetID -
// ---------------------------------------------

// GetPetsPetIDHandlerFunc - Info for a specific pet
type GetPetsPetIDHandlerFunc func(r GetPetsPetIDRequestParser) GetPetsPetIDResponse

func (f GetPetsPetIDHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(GetPetsPetIDHTTPRequest(r)).Write(w)
}

type GetPetsPetIDRequestParser interface {
	Parse() (GetPetsPetIDRequest, error)
}

func GetPetsPetIDHTTPRequest(r *http.Request) GetPetsPetIDRequestParser {
	return getPetsPetIDHTTPRequest{r}
}

type getPetsPetIDHTTPRequest struct {
	Request *http.Request
}

func (r getPetsPetIDHTTPRequest) Parse() (GetPetsPetIDRequest, error) {
	return newGetPetsPetIDParams(r.Request)
}

type GetPetsPetIDRequest struct {
	HTTPRequest *http.Request

	Path struct {

		// PetID - The id of the pet to retrieve
		PetID string
	}
}

func newGetPetsPetIDParams(r *http.Request) (zero GetPetsPetIDRequest, _ error) {
	var params GetPetsPetIDRequest
	params.HTTPRequest = r

	// Path parameters
	{
		p := r.URL.Path
		if !strings.HasPrefix(p, "/v1") {
			return zero, fmt.Errorf("wrong path: expected '/v1...'")
		}
		p = p[3:] // "/v1"

		if !strings.HasPrefix(p, "/") {
			return zero, fmt.Errorf("wrong path: expected '/v1/...'")
		}

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{petId}'")
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
				return zero, ErrParseParam{In: "path", Parameter: "petId", Reason: "required"}
			}

			v := vPath
			params.Path.PetID = v
		}
	}

	return params, nil
}

func (r GetPetsPetIDRequest) Parse() (GetPetsPetIDRequest, error) { return r, nil }

type GetPetsPetIDResponse interface {
	getPetsPetID()
	Write(w http.ResponseWriter)
}

func NewGetPetsPetIDResponse200JSON(body Pet) GetPetsPetIDResponse {
	var out GetPetsPetIDResponse200JSON
	out.Body = body
	return out
}

// GetPetsPetIDResponse200JSON - Expected response to a valid request
type GetPetsPetIDResponse200JSON struct {
	Body Pet
}

func (r GetPetsPetIDResponse200JSON) getPetsPetID() {}

func (r GetPetsPetIDResponse200JSON) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsPetIDResponse200JSON")
}

func NewGetPetsPetIDResponseDefaultJSON(code int, body Error) GetPetsPetIDResponse {
	var out GetPetsPetIDResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

// GetPetsPetIDResponseDefaultJSON - unexpected error
type GetPetsPetIDResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r GetPetsPetIDResponseDefaultJSON) getPetsPetID() {}

func (r GetPetsPetIDResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "GetPetsPetIDResponseDefaultJSON")
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
