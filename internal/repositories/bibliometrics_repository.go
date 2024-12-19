package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func BibliometricByID(db *sqlx.DB, id uuid.UUID) (*responses.BibliometricBody, error) {
	query := `
			SELECT id, h_index_wos, h_index_scopus, publication_count, ministerial_score, scientist_id, created_at, updated_at 
			FROM bibliometrics 
			WHERE id = $1`

	logging.Logger.Info("INFO: Executing query:", query)
	var bibliometric responses.BibliometricBody
	if err := db.Get(&bibliometric, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved bibliometric by ID")
	return &bibliometric, nil
}

func BibliometricByScientistID(db *sqlx.DB, id uuid.UUID) (*responses.BibliometricBody, error) {
	query := `
		SELECT id, h_index_wos, h_index_scopus, publication_count, ministerial_score, scientist_id, created_at, updated_at 
		FROM bibliometrics 
		WHERE scientist_id = $1
		ORDER BY id DESC
		LIMIT 1`
	logging.Logger.Info("INFO: Executing query:", query)
	var result responses.BibliometricBody
	if err := db.Get(&result, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved bibliometrics by author ID")
	return &result, nil
}

func CreateBibliometric(db *sqlx.DB, id uuid.UUID, input *requests.CreateBibliometric) error {
	query := `
        INSERT INTO bibliometrics (id, h_index_wos, h_index_scopus, publication_count, ministerial_score, scientist_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())`
	logging.Logger.Info("INFO: Executing query:", query)
	_, err := db.Exec(query, id, input.HIndexWos, input.HIndexScopus, input.PublicationCount, input.MinisterialScore, input.ScientistID)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
	}
	return err
}

func DeleteBibliometric(db *sqlx.DB, input *requests.DeleteBibliometric) error {
	query := "DELETE FROM bibliometrics WHERE id = $1"
	logging.Logger.Info("INFO: Executing query:", query)
	_, err := db.Exec(query, input.ID)
	if err != nil {
		logging.Logger.Error("ERROR:  Error executing query:", err)
	}
	return err
}

func UpdateBibliometric(db *sqlx.DB, input *requests.UpdateBibliometric) error {
	query := `
        UPDATE bibliometrics
        SET h_index_wos = $2, h_index_scopus = $3, publication_count = $4, ministerial_score = $5, updated_at = NOW()
        WHERE id = $1`
	logging.Logger.Info("INFO: Executing query:", query)
	_, err := db.Exec(query, input.ID, input.HIndexWos, input.HIndexScopus, input.PublicationCount, input.MinisterialScore)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
	}
	return err
}
