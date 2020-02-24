package model

// Article ...
type Article struct {
	Type              string  `json:"type,omitempty"`
	HarvesterID       string  `json:"harvesterId,omitempty"`
	CommercialPartner string  `json:"commercialPartner,omitempty"`
	LogoURL           string  `json:"logoURL,omitempty"`
	CerebroScore      float64 `json:"cerebro-score,omitempty"`
	URL               string  `json:"url,omitempty"`
	Title             string  `json:"title,omitempty"`
	CleanImage        string  `json:"cleanImage,omitempty"`
}

// Response ...
type Response struct {
	HTTPStatus int `json:"httpStatus"`
	Response   struct {
		Items []Article `json:"items"`
	}
}