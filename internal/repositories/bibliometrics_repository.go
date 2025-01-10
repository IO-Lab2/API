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
	SELECT 
		id, 
		h_index_wos, 
		h_index_scopus, 
		publication_count, 
		ministerial_score, 
		scientist_id, 
		created_at, 
		updated_at 
	FROM 
		bibliometrics 
	WHERE 
		id = $1`

	logging.Logger.Info("INFO: Executing query:", query)

	var bibliometric responses.BibliometricBody
	if err := db.Get(&bibliometric, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	metricsQuery := `
	SELECT 
		EXTRACT(YEAR FROM created_at) AS year, 
		SUM(ministerial_score) AS score 
	FROM 
		bibliometrics 
	WHERE 
		id = $1 
	GROUP BY 
		EXTRACT(YEAR FROM created_at)`

	logging.Logger.Info("INFO: Executing metrics query:", metricsQuery)

	rows, err := db.Query(metricsQuery, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing metrics query:", err)
		return nil, err
	}
	defer rows.Close()

	var publicationScores []responses.PublicationScore

	for rows.Next() {
		var year *int
		var score *float64
		if err := rows.Scan(&year, &score); err != nil {
			logging.Logger.Error("ERROR: Error scanning publication scores:", err)
			return nil, err
		}
		publicationScores = append(publicationScores, responses.PublicationScore{Year: year, Score: score})
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Row iteration error:", err)
		return nil, err
	}

	bibliometric.PublicationScores = publicationScores

	logging.Logger.Info("INFO: Successfully retrieved bibliometric by ID")
	return &bibliometric, nil
}

func BibliometricByScientistID(db *sqlx.DB, id uuid.UUID) (*responses.BibliometricBody, error) {
	query := `
	SELECT 
		id, 
		h_index_wos, 
		h_index_scopus, 
		publication_count, 
		ministerial_score, 
		scientist_id, 
		created_at, 
		updated_at 
	FROM 
		bibliometrics 
	WHERE 
		scientist_id = $1 
	ORDER BY 
		created_at DESC 
	LIMIT 1`

	logging.Logger.Info("INFO: Executing query:", query)

	var bibliometric responses.BibliometricBody
	if err := db.Get(&bibliometric, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	// Fetch publication scores grouped by year
	publicationScoresQuery := `
	SELECT 
		COALESCE(EXTRACT(YEAR FROM p.publication_date)::TEXT, 'null') AS year, 
		SUM(COALESCE(p.ministerial_score, 0)) AS total_score
	FROM 
		publications p
	JOIN 
		scientists_publications sp ON sp.publication_id = p.id
	WHERE 
		sp.scientist_id = $1
	GROUP BY 
		EXTRACT(YEAR FROM p.publication_date)
	ORDER BY 
		EXTRACT(YEAR FROM p.publication_date) ASC
	`

	logging.Logger.Info("INFO: Executing publication scores query:", publicationScoresQuery)

	rows, err := db.Query(publicationScoresQuery, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error fetching publication scores:", err)
		return nil, err
	}
	defer rows.Close()

	var publicationScores []responses.PublicationScore

	for rows.Next() {
		var year *int
		var score *float64
		if err := rows.Scan(&year, &score); err != nil {
			logging.Logger.Error("ERROR: Error scanning publication scores:", err)
			return nil, err
		}
		publicationScores = append(publicationScores, responses.PublicationScore{Year: year, Score: score})
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Row iteration error:", err)
		return nil, err
	}

	bibliometric.PublicationScores = publicationScores

	logging.Logger.Info("INFO: Successfully retrieved bibliometrics by scientist ID")
	return &bibliometric, nil
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
