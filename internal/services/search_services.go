package services

import (
	"fmt"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"
	"strings"

	"github.com/lib/pq"
)

func SearchForScientists(input *models.SearchInput) ([]responses.ScientistBody, error) {
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
		ARRAY_AGG(DISTINCT ra.name) AS research_areas
	FROM 
		scientists s
	LEFT JOIN 
		scientists_research_areas sra ON s.id = sra.scientist_id
	LEFT JOIN 
		research_areas ra ON sra.research_area_id = ra.id
	LEFT JOIN 
		bibliometrics b ON s.id = b.scientist_id
	LEFT JOIN 
		scientist_organization so ON s.id = so.scientist_id
	LEFT JOIN 
		organizations o ON so.organization_id = o.id
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
		if input.Body.ResearchAreas != nil && len(input.Body.ResearchAreas.ResearchAreas) > 0 {
			whereClauses = append(whereClauses, "ra.name = ANY(:research_areas)")
			args["research_areas"] = input.Body.ResearchAreas.ResearchAreas
		}
		if input.Body.Organizations != nil && len(input.Body.Organizations.Organizations) > 0 {
			whereClauses = append(whereClauses, "o.name = ANY(:organizations)")
			args["organizations"] = input.Body.Organizations.Organizations
		}
	}

	// Combine query
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Group and order results
	query += `
	GROUP BY s.id
	ORDER BY s.last_name, s.first_name
	`

	// Execute query
	connection := database.GetDB()
	rows, err := connection.NamedQuery(query, args)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing search query:", err)
		return nil, fmt.Errorf("failed to execute search query: %w", err)
	}
	defer rows.Close()

	// Parse results
	var scientists []responses.ScientistBody
	var count int
	for rows.Next() {
		var scientist responses.ScientistBody
		var researchAreaNames []string

		// Scan the scientist data and research areas
		if err := rows.Scan(
			&scientist.ID,
			&scientist.FirstName,
			&scientist.LastName,
			&scientist.AcademicTitle,
			&scientist.Email,
			&scientist.ProfileUrl,
			&scientist.CreatedAt,
			&scientist.UpdatedAt,
			pq.Array(&researchAreaNames),
		); err != nil {
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, fmt.Errorf("failed to scan result row: %w", err)
		}

		// Map research area names to ResearchArea structs
		for _, name := range researchAreaNames {
			scientist.ResearchAreas = append(scientist.ResearchAreas, responses.ResearchArea{Name: name})
		}

		scientists = append(scientists, scientist)
		count++
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error after iterating rows:", err)
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	logging.Logger.Info("INFO: Successfully executed search query, found: ", count)
	return scientists, nil
}
