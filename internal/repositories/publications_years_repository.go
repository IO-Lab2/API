package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func QueryGetPublicationsYears(db *sqlx.DB) ([]models.PublicationsYear, error) {
	query := `
		SELECT DISTINCT publication_date
		FROM publications
		WHERE publication_date IS NOT NULL`
	var result []models.PublicationsYear
	if err := db.Select(&result, query); err != nil {
		return nil, err
	}

	return result, nil
}
