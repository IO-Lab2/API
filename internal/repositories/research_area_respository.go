package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func ResearchAreaFilter(db *sqlx.DB) ([]models.ResearchArea, error) {
	query := "SELECT DISTINCT research_area FROM scientists"
	logging.Logger.Info("INFO: Executing query:", query)

	var areas []models.ResearchArea
	if err := db.Select(&areas, query); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved research areas")
	return areas, nil
}
