package responses

type ResearchAreasResponse struct {
	Body []ResearchArea `json:"body" doc:"Research areas"`
}

type ResearchAreaExtendedResponse struct {
	Body []ResearchAreaExtended `json:"body" doc:"Research areas"`
}
