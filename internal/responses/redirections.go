package responses

type RedirectResponse struct {
	Url string `header:"Location" doc:"The URL to redirect to." format:"url" example:"http://example.com"`
}

type NoInput struct {
}

type Response struct {
	Status int `json:"status"`
	Body   struct {
		Message string `json:"message" doc:"Response message" format:"string" example:"Success"`
	}
}

type Healthcheck struct {
	Message string `json:"message" doc:"Message" format:"string" example:"Healthy"`
}
