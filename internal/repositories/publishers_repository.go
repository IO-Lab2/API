package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func QueryGetPublishers(db *sqlx.DB) ([]models.PublisherFilter, error) {
	query := `
		SELECT DISTINCT publisher
		FROM publications
		WHERE publisher IS NOT NULL`
	var result []models.PublisherFilter
	if err := db.Select(&result, query); err != nil {
		return nil, err
	}

	return result, nil
}
