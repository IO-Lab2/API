package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func ImpactFactorFilter(db *sqlx.DB) (*models.RangeFilter, error) {
	query := "SELECT MAX(impact_factor) as largest, MIN(impact_factor) as smallest FROM bibliometrics"
	logging.Logger.Info("INFO: Executing query:", query)
	var scores models.RangeFilter
	if err := db.Get(&scores, query); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved impact factor filter")
	return &scores, nil
}
