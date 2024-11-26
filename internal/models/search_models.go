package models

type SearchInput struct {
	Query string          `query:"q" doc:"The search query to run."`
	Body  SearchInputBody `json:"filters" doc:"The body of the search query."`
}

type SearchInputBody struct {
	Titles             []AcademicTitle      `json:"titles" doc:"The academic titles to filer the search result."`
	CitationsCounts    *CitationsCounts     `json:"citations_counts" doc:"The citation counts to filter the search result."`
	MinisterialScores  *MinisterialScores   `json:"ministerial_scores" doc:"The ministerial scores to filter the search result."`
	Organizations      []OrganizationsNames `json:"organizations" doc:"The organizations to filter the search result."`
	PublicationsCounts *PublicationsCounts  `json:"publications_counts" doc:"The publications counts to filter the search result."`
	ResearchAreas      []ResearchAreas      `json:"research_areas" doc:"The research areas to filter the search result."`
}

type CitationsCounts struct {
	Min int `json:"min" doc:"The minimum number of citations."`
	Max int `json:"max" doc:"The maximum number of citations."`
}

type MinisterialScores struct {
	Min int `json:"min" doc:"The minimum ministerial score."`
	Max int `json:"max" doc:"The maximum ministerial score."`
}

type OrganizationsNames struct {
	Name string `json:"name" doc:"The name of the organization."`
}

type PublicationsCounts struct {
	Min int `json:"min" doc:"The minimum number of publications."`
	Max int `json:"max" doc:"The maximum number of publications."`
}

type ResearchAreas struct {
	Name string `json:"name" doc:"The name of the research area."`
}
