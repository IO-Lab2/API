package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func ScientistPublicationByID(db *sqlx.DB, id uuid.UUID) ([]responses.ScientistPublicationBody, error) {
	query := "SELECT id, scientist_id, publication_id, created_at, updated_at FROM scientists_publications WHERE id = $1"
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
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
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		scientistPublications = append(scientistPublications, scientistPublication)
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error iterating over rows:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist-publication relationships by ID")
	return scientistPublications, nil
}
