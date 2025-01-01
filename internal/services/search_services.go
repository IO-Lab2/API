package services

import (
	"errors"
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
		s.position,
		s.email, 
		s.profile_url, 
		s.created_at, 
		s.updated_at, 
		ARRAY_AGG(DISTINCT ra.name) AS research_areas,
		b.h_index_wos, 
		b.h_index_scopus, 
		b.publication_count, 
		b.ministerial_score
	FROM 
		scientists s
	LEFT JOIN 
		scientists_research_areas sra ON s.id = sra.scientist_id
	LEFT JOIN 
		research_areas ra ON sra.research_area_id = ra.id
	LEFT JOIN 
		bibliometrics b ON s.id = b.scientist_id
	LEFT JOIN 
		scientists_publications sp ON s.id = sp.scientist_id
	LEFT JOIN 
		publications p ON sp.publication_id = p.id
	LEFT JOIN 
		scientist_organization so ON s.id = so.scientist_id
	LEFT JOIN 
		organizations o ON so.organization_id = o.id
	`

	// Where clauses
	whereClauses := []string{}
	args := map[string]interface{}{}

	// Existing filters...
	if isNotEmpty(input.Name) {
		whereClauses = append(whereClauses, "s.first_name ILIKE :name")
		args["name"] = "%" + input.Name + "%"
	}
	if isNotEmpty(input.Surname) {
		whereClauses = append(whereClauses, "s.last_name ILIKE :surname")
		args["surname"] = "%" + input.Surname + "%"
	}
	if isNotEmpty(input.AcademicTitles) {
		whereClauses = append(whereClauses, "s.academic_title = ANY(:academic_titles)")
		args["academic_titles"] = pq.Array(input.AcademicTitles)
	}
	if isNotEmpty(input.Organizations) {
		whereClauses = append(whereClauses, "o.name = ANY(:organizations)")
		args["organizations"] = pq.Array(input.Organizations)
	}
	if isNotEmpty(input.ResearchAreas) {
		whereClauses = append(whereClauses, "ra.name = ANY(:research_areas)")
		args["research_areas"] = pq.Array(input.ResearchAreas)
	}
	if isNotEmpty(input.MinPublications) {
		whereClauses = append(whereClauses, "b.publication_count >= :min_publications")
		args["min_publications"] = input.MinPublications
	}
	if isNotEmpty(input.MaxPublications) {
		whereClauses = append(whereClauses, "b.publication_count <= :max_publications")
		args["max_publications"] = input.MaxPublications
	}
	if isNotEmpty(input.MinMinisterialScore) {
		whereClauses = append(whereClauses, "b.ministerial_score >= :min_score")
		args["min_score"] = input.MinMinisterialScore
	}
	if isNotEmpty(input.MaxMinisterialScore) {
		whereClauses = append(whereClauses, "b.ministerial_score <= :max_score")
		args["max_score"] = input.MaxMinisterialScore
	}
	if isNotEmpty(input.Positions) {
		whereClauses = append(whereClauses, "s.position = ANY(:positions)")
		args["positions"] = pq.Array(input.Positions)
	}
	if isNotEmpty(input.JournalTypes) {
		whereClauses = append(whereClauses, "p.journal_type = ANY(:journal_types)")
		args["journal_types"] = pq.Array(input.JournalTypes)
	}
	if isNotEmpty(input.Publishers) {
		whereClauses = append(whereClauses, "p.publisher = ANY(:publishers)")
		args["publishers"] = pq.Array(input.Publishers)
	}

	// New Year-Specific Ministerial Score Filter
	if len(input.YearScoreFilters) > 0 {
		yearConditions := []string{}
		for i, filter := range input.YearScoreFilters {
			condition := fmt.Sprintf(`
				(EXTRACT(YEAR FROM p.publication_date) = :year_%d 
				AND SUM(p.ministerial_score) BETWEEN :min_score_%d AND :max_score_%d)`, i, i, i)
			yearConditions = append(yearConditions, condition)
			args[fmt.Sprintf("year_%d", i)] = filter.Year
			args[fmt.Sprintf("min_score_%d", i)] = filter.MinScore
			args[fmt.Sprintf("max_score_%d", i)] = filter.MaxScore
		}
		whereClauses = append(whereClauses, "("+strings.Join(yearConditions, " OR ")+")")
	}

	// Combine query
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Group and order results
	query += `
	GROUP BY s.id, b.h_index_wos, b.h_index_scopus, b.publication_count, b.ministerial_score
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
	for rows.Next() {
		var scientist responses.ScientistBody
		var researchAreaNames []string

		if err := rows.Scan(
			&scientist.ID,
			&scientist.FirstName,
			&scientist.LastName,
			&scientist.AcademicTitle,
			&scientist.Position,
			&scientist.Email,
			&scientist.ProfileUrl,
			&scientist.CreatedAt,
			&scientist.UpdatedAt,
			pq.Array(&researchAreaNames),
			&scientist.Bibliometrics.HIndexWOS,
			&scientist.Bibliometrics.HIndexScopus,
			&scientist.Bibliometrics.PublicationCount,
			&scientist.Bibliometrics.MinisterialScore,
		); err != nil {
			logging.Logger.Error("ERROR: Error scanning row: ", err)
			return nil, fmt.Errorf("failed to scan result row: %w", err)
		}

		// Add research areas to the scientist
		for _, name := range researchAreaNames {
			scientist.ResearchAreas = append(scientist.ResearchAreas, responses.ResearchArea{Name: name})
		}

		scientists = append(scientists, scientist)
	}

	logging.Logger.Debug("Number of scientists found: ", len(scientists))
	if len(scientists) == 0 {
		return nil, errors.New("No scientists found")
	}

	return scientists, nil
}

func isNotEmpty(param interface{}) bool {
	switch v := param.(type) {
	case string:
		return v != ""
	case []string:
		return len(v) > 0
	case int:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}
