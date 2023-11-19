package test

// ------------------------
//         Schemas
// ------------------------

type Error struct {
	Message string `json:"message"`
}

type NewPet struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Pet struct {
	NewPet
	ID int64 `json:"id"`
}

type Pets []Pet
