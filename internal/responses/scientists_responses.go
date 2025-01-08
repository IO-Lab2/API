package responses

import (
	"time"

	"github.com/google/uuid"
)

type ScientistsResponse struct {
	Body *ScientistsResponseBody `json:"body" doc:"List of scientists"`
}

type ScientistsResponseBody struct {
	Scientists []ScientistBody `json:"scientists" doc:"List of scientists"`
	Count      int             `json:"count" doc:"Total number of scientists"`
}

type ScientistResponse struct {
	Body *ScientistBody `json:"body" doc:"Scientist object"`
}

type ScientistBody struct {
	ID                uuid.UUID          `db:"id" json:"id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	FirstName         string             `db:"first_name" json:"first_name" doc:"First name of the scientist" format:"string" example:"John"`
	LastName          string             `db:"last_name" json:"last_name" doc:"Last name of the scientist" format:"string" example:"Doe"`
	AcademicTitle     string             `db:"academic_title" json:"academic_title" doc:"Academic title of the scientist" format:"string" example:"PhD"`
	Position          *string            `db:"position" json:"position,omitempty" doc:"Position of the scientist" format:"string" example:"Researcher"`
	ResearchAreas     []ResearchArea     `db:"research_areas" json:"research_areas" doc:"Research areas of the scientist"`
	Email             *string            `db:"email,omitempty" json:"email" doc:"Email of the scientist" format:"string" example:"example@example.com"`
	ProfileUrl        *string            `db:"profile_url,omitempty" json:"profile_url" doc:"Profile URL of the scientist" format:"hostname" example:"https://example.com"`
	CreatedAt         time.Time          `db:"created_at" json:"created_at" doc:"Creation date of the scientist" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt         time.Time          `db:"updated_at" json:"updated_at" doc:"Last update date of the scientist" format:"date-time" example:"2021-01-01T00:00:00Z"`
	Bibliometrics     Bibliometrics      `json:"bibliometrics" doc:"Bibliometric indicators of the scientist"`
	PublicationScores []PublicationScore `json:"publication_scores" doc:"Ministerial score points grouped by year"`
}

type PublicationScore struct {
	Year  *string  `json:"year" doc:"Year of the publication score" format:"string" example:"2021"`
	Score *float64 `json:"score" doc:"Total score for the publications" format:"float" example:"1.0"`
}

type Bibliometrics struct {
	HIndexWOS        *int     `db:"h_index_wos" json:"h_index_wos,omitempty"`
	HIndexScopus     *int     `db:"h_index_scopus" json:"h_index_scopus,omitempty"`
	PublicationCount *int     `db:"publication_count" json:"publication_count"`
	MinisterialScore *float64 `db:"ministerial_score" json:"ministerial_score" doc:"Total ministerial score of the scientist" format:"float" example:"1.0"`
}

type ResearchArea struct {
	Name string `db:"name" json:"name" doc:"Name of the research area" format:"string" example:"health sciences (HS)"`
}

type ResearchAreaExtended struct {
	ID   uuid.UUID `db:"id" json:"id" doc:"Research area ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Name string    `db:"name" json:"name" doc:"Name of the research area" format:"string" example:"health sciences (HS)"`
}
