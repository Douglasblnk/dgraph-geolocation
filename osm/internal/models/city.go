package models

import (
	"time"

	"github.com/uptrace/bun"
)

type City struct {
	UID        string    `json:"uid,omitempty" bun:"-"`
	GeonameId  int       `json:"City.geonameId" bun:"geonameid" dgraph:"index=int unique upsert"`
	StateId    int       `json:"City.stateId" bun:"-"`
	Name       string    `json:"City.name" dgraph:"index=term lang"`
	NameEn     string    `json:"City.name@en,omitempty" bun:"-"`
	NameEs     string    `json:"City.name@es,omitempty" bun:"-"`
	NamePt     string    `json:"City.name@pt,omitempty" bun:"-"`
	NameRu     string    `json:"City.name@ru,omitempty" bun:"-"`
	NameZh     string    `json:"City.name@zh,omitempty" bun:"-"`
	NameAbbr   string    `json:"City.name@abbr,omitempty" bun:"-"`
	AsciiName  string    `json:"City.asciiName" bun:"asciiname"`
	Latitude   float32   `json:"City.latitude"`
	Longitude  float32   `json:"City.longitude"`
	Population int       `json:"City.population"`
	Timezone   string    `json:"City.timezone" bun:"timezone"`
	UpdatedAt  time.Time `json:"City.updatedAt" bun:"moddate"`
	DType      []string  `json:"dgraph.type" bun:"-"`

	bun.BaseModel `json:",omitempty" bun:"table:geoname"`
}
