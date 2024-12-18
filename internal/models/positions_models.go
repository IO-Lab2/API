package models

type PositionFilter struct {
	Position *string `db:"position" json:"position" doc:"Position name"`
}
