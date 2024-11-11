package repositories

import (
	"database/sql"
	"io-project-api/internal/models"
)

func GetCitationCountFilter(db *sql.DB) ([]models.CitationCount, error) {
	query := ""
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var counts []models.CitationCount
	for rows.Next() {
		var count models.CitationCount
		if err := rows.Scan(&count.Count); err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}
