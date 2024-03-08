package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// ListPets -
// ---------------------------------------------

// ListPetsHandlerFunc - List all pets
type ListPetsHandlerFunc func(ctx context.Context, r ListPetsRequest) ListPetsResponse

func (f ListPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), ListPetsHTTPRequest(r)).Write(w)
}

type ListPetsRequest interface {
	HTTP() *http.Request
	Parse() (ListPetsParams, error)
}

func ListPetsHTTPRequest(r *http.Request) ListPetsRequest {
	return listPetsHTTPRequest{r}
}

type listPetsHTTPRequest struct {
	Request *http.Request
}

func (r listPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r listPetsHTTPRequest) Parse() (ListPetsParams, error) {
	return newListPetsParams(r.Request)
}

type ListPetsParams struct {
	Query struct {

		// Limit - How many items to return at one time (max 100)
		Limit *int32
	}
}

func newListPetsParams(r *http.Request) (zero ListPetsParams, _ error) {
	var params ListPetsParams

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

func (r ListPetsParams) HTTP() *http.Request { return nil }

func (r ListPetsParams) Parse() (ListPetsParams, error) { return r, nil }

type ListPetsResponse interface {
	listPets()
	Write(w http.ResponseWriter)
}

func NewListPetsResponse200JSON(body Pets, xNext string) ListPetsResponse {
	var out ListPetsResponse200JSON
	out.Body = body
	out.Headers.XNext = xNext
	return out
}

// ListPetsResponse200JSON - A paged array of pets
type ListPetsResponse200JSON struct {
	Body    Pets
	Headers struct {
		XNext string
	}
}

func (r ListPetsResponse200JSON) listPets() {}

func (r ListPetsResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("x-next", r.Headers.XNext)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "ListPetsResponse200JSON")
}

func NewListPetsResponseDefaultJSON(code int, body Error) ListPetsResponse {
	var out ListPetsResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

// ListPetsResponseDefaultJSON - unexpected error
type ListPetsResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r ListPetsResponseDefaultJSON) listPets() {}

func (r ListPetsResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "ListPetsResponseDefaultJSON")
}

// ---------------------------------------------
// CreatePets -
// ---------------------------------------------

// CreatePetsHandlerFunc - Create a pet
type CreatePetsHandlerFunc func(ctx context.Context, r CreatePetsRequest) CreatePetsResponse

func (f CreatePetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), CreatePetsHTTPRequest(r)).Write(w)
}

type CreatePetsRequest interface {
	HTTP() *http.Request
	Parse() CreatePetsParams
}

func CreatePetsHTTPRequest(r *http.Request) CreatePetsRequest {
	return createPetsHTTPRequest{r}
}

type createPetsHTTPRequest struct {
	Request *http.Request
}

func (r createPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r createPetsHTTPRequest) Parse() CreatePetsParams {
	return newCreatePetsParams(r.Request)
}

type CreatePetsParams struct {
}

func newCreatePetsParams(r *http.Request) (zero CreatePetsParams) {
	var params CreatePetsParams

	return params
}

func (r CreatePetsParams) HTTP() *http.Request { return nil }

func (r CreatePetsParams) Parse() CreatePetsParams { return r }

type CreatePetsResponse interface {
	createPets()
	Write(w http.ResponseWriter)
}

func NewCreatePetsResponse201() CreatePetsResponse {
	var out CreatePetsResponse201
	return out
}

// CreatePetsResponse201 - Null response
type CreatePetsResponse201 struct{}

func (r CreatePetsResponse201) createPets() {}

func (r CreatePetsResponse201) Write(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func NewCreatePetsResponseDefaultJSON(code int, body Error) CreatePetsResponse {
	var out CreatePetsResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

// CreatePetsResponseDefaultJSON - unexpected error
type CreatePetsResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r CreatePetsResponseDefaultJSON) createPets() {}

func (r CreatePetsResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "CreatePetsResponseDefaultJSON")
}

// ---------------------------------------------
// ShowPetByID -
// ---------------------------------------------

// ShowPetByIDHandlerFunc - Info for a specific pet
type ShowPetByIDHandlerFunc func(ctx context.Context, r ShowPetByIDRequest) ShowPetByIDResponse

func (f ShowPetByIDHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), ShowPetByIDHTTPRequest(r)).Write(w)
}

type ShowPetByIDRequest interface {
	HTTP() *http.Request
	Parse() (ShowPetByIDParams, error)
}

func ShowPetByIDHTTPRequest(r *http.Request) ShowPetByIDRequest {
	return showPetByIDHTTPRequest{r}
}

type showPetByIDHTTPRequest struct {
	Request *http.Request
}

func (r showPetByIDHTTPRequest) HTTP() *http.Request { return r.Request }

func (r showPetByIDHTTPRequest) Parse() (ShowPetByIDParams, error) {
	return newShowPetByIDParams(r.Request)
}

type ShowPetByIDParams struct {
	Path struct {

		// PetID - The id of the pet to retrieve
		PetID string
	}
}

func newShowPetByIDParams(r *http.Request) (zero ShowPetByIDParams, _ error) {
	var params ShowPetByIDParams

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

func (r ShowPetByIDParams) HTTP() *http.Request { return nil }

func (r ShowPetByIDParams) Parse() (ShowPetByIDParams, error) { return r, nil }

type ShowPetByIDResponse interface {
	showPetByID()
	Write(w http.ResponseWriter)
}

func NewShowPetByIDResponse200JSON(body Pet) ShowPetByIDResponse {
	var out ShowPetByIDResponse200JSON
	out.Body = body
	return out
}

// ShowPetByIDResponse200JSON - Expected response to a valid request
type ShowPetByIDResponse200JSON struct {
	Body Pet
}

func (r ShowPetByIDResponse200JSON) showPetByID() {}

func (r ShowPetByIDResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "ShowPetByIDResponse200JSON")
}

func NewShowPetByIDResponseDefaultJSON(code int, body Error) ShowPetByIDResponse {
	var out ShowPetByIDResponseDefaultJSON
	out.Code = code
	out.Body = body
	return out
}

// ShowPetByIDResponseDefaultJSON - unexpected error
type ShowPetByIDResponseDefaultJSON struct {
	Code int
	Body Error
}

func (r ShowPetByIDResponseDefaultJSON) showPetByID() {}

func (r ShowPetByIDResponseDefaultJSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "ShowPetByIDResponseDefaultJSON")
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
