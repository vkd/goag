package test

import "time"

// ------------------------
//         Schemas
// ------------------------

type Pet struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
}
