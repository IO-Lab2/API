package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func MinisterialScoreFilter(db *sqlx.DB) (*models.MinisterialScore, error) {
	query := "SELECT MAX(ministerial_score ) as largest, MIN(ministerial_score ) as smallest FROM ministerial_score"
	var scores models.MinisterialScore
	if err := db.Get(&scores, query); err != nil {
		return nil, err
	}

	return &scores, nil
}
