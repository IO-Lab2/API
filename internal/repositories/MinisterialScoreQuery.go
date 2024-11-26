package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func MinisterialScoreFilter(db *sqlx.DB) ([]models.MinisterialScore, error) {
	query := ""
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []models.MinisterialScore
	for rows.Next() {
		var score models.MinisterialScore
		if err := rows.Scan(&score.Score); err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}
