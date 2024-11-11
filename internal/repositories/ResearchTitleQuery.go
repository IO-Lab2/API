package repositories

import (
	"database/sql"
	"io-project-api/internal/models"
)

func GetResearchTitleFilter(db *sql.DB) ([]models.ResearchTitle, error) {
	query := ""
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
