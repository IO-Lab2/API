package repositories

import (
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func ScientistPublicationByID(db *sqlx.DB, id uuid.UUID) ([]responses.ScientistPublicationBody, error) {
	query := "SELECT id, scientist_id, publication_id, created_at, updated_at FROM scientists_publications WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scientistPublications []responses.ScientistPublicationBody

	for rows.Next() {
		var scientistPublication responses.ScientistPublicationBody
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
