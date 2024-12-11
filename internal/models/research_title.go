package models

type ResearchArea struct {
	Area string `db:"research_area" json:"research_area" doc:"Research area" format:"string" example:"Artificial Intelligence"`
}
