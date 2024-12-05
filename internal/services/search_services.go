package services

import (
	"fmt"
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"
	"strings"
)

func SearchForScientists(input *models.SearchInput) ([]responses.ScientistBody, error) {
	query := `
		SELECT 
			s.id, 
			s.first_name, 
			s.last_name, 
			s.academic_title, 
			s.research_area, 
			s.email, 
			s.profile_url, 
			s.created_at, 
			s.updated_at 
		FROM scientists s
		LEFT JOIN bibliometrics b ON s.id = b.scientist_id
	`

	// Where clauses
	whereClauses := []string{}
	args := map[string]interface{}{}

	// Filters
	if input.Name != "" {
		whereClauses = append(whereClauses, "s.first_name ILIKE :name")
		args["name"] = "%" + input.Name + "%"
	}
	if input.Surname != "" {
		whereClauses = append(whereClauses, "s.last_name ILIKE :surname")
		args["surname"] = "%" + input.Surname + "%"
	}
	if input.Body != nil {
		if input.Body.AcademicTitles != nil {
			whereClauses = append(whereClauses, "s.academic_title = ANY(:titles)")
			args["titles"] = input.Body.AcademicTitles.Titles
		}
		if input.Body.PublicationsCounts != nil {
			if input.Body.PublicationsCounts.LowerBound > 0 {
				whereClauses = append(whereClauses, "b.publication_count >= :publications_lower")
				args["publications_lower"] = input.Body.PublicationsCounts.LowerBound
			}
			if input.Body.PublicationsCounts.UpperBound > 0 {
				whereClauses = append(whereClauses, "b.publication_count <= :publications_upper")
				args["publications_upper"] = input.Body.PublicationsCounts.UpperBound
			}
		}
		if input.Body.MinisterialScores != nil {
			if input.Body.MinisterialScores.LowerBound > 0 {
				whereClauses = append(whereClauses, "b.ministerial_score >= :score_lower")
				args["score_lower"] = input.Body.MinisterialScores.LowerBound
			}
			if input.Body.MinisterialScores.UpperBound > 0 {
				whereClauses = append(whereClauses, "b.ministerial_score <= :score_upper")
				args["score_upper"] = input.Body.MinisterialScores.UpperBound
			}
		}
		if input.Body.ResearchAreas != nil {
			whereClauses = append(whereClauses, "s.research_area = ANY(:research_areas)")
			args["research_areas"] = input.Body.ResearchAreas.ResearchAreas
		}
	}

	// Combine query
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Order by last name, first name
	query += " ORDER BY s.last_name, s.first_name"

	// Execute query
	connection := database.GetDB()
	rows, err := connection.NamedQuery(query, args)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search query: %w", err)
	}
	defer rows.Close()

	// Parse results
	var scientists []responses.ScientistBody
	for rows.Next() {
		var scientist responses.ScientistBody
		if err := rows.StructScan(&scientist); err != nil {
			return nil, fmt.Errorf("failed to scan result row: %w", err)
		}
		scientists = append(scientists, scientist)
	}

	return scientists, nil
}
