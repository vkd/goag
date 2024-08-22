package test

// ------------------------
//         Schemas
// ------------------------

type NewPet struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Pet struct {
	NewPet
	ID int64 `json:"id"`
}
