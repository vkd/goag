// Code generated by goag (https://github.com/vkd/goag). DO NOT EDIT.
package test

import (
	"bytes"
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
	f(r.Context(), ListPetsHTTPRequest(r)).writeListPets(w)
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
	Query ListPetsParamsQuery
}

type ListPetsParamsQuery struct {

	// Limit - How many items to return at one time (max 100)
	Limit Maybe[int32]
}

func newListPetsParams(r *http.Request) (zero ListPetsParams, _ error) {
	var params ListPetsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["limit"]
			if ok && len(q) > 0 {
				var vOpt int32
				if len(q) == 1 {
					vInt64, err := strconv.ParseInt(q[0], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "limit", Reason: "parse int32", Err: err}
					}
					vOpt = int32(vInt64)
				} else {
					return zero, ErrParseParam{In: "query", Parameter: "limit", Reason: "multiple values found: single value expected"}
				}
				params.Query.Limit.Set(vOpt)
			}
		}
	}

	return params, nil
}

func (r ListPetsParams) HTTP() *http.Request { return nil }

func (r ListPetsParams) Parse() (ListPetsParams, error) { return r, nil }

type ListPetsResponse interface {
	writeListPets(http.ResponseWriter)
}

func NewListPetsResponse200JSON(body Pets, xNext Maybe[string]) ListPetsResponse {
	var out ListPetsResponse200JSON
	out.Body = body
	out.Headers.XNext = xNext
	return out
}

// ListPetsResponse200JSON - A paged array of pets
type ListPetsResponse200JSON struct {
	Body    Pets
	Headers struct {
		XNext Maybe[string]
	}
}

func (r ListPetsResponse200JSON) writeListPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r ListPetsResponse200JSON) Write(w http.ResponseWriter) {
	if r.Headers.XNext.IsSet {
		hs := []string{r.Headers.XNext.Value}
		for _, h := range hs {
			w.Header().Add("x-next", h)
		}
	}
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

func (r ListPetsResponseDefaultJSON) writeListPets(w http.ResponseWriter) {
	r.Write(w)
}

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
	f(r.Context(), CreatePetsHTTPRequest(r)).writeCreatePets(w)
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
	writeCreatePets(http.ResponseWriter)
}

func NewCreatePetsResponse201() CreatePetsResponse {
	var out CreatePetsResponse201
	return out
}

// CreatePetsResponse201 - Null response
type CreatePetsResponse201 struct{}

func (r CreatePetsResponse201) writeCreatePets(w http.ResponseWriter) {
	r.Write(w)
}

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

func (r CreatePetsResponseDefaultJSON) writeCreatePets(w http.ResponseWriter) {
	r.Write(w)
}

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
	f(r.Context(), ShowPetByIDHTTPRequest(r)).writeShowPetByID(w)
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
	Path ShowPetByIDParamsPath
}

type ShowPetByIDParamsPath struct {

	// PetID - The id of the pet to retrieve
	PetID string
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

			params.Path.PetID = vPath
		}
	}

	return params, nil
}

func (r ShowPetByIDParams) HTTP() *http.Request { return nil }

func (r ShowPetByIDParams) Parse() (ShowPetByIDParams, error) { return r, nil }

type ShowPetByIDResponse interface {
	writeShowPetByID(http.ResponseWriter)
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

func (r ShowPetByIDResponse200JSON) writeShowPetByID(w http.ResponseWriter) {
	r.Write(w)
}

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

func (r ShowPetByIDResponseDefaultJSON) writeShowPetByID(w http.ResponseWriter) {
	r.Write(w)
}

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

var _ json.Marshaler = (*Nullable[any])(nil)

func (m Nullable[T]) MarshalJSON() ([]byte, error) {
	if m.IsSet {
		return json.Marshal(&m.Value)
	}
	return []byte(nullValue), nil
}

var _ json.Unmarshaler = (*Nullable[any])(nil)

const nullValue = "null"

var nullValueBs = []byte(nullValue)

func (m *Nullable[T]) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, nullValueBs) {
		m.IsSet = false
		return nil
	}
	m.IsSet = true
	return json.Unmarshal(bs, &m.Value)
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
