package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func MinisterialScoreFilter(db *sqlx.DB) (*models.MinisterialScore, error) {
	query := "SELECT MAX(ministerial_score ) as largest, MIN(ministerial_score ) as smallest FROM ministerial_score"
	logging.Logger.Info("INFO: Executing query:", query)
	var scores models.MinisterialScore
	if err := db.Get(&scores, query); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved ministerial score filter")
	return &scores, nil
}
