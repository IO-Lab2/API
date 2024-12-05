package models

import "github.com/google/uuid"

type ResearchArea struct {
	ID   uuid.UUID `db:"id" json:"id" doc:"Research area ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Area string    `db:"research_area" json:"research_area" doc:"Research area" format:"string" example:"Artificial Intelligence"`
}
