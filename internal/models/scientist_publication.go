package models

import (
	"time"

	"github.com/google/uuid"
)

type ScientistPublication struct {
	ID            uuid.UUID `db:"id" json:"id"`
	ScientistID   uuid.UUID `db:"scientist_id" json:"scientist_id"`
	PublicationID uuid.UUID `db:"publication_id" json:"publication_id"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
type CreateScientistPublication struct {
	ScientistID   uuid.UUID `db:"scientist_id" json:"scientist_id"`
	PublicationID uuid.UUID `db:"publication_id" json:"publication_id"`
}
type UpdateScientistPublication struct {
	ID            uuid.UUID `db:"id" json:"id"`
	ScientistID   uuid.UUID `db:"scientist_id" json:"scientist_id"`
	PublicationID uuid.UUID `db:"publication_id" json:"publication_id"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
type DeleteScientistPublication struct {
	ID uuid.UUID `db:"id" json:"id"`
}
