package services

import (
	"dgraph-osm/osm/internal/models"
	"dgraph-osm/osm/internal/utils"
	"encoding/json"
	"os"

	"github.com/paulmach/osm"
)

func ScanCountries(node *osm.Node, country *os.File) *models.Country {
	var countryModel *models.Country

	if node.Tags.Find("place") == "country" {
		countryModel = &models.Country{
			Name:         node.Tags.Find("name"),
			Latitude:     node.Lat,
			Longitude:    node.Lon,
			Capital_city: utils.StringPointer(node.Tags.Find("capital_city")),
			Country_code: utils.StringPointer(
				utils.Ternary(
					node.Tags.Find("ISO3166-1:alpha2"),
					node.Tags.Find("country_code_fips"),
				),
			),
			Population: utils.StringPointer(node.Tags.Find("population")),
		}

		data, _ := json.Marshal(countryModel)

		country.WriteString(string(data))
		country.WriteString("\n")
	}

	return countryModel
}
