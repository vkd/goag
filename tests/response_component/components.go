package test

import "net/http"

// ------------------------
//         Schemas
// ------------------------

type Error struct {
	Message string `json:"message"`
}

type Pet struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ------------------------------
//         Responses
// ------------------------------

func NewErrorResponse(code int, body Error) ErrorResponse {
	var out ErrorResponse
	out.Code = code
	out.Body = body
	return out
}

// ErrorResponse - Error output response
type ErrorResponse struct {
	Code int
	Body Error
}

func (r ErrorResponse) writeGetV2Pet(w http.ResponseWriter) {
	r.Write(w)
}

func (r ErrorResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "ErrorResponse")
}

func NewPetResponse(body Pet) PetResponse {
	var out PetResponse
	out.Body = body
	return out
}

// PetResponse - Pet output response
type PetResponse struct {
	Body Pet
}

func (r PetResponse) writeGetPet(w http.ResponseWriter) {
	r.Write(w, 200)
}

func (r PetResponse) writeGetV2Pet(w http.ResponseWriter) {
	r.Write(w, 201)
}

func (r PetResponse) writeGetV3Pet(w http.ResponseWriter) {
	r.Write(w, 202)
}

func (r PetResponse) Write(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, r.Body, "PetResponse")
}

// NewPet2Response - Pet output response
func NewPet2Response(body Pet) Pet2Response {
	return NewPetResponse(body)
}

// Pet2Response - Pet output response
type Pet2Response = PetResponse

// NewPet3Response - Pet output response
func NewPet3Response(body Pet) Pet3Response {
	return NewPet2Response(body)
}

// Pet3Response - Pet output response
type Pet3Response = Pet2Response
