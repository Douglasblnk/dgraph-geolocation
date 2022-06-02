package models

type State struct {
	Name        string  `json:"State.Name"`
	Latitude    float64 `json:"State.Latitude"`
	Longitude   float64 `json:"State.Longitude"`
	Country     string  `json:"State.Country_name"`
	CountryCode string  `json:"State.Country_code"`
	StateCode   string  `json:"State.State_code"`
}
