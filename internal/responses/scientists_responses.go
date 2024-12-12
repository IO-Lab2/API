package responses

import (
	"time"

	"github.com/google/uuid"
)

type ScientistsResponse struct {
	Body []ScientistBody `json:"body" doc:"Scientists object"`
}

type ScientistResponse struct {
	Body *ScientistBody `json:"body" doc:"Scientist object"`
}

type ScientistBody struct {
	ID            uuid.UUID      `db:"id" json:"id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	FirstName     string         `db:"first_name" json:"first_name" doc:"First name of the scientist" format:"string" example:"John"`
	LastName      string         `db:"last_name" json:"last_name" doc:"Last name of the scientist" format:"string" example:"Doe"`
	AcademicTitle string         `db:"academic_title" json:"academic_title" doc:"Academic title of the scientist" format:"string" example:"PhD"`
	ResearchAreas []ResearchArea `db:"research_areas" json:"research_areas" doc:"Research areas of the scientist"`
	Email         *string        `db:"email,omitempty" json:"email" doc:"Email of the scientist" format:"string" example:"example@example.com"`
	ProfileUrl    *string        `db:"profile_url,omitempty" json:"profile_url" doc:"Profile URL of the scientist" format:"hostname" example:"https://example.com"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at" doc:"Creation date of the scientist" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at" doc:"Last update date of the scientist" format:"date-time" example:"2021-01-01T00:00:00Z"`
}

type ResearchArea struct {
	Name string `db:"name" json:"name" doc:"Name of the research area" format:"string" example:"health sciences (HS)"`
}

type ResearchAreaExtended struct {
	ID   uuid.UUID `db:"id" json:"id" doc:"Research area ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Name string    `db:"name" json:"name" doc:"Name of the research area" format:"string" example:"health sciences (HS)"`
}
