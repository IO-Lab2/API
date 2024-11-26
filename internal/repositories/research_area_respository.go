package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func ResearchAreaFilter(db *sqlx.DB) ([]models.ResearchTitle, error) {
	query := "SELECT DISTINCT research_area FROM scientists"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var titles []models.ResearchTitle
	for rows.Next() {
		var title models.ResearchTitle
		if err := rows.Scan(&title.Title); err != nil {
			return nil, err
		}
		titles = append(titles, title)
	}
	return titles, nil
}
