package requests

import "github.com/google/uuid"

type ScientistID struct {
	ID uuid.UUID `path:"id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}

type ScientistName struct {
	FirstName string `db:"first_name" query:"first_name" doc:"Scientist first name" example:"John"`
	LastName  string `db:"last_name" query:"last_name" doc:"Scientist last name" example:"Doe"`
}
