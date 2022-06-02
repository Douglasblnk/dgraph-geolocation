package models

type Country struct {
	Name           string `json:"Country.Name"`
	Capital        string `json:"Country.Capital"`
	Iso2           string `json:"Country.Iso2"`
	Iso3           string `json:"Country.Iso3"`
	Currency       string `json:"Country.Currency"`
	CurrencySymbol string `json:"Country.Currency_symbol"`
	Emoji          string `json:"Country.Emoji"`
}
