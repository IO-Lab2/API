package models

import (
	"time"

	"github.com/google/uuid"
)

type Scientist struct {
	ID            uuid.UUID      `db:"id" json:"id"`
	FirstName     string         `db:"first_name" json:"first_name"`
	LastName      string         `db:"last_name" json:"last_name"`
	AcademicTitle string         `db:"academic_title" json:"academic_title"`
	Publication   *string        `db:"publication" json:"publication"`
	ResearchAreas []ResearchArea `db:"research_areas" json:"research_areas" doc:"Research areas of the scientist"`
	Email         string         `db:"email" json:"email"`
	ProfileUrl    string         `db:"profile_url" json:"profile_url"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
}

type ResearchArea struct {
	Name string `db:"name" json:"name" doc:"Name of the research area" format:"string" example:"health sciences (HS)"`
}
type CreateScientist struct {
	FirstName     string  `db:"first_name" json:"first_name"`
	LastName      string  `db:"last_name" json:"last_name"`
	AcademicTitle string  `db:"academic_title" json:"academic_title"`
	Position      *string `db:"position" json:"position"`
	ResearchArea  string  `db:"research_area" json:"research_area"`
	Email         string  `db:"email" json:"email"`
	ProfileUrl    string  `db:"profile_url" json:"profile_url"`
}
type UpdateScientist struct {
	ID            uuid.UUID `db:"id" json:"id"`
	FirstName     string    `db:"first_name" json:"first_name"`
	LastName      string    `db:"last_name" json:"last_name"`
	AcademicTitle string    `db:"academic_title" json:"academic_title"`
	Position      *string   `db:"position" json:"position"`
	ResearchArea  string    `db:"research_area" json:"research_area"`
	Email         string    `db:"email" json:"email"`
	ProfileUrl    string    `db:"profile_url" json:"profile_url"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
type DeleteScientist struct {
	ID uuid.UUID `db:"id" json:"id"`
}
