package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Country struct {
	UID                string    `json:"uid,omitempty" bun:"-"`
	GeonameId          int       `json:"Country.geonameId" bun:"geonameid" dgraph:"index=int unique upsert"`
	ContinentId        int       `json:"Country.continentId" bun:"-"`
	Name               string    `json:"Country.name" dgraph:"index=term lang"`
	NameEn             string    `json:"Country.name@en,omitempty" bun:"-"`
	NameEs             string    `json:"Country.name@es,omitempty" bun:"-"`
	NamePt             string    `json:"Country.name@pt,omitempty" bun:"-"`
	NameRu             string    `json:"Country.name@ru,omitempty" bun:"-"`
	NameZh             string    `json:"Country.name@zh,omitempty" bun:"-"`
	NameAbbr           string    `json:"Country.name@abbr,omitempty" bun:"-"`
	AsciiName          string    `json:"Country.asciiName,omitempty" bun:"-"`
	IsoAlpha2          string    `json:"Country.isoAlpha2,omitempty" bun:"iso_alpha2"`
	IsoAlpha3          string    `json:"Country.isoAlpha3,omitempty" bun:"iso_alpha3"`
	IsoNumeric         int       `json:"Country.isoNumeric,omitempty" bun:"iso_numeric"`
	FipsCode           string    `json:"Country.fipsCode,omitempty" bun:"fips_code"`
	Capital            string    `json:"Country.capital,omitempty"`
	AreaInKm           float32   `json:"Country.areaInKm,omitempty" bun:"areainsqkm"`
	Latitude           float32   `json:"Country.latitude,omitempty" bun:"-"`
	Longitude          float32   `json:"Country.longitude,omitempty" bun:"-"`
	Continent          string    `json:"Country.continent,omitempty"`
	Tld                string    `json:"Country.tld,omitempty"`
	Currency           string    `json:"Country.currency,omitempty"`
	CurrencyName       string    `json:"Country.currencyName,omitempty" bun:"currencyname"`
	Phone              string    `json:"Country.phone,omitempty"`
	PostalCodeFormat   string    `json:"Country.postalCodeFormat,omitempty" bun:"postalcodeformat"`
	PostalCodeRegex    string    `json:"Country.postalCodeRegex,omitempty" bun:"postalcoderegex"`
	Languages          string    `json:"Country.languages,omitempty"`
	Neighbours         string    `json:"Country.neighbours,omitempty"`
	EquivalentFipsCode string    `json:"Country.equivalentFipsCode,omitempty" bun:"equivalentfipscode"`
	Population         int       `json:"Country.population,omitempty"`
	Timezone           string    `json:"Country.timezone,omitempty" bun:"-"`
	States             []*State  `json:"State,omitempty" dgraph:"count reverse" bun:"-"`
	UpdatedAt          time.Time `json:"Country.updatedAt" bun:"-"`
	DType              []string  `json:"dgraph.type" bun:"-"`

	bun.BaseModel `json:",omitempty" bun:"table:countryinfo"`
}
