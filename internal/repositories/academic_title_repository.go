package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func AcademicTitleFilter(db *sqlx.DB) ([]models.AcademicTitle, error) {
	query := `SELECT DISTINCT academic_title FROM scientists`
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var titles []models.AcademicTitle
	for rows.Next() {
		var title models.AcademicTitle
		if err := rows.Scan(&title.Title); err != nil {
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		titles = append(titles, title)
	}

	logging.Logger.Info("INFO: Successfully retrieved academic titles")
	return titles, nil
}
