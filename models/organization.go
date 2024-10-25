package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID        uuid.UUID        `db:"id" json:"id"`
	Name      string           `db:"name" json:"name"`
	Type      OrganizationEnum `db:"type" json:"type"`
	CreatedAt time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt time.Time        `db:"updated_at" json:"updated_at"`
}
