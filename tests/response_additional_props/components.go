package test

// ------------------------
//         Schemas
// ------------------------

type Pet struct {
	Name string `json:"name"`
}

type Pets []Pet
