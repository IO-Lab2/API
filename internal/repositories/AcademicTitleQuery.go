package repositories

import (
	"database/sql"
	"io-project-api/internal/models"
)

func GetAcademicTitleFilter(db *sql.DB) ([]models.AcademicTitle, error) {
	query := ""
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var titles []models.AcademicTitle
	for rows.Next() {
		var title models.AcademicTitle
		if err := rows.Scan(&title.Title); err != nil {
			return nil, err
		}
		titles = append(titles, title)
	}
	return titles, nil
}
