package test

// ------------------------
//         Schemas
// ------------------------

type NewPet struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// ------------------------------
//         RequestBodies
// ------------------------------

type NewPetJSON NewPet

type Pets2JSON NewPetJSON
