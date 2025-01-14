package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func PublicationByID(db *sqlx.DB, id uuid.UUID) (*responses.PublicationBody, error) {
	query := "SELECT id, title, journal, publication_date, journal_impact_factor, journal_type, ministerial_score, created_at, updated_at FROM publications WHERE id = $1"
	logging.Logger.Info("INFO: Executing query:", query)

	var publication responses.PublicationBody
	if err := db.Get(&publication, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved publication by ID")
	return &publication, nil
}

func RandomPublication(db *sqlx.DB) (*responses.PublicationBody, error) {
	query := "SELECT id, title, journal, publication_date, journal_impact_factor, journal_type, ministerial_score, created_at, updated_at FROM publications ORDER BY RANDOM() LIMIT 1"
	logging.Logger.Info("INFO: Executing query:", query)

	var publication responses.PublicationBody
	if err := db.Get(&publication, query); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved random publication")
	return &publication, nil
}

func PublicationsByScientistID(db *sqlx.DB, id uuid.UUID) ([]responses.PublicationBody, error) {
	query := `
		SELECT p.id, p.title, p.journal, p.publication_date, p.journal_impact_factor, p.journal_type, p.ministerial_score,
		p.created_at, p.updated_at
		FROM publications p
		JOIN scientist_publication sp ON p.id = sp.publication_id
		WHERE sp.scientist_id = $1`
	logging.Logger.Info("INFO: Executing query:", query)

	var publications []responses.PublicationBody
	if err := db.Select(&publications, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved publications by scientist ID")
	return publications, nil
}

func PublicationCountFilter(db *sqlx.DB) (*models.PublicationCount, error) {
	query := "SELECT MAX(publication_count) as largest, MIN(publication_count) as smallest FROM bibliometrics"
	logging.Logger.Info("INFO: Executing query:", query)

	var publicationCount models.PublicationCount
	if err := db.Get(&publicationCount, query); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved publication count filter")
	return &publicationCount, nil
}
func CreatePublication(db *sqlx.DB, id uuid.UUID, publication *requests.CreatePublicationRequest) error {
	query := `
		INSERT INTO publications (id, title, journal, publication_date, journal_impact_factor, publisher, journal_type, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
	`
	logging.Logger.Info("INFO: Executing query:", query)
	_, err := db.Exec(query, id, publication.Title, publication.Journal, publication.PublicationDate, publication.JournalImpactFactor, publication.Publisher, publication.JournalType)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return err
	}
	logging.Logger.Info("INFO: Successfully created publication")
	return nil
}
func DeletePublication(db *sqlx.DB, input *requests.DeletePublication) error {
	query := `
	DELETE FROM  publications WHERE id = $1
	`
	logging.Logger.Info("INFO: Executing query: ", query)
	_, err := db.Exec(query, input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to execute query.", err)
	}
	return nil
}
