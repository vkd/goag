package test

import "net/http"

// ------------------------
//         Schemas
// ------------------------

type NewPet struct {
	Name string                  `json:"name"`
	Tag  Nullable[string]        `json:"tag"`
	Tago Maybe[Nullable[string]] `json:"tago"`
}

type Pet struct {
	ID   int64                   `json:"id"`
	Name string                  `json:"name"`
	Tag  Nullable[string]        `json:"tag"`
	Tago Maybe[Nullable[string]] `json:"tago"`
}

// ------------------------------
//         Responses
// ------------------------------

func NewPetResponse(body Pet) PetResponse {
	var out PetResponse
	out.Body = body
	return out
}

// PetResponse - Pet output response
type PetResponse struct {
	Body Pet
}

func (r PetResponse) writePostPets(w http.ResponseWriter) {
	r.Write(w, 200)
}

func (r PetResponse) Write(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, r.Body, "PetResponse")
}
