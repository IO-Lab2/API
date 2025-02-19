package services

import (
	"errors"
	"fmt"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"
	"sort"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/lib/pq"
)

func SearchForScientists(input *models.SearchInput) ([]responses.ScientistBody, int, error) {
	if input.Limit < 1 {
		input.Limit = 50
	}
	if input.Limit == 25 {
		input.Limit = 50
	}
	if input.Page < 1 {
		input.Page = 1
	}

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
        b.ministerial_score,
		b.impact_factor,
        json_agg(json_build_object(
            'year', EXTRACT(YEAR FROM p.publication_date),
            'score', p.ministerial_score
        )) AS publication_scores,
		COUNT(*) OVER() AS total_count
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
    	scientist_organization so ON s.id = so.scientist_id
	LEFT JOIN 
    	organizations o ON so.organization_id = o.id
    LEFT JOIN 
        publications p ON sp.publication_id = p.id`

	args := map[string]interface{}{
		"limit":  input.Limit,
		"offset": (input.Page - 1) * input.Limit,
	}

	var whereClauses []string

	if isNotEmpty(input.Name) {
		whereClauses = append(whereClauses, "s.first_name ILIKE :name")
		args["name"] = "%" + input.Name + "%"
	}
	if isNotEmpty(input.Surname) {
		whereClauses = append(whereClauses, "s.last_name ILIKE :surname")
		args["surname"] = "%" + input.Surname + "%"
	}

	if isNotEmpty(input.Organizations) {
		organizations := parseList(input.Organizations)
		whereClauses = append(whereClauses, "o.name = ANY(:organizations)")
		args["organizations"] = pq.Array(organizations)
	}

	if isNotEmpty(input.AcademicTitles) {
		academicTitles := parseList(input.AcademicTitles)
		whereClauses = append(whereClauses, "s.academic_title = ANY(:academic_titles)")
		args["academic_titles"] = pq.Array(academicTitles)
	}

	if isNotEmpty(input.ResearchAreas) {
		researchAreas := parseList(input.ResearchAreas)
		whereClauses = append(whereClauses, "ra.name = ANY(:research_areas)")
		args["research_areas"] = pq.Array(researchAreas)
	}

	if isNotEmpty(input.Positions) {
		positions := parseList(input.Positions)
		whereClauses = append(whereClauses, "s.position = ANY(:positions)")
		args["positions"] = pq.Array(positions)
	}

	if isNotEmpty(input.JournalTypes) {
		journalTypes := parseList(input.JournalTypes)
		whereClauses = append(whereClauses, "p.journal_type = ANY(:journal_types)")
		args["journal_types"] = pq.Array(journalTypes)
	}

	if isNotEmpty(input.Publishers) {
		publishers := parseList(input.Publishers)
		whereClauses = append(whereClauses, "p.publisher = ANY(:publishers)")
		args["publishers"] = pq.Array(publishers)
	}
	if isNotEmpty(input.Positions) {
		positions := parseList(input.Positions)
		whereClauses = append(whereClauses, "s.position = ANY(:positions)")
		args["positions"] = pq.Array(positions)
	}
	if isNotEmpty(input.Publishers) {
		publishers := parseList(input.Publishers)
		whereClauses = append(whereClauses, "p.publisher = ANY(:publishers)")
		args["publishers"] = pq.Array(publishers)
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

	if isNotEmpty(input.MinImpactFactor) {
		whereClauses = append(whereClauses, "b.impact_factor >= :min_impact_factor")
		args["min_impact_factor"] = input.MinImpactFactor
	}

	if isNotEmpty(input.MaxImpactFactor) {
		whereClauses = append(whereClauses, "b.impact_factor <= :max_impact_factor")
		args["max_impact_factor"] = input.MaxImpactFactor
	}

	// Combine query
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += `
    GROUP BY s.id, b.h_index_wos, b.h_index_scopus, b.publication_count, b.ministerial_score, b.impact_factor
	ORDER BY s.last_name, s.first_name
	LIMIT :limit OFFSET :offset`

	logging.Logger.Debug("Search query: ", query)
	logging.Logger.Debug("Search query args: ", args)
	connection := database.GetDB()
	rows, err := connection.NamedQuery(query, args)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing search query:", err)
		return nil, 0, fmt.Errorf("failed to execute search query: %w", err)
	}
	defer rows.Close()

	var totalRows int
	var scientists []responses.ScientistBody
	for rows.Next() {
		var scientist responses.ScientistBody
		var researchAreaNames []string
		var publicationScoresJSON []byte

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
			&scientist.Bibliometrics.ImpactFactor,
			&publicationScoresJSON,
			&totalRows,
		); err != nil {
			logging.Logger.Error("ERROR: Error scanning row: ", err)
			return nil, 0, fmt.Errorf("failed to scan result row: %w", err)
		}

		var publicationScores []responses.PublicationScore
		if err := json.Unmarshal(publicationScoresJSON, &publicationScores); err != nil {
			logging.Logger.Error("ERROR: Error unmarshalling publication scores: ", err)
			return nil, 0, fmt.Errorf("failed to unmarshal publication scores: %w", err)
		}

		// Combine scores for each year
		combinedScores := make(map[int]float64)
		for _, score := range publicationScores {
			if score.Year != nil && score.Score != nil {
				combinedScores[*score.Year] += *score.Score
			}
		}

		// Filter by YearScoreFilter
		if len(input.YearScoreFilter) > 0 {
			totalRows = len(scientists)
			yearScoreFilters, err := ParseYearScoreFilters(input.YearScoreFilter)
			if err != nil {
				logging.Logger.Error("ERROR: Invalid year score filter: ", err)
				return nil, 0, fmt.Errorf("invalid year score filter: %w", err)
			}

			meetsCriteria := true
			for _, filter := range yearScoreFilters {
				score, exists := combinedScores[filter.Year]
				if !exists || score < filter.MinScore || score > filter.MaxScore {
					meetsCriteria = false
					break
				}
			}

			if !meetsCriteria {
				continue // Skip this scientist
			}
		}

		// Filter by the PublicationsYears
		if isNotEmpty(input.PublicationsYears) {
			meetsCriteria := false
			for _, year := range input.PublicationsYears {
				if _, exists := combinedScores[year]; exists {
					meetsCriteria = true
					break
				}
			}

			if !meetsCriteria {
				continue // Skip this scientist
			}
		}

		// Convert combined scores map to slice
		var combinedPublicationScores []responses.PublicationScore
		for year, score := range combinedScores {
			combinedPublicationScores = append(combinedPublicationScores, responses.PublicationScore{
				Year:  &year,
				Score: &score,
			})
		}

		scientist.ResearchAreas = make([]responses.ResearchArea, len(researchAreaNames))
		for i, name := range researchAreaNames {
			scientist.ResearchAreas[i] = responses.ResearchArea{Name: &name}
		}

		// Sort the combinedPublicationScores by year in ascending order
		sort.Slice(combinedPublicationScores, func(i, j int) bool {
			return *combinedPublicationScores[i].Year < *combinedPublicationScores[j].Year
		})

		scientist.PublicationScores = combinedPublicationScores
		scientists = append(scientists, scientist)
	}

	logging.Logger.Debug("Number of scientists found: ", len(scientists))
	if len(scientists) == 0 {
		return nil, 0, errors.New("No scientists found")
	}

	return scientists, totalRows, nil
}

func isNotEmpty(param interface{}) bool {
	switch v := param.(type) {
	case string:
		return v != ""
	case []string:
		return len(v) > 0
	case []int:
		return len(v) > 0
	case []float64:
		return len(v) > 0
	case []float32:
		return len(v) > 0
	case int:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}

func parseList(input []string) []string {
	unique := make(map[string]struct{}) // A set-like structure to track unique values
	var result []string

	for _, item := range input {
		parts := strings.Split(item, ",") // Split by commas
		for _, part := range parts {
			trimmed := strings.TrimSpace(part) // Trim whitespace
			if trimmed != "" {
				if _, exists := unique[trimmed]; !exists { // Check for uniqueness
					unique[trimmed] = struct{}{}
					result = append(result, trimmed)
				}
			}
		}
	}

	return result
}

func ParseYearScoreFilters(filters []string) ([]models.YearScoreFilter, error) {
	var result []models.YearScoreFilter

	for _, filter := range filters {
		parts := strings.Split(filter, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format: %s, expected year:min_score-max_score", filter)
		}

		// Parse the year
		year, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid year in filter: %s", filter)
		}

		// Parse min_score and max_score
		scoreRange := strings.Split(parts[1], "-")
		if len(scoreRange) != 2 {
			return nil, fmt.Errorf("invalid score range in filter: %s", filter)
		}

		minScore, err := strconv.ParseFloat(scoreRange[0], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid min score in filter: %s", filter)
		}

		maxScore, err := strconv.ParseFloat(scoreRange[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid max score in filter: %s", filter)
		}

		result = append(result, models.YearScoreFilter{
			Year:     year,
			MinScore: minScore,
			MaxScore: maxScore,
		})
	}

	return result, nil
}
