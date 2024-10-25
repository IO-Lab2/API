package models

import (
	"time"

	"github.com/google/uuid"
)

type ScientistOrganization struct {
	ID             uuid.UUID `db:"id" json:"id"`
	ScientistID    uuid.UUID `db:"scientist_id" json:"scientist_id"`
	OrganizationID uuid.UUID `db:"organization_id" json:"organization_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
