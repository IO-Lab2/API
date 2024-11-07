package responses

import (
	"time"

	"github.com/google/uuid"
)

type ScientistsResponse struct {
	Body []ScientistBody `json:"body" doc:"Scientist object"`
}

type ScientistBody struct {
	ID            uuid.UUID `json:"id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	FirstName     string    `json:"first_name" doc:"First name of the scientist" format:"string" example:"John"`
	LastName      string    `json:"last_name" doc:"Last name of the scientist" format:"string" example:"Doe"`
	AcademicTitle string    `json:"academic_title" doc:"Academic title of the scientist" format:"string" example:"PhD"`
	ResearchArea  string    `json:"research_area" doc:"Research area of the scientist" format:"string" example:"Computer Science"`
	Email         string    `json:"email" doc:"Email of the scientist" format:"string" example:"example@example.com"`
	ProfileUrl    string    `json:"profile_url" doc:"Profile URL of the scientist" format:"hostname" example:"https://example.com"`
	CreatedAt     time.Time `json:"created_at" doc:"Creation date of the scientist" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt     time.Time `json:"updated_at" doc:"Last update date of the scientist" format:"date-time" example:"2021-01-01T00:00:00Z"`
}
