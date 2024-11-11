package repositories

import (
	"database/sql"
	"io-project-api/internal/models"
)

func GetPublicationCountFilter(db *sql.DB) ([]models.PublicationCount, error) {
	query := ""
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var counts []models.PublicationCount
	for rows.Next() {
		var count models.PublicationCount
		if err := rows.Scan(&count.Count); err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}
