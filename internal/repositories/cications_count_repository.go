package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

// Deprecated
func CitationCountFilter(db *sqlx.DB) (*models.CitationsFilter, error) {
	query := "SELECT MAX(citation_count) as largest, MIN(citation_count) as smallest FROM bibliometrics"
	logging.Logger.Info("INFO: Executing query:", query)
	var citations models.CitationsFilter
	if err := db.QueryRow(query).Scan(&citations.Largest, &citations.Smallest); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved citation count filter")
	return &citations, nil
}
