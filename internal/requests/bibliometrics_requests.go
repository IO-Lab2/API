package requests

import "github.com/google/uuid"

type BibliometricsID struct {
	ID uuid.UUID `path:"id" doc:"Bibliometric ID" format:"UUID" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}

type BibliometricsScientistID struct {
	ID uuid.UUID `path:"id" doc:"Scientist ID" format:"UUID" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}

type BibliometricsAuthor struct {
	Author string `path:"author" doc:"Author name" example:"John Doe"`
}
type CreateBibliometric struct {
	HIndexWos        int       `query:"h_index_wos" json:"h_index_wos" doc:"HIndex Wos" format:"int" example:"1"`
	HIndexScopus     int       `query:"h_index_scopus" json:"h_index_scopus" doc:"HIndex Scopus" format:"int" example:"2"`
	PublicationCount int       `query:"publication_count" json:"publication_count" doc:"Publication count" format:"int" example:"123"`
	MinisterialScore float64   `query:"ministerial_score" json:"ministerial_score" doc:"Ministerial score" format:"float" example:"65.7"`
	ScientistID      uuid.UUID `query:"scientist_id" json:"scientist_id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
