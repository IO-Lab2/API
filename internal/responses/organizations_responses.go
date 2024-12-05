package responses

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationResponse struct {
	Body *OrganizationBodyExtended `json:"body" doc:"Organization object"`
}

type ListOfOrganizations struct {
	Body []OrganizationBody `json:"body" doc:"Organization object"`
}

type OrganizationBody struct {
	ID   uuid.UUID `db:"id" json:"id" doc:"Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Name string    `db:"name" json:"name" doc:"Name of the organization" format:"string" example:"Politechnika Warszawska"`
	Type string    `db:"type" json:"type" doc:"Type of the organization" format:"string" example:"Uniwersytet"`
}

type ListOfOrganizationsResponse struct {
	Body []OrganizationBodyExtended `json:"body" doc:"Organization object"`
}
type OrganizationBodyExtended struct {
	ID        uuid.UUID `db:"id" json:"id" doc:"Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Name      string    `db:"name" json:"name" doc:"Name of the organization" format:"string" example:"Politechnika Warszawska"`
	Type      string    `db:"type" json:"type" doc:"Type of the organization" format:"string" example:"Uniwersytet"`
	CreatedAt time.Time `db:"created_at" json:"created_at" doc:"Creation date of the organization" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" doc:"Last update date of the organization" format:"date-time" example:"2021-01-01T00:00:00Z"`
}
type CreateOrganization struct {
	ID uuid.UUID `db:"id" json:"id" doc:"ID of created organization" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
type CreateOrganizationResponse struct {
	Body CreateOrganization `json:"body" doc:"Organization creation object"`
}
type UpdateOrganizationResponse struct {
	Body OrganizationBodyExtended `json:"body" doc:"Updated organization"`
}
