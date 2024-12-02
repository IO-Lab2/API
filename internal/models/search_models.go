package models

type SearchInput struct {
	Query string           `query:"q" doc:"The search query to run."`
	Body  *SearchInputBody `json:"filters" doc:"The body of the search query."`
}

type SearchInputBody struct {
	AcademicTitles     *TitlesRequest             `json:"academic_titles" doc:"The academic titles to filter the search result."`
	CitationsCounts    *CitationsCountsRequest    `json:"citations_counts" doc:"The number of citations to filter the search result."`
	MinisterialScores  *MinisterialScoresRequest  `json:"ministerial_scores" doc:"The ministerial scores to filter the search result."`
	Organizations      *OrganizationsRequest      `json:"organizations" doc:"The organizations to filter the search result."`
	PublicationsCounts *PublicationsCountsRequest `json:"publications_counts" doc:"The number of publications to filter the search result."`
	ResearchAreas      *ResearchAreasRequest      `json:"research_areas" doc:"The research areas to filter the search result."`
}

type TitlesRequest struct {
	Titles []AcademicTitle `json:"titles" doc:"The academic titles to filer the search result."`
}

type CitationsCountsRequest struct {
	LowerBound int `json:"lower_bound" doc:"The minimum number of citations to filter the search result."`
	UpperBound int `json:"upper_bound" doc:"The maximum number of citations to filter the search result."`
}

type MinisterialScoresRequest struct {
	LowerBound int `json:"lower_bound" doc:"The minimum ministerial score to filter the search result."`
	UpperBound int `json:"upper_bound" doc:"The maximum ministerial score to filter the search result."`
}

type OrganizationsRequest struct {
	Organizations []OrganizationsNames `json:"organizations" doc:"The organizations to filter the search result."`
}

type PublicationsCountsRequest struct {
	LowerBound int `json:"lower_bound" doc:"The minimum number of publications to filter the search result."`
	UpperBound int `json:"upper_bound" doc:"The maximum number of publications to filter the search result."`
}

type ResearchAreasRequest struct {
	ResearchAreas []ResearchAreas `json:"research_areas" doc:"The research areas to filter the search result."`
}

type CitationsCounts struct {
	Min *int `json:"min" doc:"The minimum number of citations."`
	Max *int `json:"max" doc:"The maximum number of citations."`
}

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
