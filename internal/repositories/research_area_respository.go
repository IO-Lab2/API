package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func ResearchAreaFilter(db *sqlx.DB) ([]models.ResearchArea, error) {
	query := "SELECT DISTINCT research_area FROM scientists"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var areas []models.ResearchArea
	for rows.Next() {
		var area models.ResearchArea
		if err := rows.Scan(&area.Area); err != nil {
			return nil, err
		}
		areas = append(areas, area)
	}
	return areas, nil
}
