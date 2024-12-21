package repositories

import (
	"database/sql"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func ScientistByID(db *sqlx.DB, id uuid.UUID) (*responses.ScientistBody, error) {
	query := `
	SELECT 
		s.id, 
		s.first_name, 
		s.last_name, 
		s.academic_title, 
		s.email, 
		s.profile_url, 
		s.created_at, 
		s.updated_at, 
		ARRAY_AGG(ra.name) AS research_areas
	FROM 
		scientists s
	LEFT JOIN 
		scientists_research_areas sra ON s.id = sra.scientist_id
	LEFT JOIN 
		research_areas ra ON sra.research_area_id = ra.id
	WHERE 
		s.id = $1
	GROUP BY 
		s.id
	`

	logging.Logger.Info("INFO: Executing query:", query)

	var scientist responses.ScientistBody
	var researchAreaNames []string

	// Execute the query and scan the results
	row := db.QueryRow(query, id)
	err := row.Scan(
		&scientist.ID,
		&scientist.FirstName,
		&scientist.LastName,
		&scientist.AcademicTitle,
		&scientist.Email,
		&scientist.ProfileUrl,
		&scientist.CreatedAt,
		&scientist.UpdatedAt,
		pq.Array(&researchAreaNames), // Use pq.Array to scan the ARRAY_AGG result
	)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	// Map research area names to ResearchArea structs
	for _, name := range researchAreaNames {
		scientist.ResearchAreas = append(scientist.ResearchAreas, responses.ResearchArea{Name: name})
	}

	// Fetch bibliometrics information
	bibliometricsQuery := `
	SELECT 
		COALESCE(h_index_wos, 0), 
		COALESCE(h_index_scopus, 0), 
		COALESCE(publication_count, 0), 
		COALESCE(ministerial_score, 0)
	FROM 
		bibliometrics
	WHERE 
		scientist_id = $1
	`

	err = db.QueryRow(bibliometricsQuery, id).Scan(
		&scientist.Bibliometrics.HIndexWOS,
		&scientist.Bibliometrics.HIndexScopus,
		&scientist.Bibliometrics.PublicationCount,
		&scientist.Bibliometrics.MinisterialScore,
	)
	if err != nil && err != sql.ErrNoRows {
		logging.Logger.Error("ERROR: Error fetching bibliometrics:", err)
		return nil, err
	}

	// Fetch publication scores grouped by year
	publicationScoresQuery := `
	SELECT 
		EXTRACT(YEAR FROM p.publication_date) AS year, 
		SUM(p.ministerial_score) AS total_score
	FROM 
		publications p
	JOIN 
		scientists_publications sp ON sp.publication_id = p.id
	WHERE 
		sp.scientist_id = $1
	GROUP BY 
		EXTRACT(YEAR FROM p.publication_date)
	`

	rows, err := db.Query(publicationScoresQuery, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error fetching publication scores:", err)
		return nil, err
	}
	defer rows.Close()

	var publicationScores []responses.PublicationScore

	for rows.Next() {
		var year *string
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

	scientist.PublicationScores = publicationScores

	logging.Logger.Info("INFO: Successfully retrieved scientist by ID")
	return &scientist, nil
}

func ScientistByName(db *sqlx.DB, name *requests.ScientistName) (*responses.ScientistBody, error) {
	query := `
	SELECT 
		s.id, 
		s.first_name, 
		s.last_name, 
		s.academic_title, 
		s.email, 
		s.profile_url, 
		s.created_at, 
		s.updated_at, 
		ARRAY_AGG(ra.name) AS research_areas
	FROM 
		scientists s
	LEFT JOIN 
		scientists_research_areas sra ON s.id = sra.scientist_id
	LEFT JOIN 
		research_areas ra ON sra.research_area_id = ra.id
	WHERE 
		s.first_name = $1 AND s.last_name = $2
	GROUP BY 
		s.id
	`
	logging.Logger.Info("INFO: Executing query:", query)

	var scientist responses.ScientistBody
	var researchAreaNames []string

	// Execute the query and scan the results
	row := db.QueryRow(query, name.FirstName, name.LastName)
	err := row.Scan(
		&scientist.ID,
		&scientist.FirstName,
		&scientist.LastName,
		&scientist.AcademicTitle,
		&scientist.Email,
		&scientist.ProfileUrl,
		&scientist.CreatedAt,
		&scientist.UpdatedAt,
		pq.Array(&researchAreaNames), // Use pq.Array to scan the ARRAY_AGG result
	)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	// Map research area names to ResearchArea structs
	for _, name := range researchAreaNames {
		scientist.ResearchAreas = append(scientist.ResearchAreas, responses.ResearchArea{Name: name})
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist by name")
	return &scientist, nil
}
