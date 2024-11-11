package responses

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationsResponse struct {
	Body []OrganizationBody `json:"body" doc:"Organization object"`
}
type OrganizationBody struct {
	ID        uuid.UUID `json:"id" doc:"Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Name      string    `json:"name" doc:"Name of the organization" format:"string" example:"Politechnika Warszawska"`
	Type      string    `json:"type" doc:"Type of the organization" format:"string" example:"Uniwersytet"`
	CreatedAt time.Time `json:"createdAt" doc:"Creation date of the organization" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"UpdatedAt" doc:"Last update date of the organization" format:"date-time" example:"2021-01-01T00:00:00Z"`
}
