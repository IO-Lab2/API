package models

import (
	"database/sql"
	"io-project-api/models"

	"github.com/google/uuid"
)

func BibliometricsByID(db *sql.DB, id uuid.UUID) ([]models.Bibliometrics, error) {
	query := "SELECT id, h_index, citation_count, publication_count, ministerial_score, scientist_id FROM bibliometrics WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var Bibliometrics []models.Bibliometrics

	for rows.Next() {
		var Bibliometric models.Bibliometrics
		err := rows.Scan(
			&Bibliometric.ID,
			&Bibliometric.HIndex,
			&Bibliometric.CitationCount,
			&Bibliometric.PublicationCount,
			&Bibliometric.MinisterialScore,
			&Bibliometric.ScientistID,
			&Bibliometric.CreatedAt,
			&Bibliometric.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		Bibliometrics = append(Bibliometrics, Bibliometric)
	}
	return Bibliometrics, nil
}
