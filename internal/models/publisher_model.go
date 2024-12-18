package models

type PublisherFilter struct {
	Publisher *string `db:"publisher" json:"publisher" doc:"Publisher name"`
}
