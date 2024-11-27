package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func AcademicTitleFilter(db *sqlx.DB) ([]models.AcademicTitle, error) {
	query := "SELECT DISTINCT academic_title FROM scientists"
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
