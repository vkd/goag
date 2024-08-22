package test

import (
	"fmt"
	"net/http"
	"strconv"
)

// ------------------------
//         Schemas
// ------------------------

type Error struct {
	Detail string `json:"detail"`
}

type Pet struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Pets []Pet

// ------------------------
//         Headers
// ------------------------

// ErrorCode - Error code
type ErrorCode int

func (h ErrorCode) String() string {
	return strconv.FormatInt(int64(int(h)), 10)
}

func (c *ErrorCode) Parse(s string) error {
	vInt64, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v := int(vInt64)
	*c = ErrorCode(v)
	return nil
}

// ------------------------------
//         RequestBodies
// ------------------------------

type CreatePetJSON struct {
	Name string `json:"name"`
}

// ------------------------------
//         Responses
// ------------------------------

func NewErrorResponseResponse(body Error, xErrorCode Maybe[int]) ErrorResponseResponse {
	var out ErrorResponseResponse
	out.Body = body
	out.Headers.XErrorCode = xErrorCode
	return out
}

// ErrorResponseResponse - Error response
type ErrorResponseResponse struct {
	Body    Error
	Headers struct {
		XErrorCode Maybe[int]
	}
}

func (r ErrorResponseResponse) writePostPets(w http.ResponseWriter) {
	r.Write(w, 500)
}

func (r ErrorResponseResponse) writePostShops(w http.ResponseWriter) {
	r.Write(w, 500)
}

func (r ErrorResponseResponse) Write(w http.ResponseWriter, code int) {
	if r.Headers.XErrorCode.IsSet {
		hs := []string{strconv.FormatInt(int64(r.Headers.XErrorCode.Value), 10)}
		for _, h := range hs {
			w.Header().Add("X-Error-Code", h)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, r.Body, "ErrorResponseResponse")
}
