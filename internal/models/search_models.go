package models

type SearchInput struct {
	Name                string   `query:"name" doc:"The name of the scientist to search for."`
	Surname             string   `query:"surname" doc:"The surname of the scientist to search for."`
	AcademicTitles      []string `query:"academic_titles[]" doc:"List of academic titles to filter the search result."`
	Organizations       []string `query:"organizations[]" doc:"List of organizations to filter the search result."`
	ResearchAreas       []string `query:"research_areas[]" doc:"List of research areas to filter the search result."`
	MinMinisterialScore float64  `query:"ministerial_score_min" doc:"Minimum ministerial score."`
	MaxMinisterialScore float64  `query:"ministerial_score_max" doc:"Maximum ministerial score."`
	MinPublications     int      `query:"publications_min" doc:"Minimum number of publications."`
	MaxPublications     int      `query:"publications_max" doc:"Maximum number of publications."`
	Positions           []string `query:"positions[]" doc:"List of positions to filter the search result."`
	JournalTypes        []string `query:"journal_types[]" doc:"List of journal types to filter the search result."`
	PublicationsYears   []int    `query:"publications_years[]" doc:"List of publication years to filter the search result."`
}

type TitlesRequest struct {
	Titles []AcademicTitle `json:"titles,omitempty" doc:"The academic titles to filer the search result."`
}

type MinisterialScoresRequest struct {
	LowerBound float64 `json:"lower_bound,omitempty" doc:"The minimum ministerial score to filter the search result." format:"float64"`
	UpperBound float64 `json:"upper_bound,omitempty" doc:"The maximum ministerial score to filter the search result." format:"float64"`
}

type OrganizationsRequest struct {
	Organizations []OrganizationsNames `json:"organizations,omitempty" doc:"The organizations to filter the search result."`
}

type PublicationsCountsRequest struct {
	LowerBound int `json:"lower_bound,omitempty" doc:"The minimum number of publications to filter the search result."`
	UpperBound int `json:"upper_bound,omitempty" doc:"The maximum number of publications to filter the search result."`
}

type ResearchAreasRequest struct {
	ResearchAreas []ResearchAreas `json:"research_areas,omitempty" doc:"The research areas to filter the search result."`
}

// Additional models
type MinisterialScores struct {
	Min *float64 `json:"min" doc:"The minimum ministerial score."`
	Max *float64 `json:"max" doc:"The maximum ministerial score."`
}

type OrganizationsNames struct {
	Name *string `json:"name" doc:"The name of the organization."`
}

type PublicationsCounts struct {
	Min *int `json:"min" doc:"The minimum number of publications."`
	Max *int `json:"max" doc:"The maximum number of publications."`
}

type ResearchAreas struct {
	Name *string `json:"name" doc:"The name of the research area."`
}
