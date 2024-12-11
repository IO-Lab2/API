package models

type SearchInput struct {
	Name    string           `query:"name" doc:"The name of the scientist to search for."`
	Surname string           `query:"surname" doc:"The surname of the scientist to search for."`
	Body    *SearchInputBody `json:"filters" doc:"The body of the search query."`
}

type SearchInputBody struct {
	AcademicTitles     *TitlesRequest             `json:"academic_titles,omitempty" doc:"The academic titles to filter the search result."`
	MinisterialScores  *MinisterialScoresRequest  `json:"ministerial_scores,omitempty" doc:"The ministerial scores to filter the search result."`
	Organizations      *OrganizationsRequest      `json:"organizations,omitempty" doc:"The organizations to filter the search result."`
	PublicationsCounts *PublicationsCountsRequest `json:"publications_counts,omitempty" doc:"The number of publications to filter the search result."`
	ResearchAreas      *ResearchAreasRequest      `json:"research_areas,omitempty" doc:"The research areas to filter the search result."`
}

type TitlesRequest struct {
	Titles []AcademicTitle `json:"titles,omitempty" doc:"The academic titles to filer the search result."`
}

type MinisterialScoresRequest struct {
	LowerBound int `json:"lower_bound,omitempty" doc:"The minimum ministerial score to filter the search result." format:"int64" minimum:"0"`
	UpperBound int `json:"upper_bound,omitempty" doc:"The maximum ministerial score to filter the search result." format:"int64"`
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
	Min *int `json:"min" doc:"The minimum ministerial score."`
	Max *int `json:"max" doc:"The maximum ministerial score."`
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
