package test

// ------------------------
//         Schemas
// ------------------------

type NewPet struct {
	Name string `json:"name"`
	Tag  PetTag `json:"tag"`
}
