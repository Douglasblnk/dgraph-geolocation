package models

import (
	"time"

	"github.com/uptrace/bun"
)

type State struct {
	UID        string    `json:"uid,omitempty" bun:"-"`
	GeonameId  int       `json:"State.geonameId" bun:"geonameid" dgraph:"index=int unique upsert"`
	CountryId  int       `json:"State.countryId" bun:"-"`
	Name       string    `json:"State.name" dgraph:"index=term lang"`
	NameEn     string    `json:"State.name@en,omitempty" bun:"-"`
	NameEs     string    `json:"State.name@es,omitempty" bun:"-"`
	NamePt     string    `json:"State.name@pt,omitempty" bun:"-"`
	NameRu     string    `json:"State.name@ru,omitempty" bun:"-"`
	NameZh     string    `json:"State.name@zh,omitempty" bun:"-"`
	NameAbbr   string    `json:"State.name@abbr,omitempty" bun:"-"`
	AsciiName  string    `json:"State.asciiName" bun:"asciiname"`
	Latitude   float32   `json:"State.latitude"`
	Longitude  float32   `json:"State.longitude"`
	Population int       `json:"State.population"`
	Timezone   string    `json:"State.timezone" bun:"timezone"`
	UpdatedAt  time.Time `json:"State.updatedAt" bun:"moddate"`
	Cities     []*City   `json:"City,omitempty" dgraph:"count reverse" bun:"-"`
	DType      []string  `json:"dgraph.type" bun:"-"`

	bun.BaseModel `json:",omitempty" bun:"table:geoname"`
}
