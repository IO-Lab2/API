package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func ResearchAreaFilter(db *sqlx.DB) ([]models.ResearchArea, error) {
	query := "SELECT DISTINCT id, research_area FROM scientists"
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var areas []models.ResearchArea
	for rows.Next() {
		var area models.ResearchArea
		if err := rows.Scan(
			&area.Area,
			&area.ID); err != nil {
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		areas = append(areas, area)
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error iterating over rows:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved research areas")
	return areas, nil
}
