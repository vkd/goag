package test

import (
	"context"
	"fmt"
	"net/http"
)

var db interface {
	GetPet(_ context.Context, id string) (Pet, error)
}

func ExampleAPI_petsStore() {
	api := &API{
		GetPetsPetIDHandler: func(r GetPetsPetIDRequest) GetPetsPetIDResponse {
			req, err := r.Parse()
			if err != nil {
				return NewGetPetsPetIDResponseDefaultJSON(http.StatusBadRequest, Error{
					Code:    400,
					Message: fmt.Sprintf("Bad request: %v", err),
				})
			}

			out, err := db.GetPet(r.HTTP().Context(), req.Path.PetID)
			if err != nil {
				return NewGetPetsPetIDResponseDefaultJSON(http.StatusInternalServerError, Error{
					Code:    500,
					Message: fmt.Sprintf("Internal server error: %v", err),
				})
			}

			return NewGetPetsPetIDResponse200JSON(out)
		},
		// ...
	}

	_ = http.ListenAndServe(":8080", api)
}
