package test

import "github.com/vkd/goag/tests/post_custom_type/pkg"

// ------------------------
//         Schemas
// ------------------------

type NewPet struct {
	Name string     `json:"name"`
	Tag  pkg.PetTag `json:"tag"`
}
