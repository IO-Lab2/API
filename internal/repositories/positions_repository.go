package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func QueryGetPositions(db *sqlx.DB) ([]models.PositionFilter, error) {

	query := `
		SELECT DISTINCT position
		FROM scientists
		WHERE position IS NOT NULL`
	var result []models.PositionFilter
	if err := db.Select(&result, query); err != nil {
		return nil, err
	}

	return result, nil
}
