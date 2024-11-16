package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func GetPublicationCountFilter(db *sqlx.DB) ([]models.PublicationCount, error) {
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
