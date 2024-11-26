package models

type SearchInput struct {
	Query string          `query:"q" doc:"The search query to run."`
	Body  SearchInputBody `json:"filters" doc:"The body of the search query."`
}

type SearchInputBody struct {
}
