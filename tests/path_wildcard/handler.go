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
// GetPetsPetID -
// ---------------------------------------------

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
		PetID int32
	}
}

func newGetPetsPetIDParams(r *http.Request) (zero GetPetsPetIDRequest, _ error) {
	var params GetPetsPetIDRequest
	params.HTTPRequest = r

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/pets/") {
			return zero, fmt.Errorf("wrong path: expected '/pets/{pet_id}'")
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
				return zero, ErrParseParam{In: "path", Parameter: "pet_id", Reason: "required"}
			}

			vInt, err := strconv.ParseInt(vPath, 10, 32)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "pet_id", Reason: "parse int32", Err: err}
			}
			v := int32(vInt)
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

type GetPetsPetIDResponse200JSON struct {
	Body Pet
}

func (r GetPetsPetIDResponse200JSON) getPetsPetID() {}

func (r GetPetsPetIDResponse200JSON) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetPetsPetIDResponse200JSON")
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
