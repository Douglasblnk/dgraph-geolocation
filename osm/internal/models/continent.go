package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Continent struct {
	UID        string     `json:"uid,omitempty"`
	ID         int        `json:"Continent.id" dgraph:"index=int unique upsert"`
	Name       string     `json:"Continent.name,omitempty" dgraph:"index=term lang"`
	NameEn     string     `json:"Continent.name@en,omitempty" bun:"-"`
	NameEs     string     `json:"Continent.name@es,omitempty" bun:"-"`
	NamePt     string     `json:"Continent.name@pt,omitempty" bun:"-"`
	NameRu     string     `json:"Continent.name@ru,omitempty" bun:"-"`
	NameZh     string     `json:"Continent.name@zh,omitempty" bun:"-"`
	NameAbbr   string     `json:"Continent.name@abbr,omitempty" bun:"-"`
	AsciiName  string     `json:"Continent.asciiName,omitempty" bun:"-"`
	Iso        string     `json:"Continent.iso,omitempty" bun:"code" dgraph:"index=exact unique upsert"`
	Latitude   float32    `json:"Continent.latitude,omitempty" bun:"-"`
	Longitude  float32    `json:"Continent.longitude,omitempty" bun:"-"`
	Population int        `json:"Continent.population,omitempty" bun:"-"`
	Timezone   string     `json:"Continent.timezone,omitempty" bun:"-"`
	UpdatedAt  time.Time  `json:"Continent.updatedAt,omitempty" bun:"-"`
	Countries  []*Country `json:"Country,omitempty" dgraph:"count reverse" bun:"-"`
	DType      []string   `json:"dgraph.type" bun:"-"`

	bun.BaseModel `json:",omitempty" bun:"table:continentcodes"`
}
