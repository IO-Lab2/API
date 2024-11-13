package requests

import "github.com/google/uuid"

type ScientistOrganizationID struct {
	ID uuid.UUID `path:"ID" doc:"Scientist Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
