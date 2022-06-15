package models

type Country struct {
	Name           string   `json:"Country.name"`
	Capital        string   `json:"Country.capital"`
	Iso2           string   `json:"Country.iso2"`
	Iso3           string   `json:"Country.iso3"`
	Currency       string   `json:"Country.currency"`
	CurrencySymbol string   `json:"Country.currency_symbol"`
	Emoji          string   `json:"Country.emoji"`
	States         []*State `json:"Country.states"`
}
