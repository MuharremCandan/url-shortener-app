package entities

type CreateRedirectDTO struct {
	URL string `json:"url"`
}

type CreatedRedirectResponseDTO struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

type GetRedirectRequestDTO struct {
	Code string `json:"code"`
}

type GetRedirectResponseDTO struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}
