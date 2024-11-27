package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

// Deprecated
func CitationCountFilter(db *sqlx.DB) (*models.CitationsFilter, error) {
	query := "SELECT MAX(citation_count) as largest, MIN(citation_count) as smallest FROM bibliometrics"
	var citations models.CitationsFilter
	if err := db.QueryRow(query).Scan(&citations.Largest, &citations.Smallest); err != nil {
		return nil, err
	}

	return &citations, nil
}
