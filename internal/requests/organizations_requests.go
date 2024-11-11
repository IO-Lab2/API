package requests

import "github.com/google/uuid"

type OrganizationID struct {
	ID uuid.UUID `path:"id" doc:"Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
