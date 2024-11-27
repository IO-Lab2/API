package repositories

import (
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// BibliometricByID retrieves a bibliometric entry by its ID.
func BibliometricByID(db *sqlx.DB, id uuid.UUID) (*responses.BibliometricBody, error) {
	query := "SELECT id, h_index_wos, h_index_scopus, publication_count, ministerial_score, scientist_id, created_at, updated_at FROM bibliometrics WHERE id = $1"
	var bibliometric responses.BibliometricBody
	if err := db.Get(&bibliometric, query, id); err != nil {
		return nil, err
	}

	return &bibliometric, nil
}

// BibliometricByAuthor retrieves bibliometric entries by the scientist's ID.
func BibliometricByAuthor(db *sqlx.DB, id uuid.UUID) ([]responses.BibliometricBody, error) {
	query := "SELECT id, h_index_wos, h_index_scopus, publication_count, ministerial_score, scientist_id, created_at, updated_at FROM bibliometrics WHERE scientist_id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bibliometrics []responses.BibliometricBody

	for rows.Next() {
		var bibliometric responses.BibliometricBody
		err := rows.Scan(
			&bibliometric.ID,
			&bibliometric.HIndexWos,
			&bibliometric.HIndexScopus,
			&bibliometric.PublicationCount,
			&bibliometric.MinisterialScore,
			&bibliometric.ScientistID,
			&bibliometric.CreatedAt,
			&bibliometric.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bibliometrics = append(bibliometrics, bibliometric)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bibliometrics, nil
}
