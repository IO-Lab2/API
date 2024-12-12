package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"

	"github.com/jmoiron/sqlx"
)

func ResearchAreaFilter(db *sqlx.DB) ([]responses.ResearchAreaExtended, error) {
	query := "SELECT DISTINCT id, name FROM research_areas"
	logging.Logger.Info("INFO: Executing query:", query)

	var areas []responses.ResearchAreaExtended
	if err := db.Select(&areas, query); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved research areas")
	return areas, nil
}
