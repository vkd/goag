package test

// ------------------------
//         Schemas
// ------------------------

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type Pet struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Pets []Pet
