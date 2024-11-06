package repositories

import (
	"database/sql"
	"io-project-api/models"

	"github.com/google/uuid"
)

func ScientistPublicationByID(db *sql.DB, id uuid.UUID) ([]models.ScientistPublication, error) {
	query := "SELECT id, scientist_id, publication_id, created_at, updated_at FROM scientists_publications WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scientistPublications []models.ScientistPublication

	for rows.Next() {
		var scientistPublication models.ScientistPublication
		err := rows.Scan(
			&scientistPublication.ID,
			&scientistPublication.ScientistID,
			&scientistPublication.PublicationID,
			&scientistPublication.CreatedAt,
			&scientistPublication.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		scientistPublications = append(scientistPublications, scientistPublication)
	}
	return scientistPublications, nil
}
