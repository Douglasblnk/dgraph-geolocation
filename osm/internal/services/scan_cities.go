package services

import (
	"dgraph-osm/osm/internal/models"
	"dgraph-osm/osm/internal/utils"
	"encoding/json"
	"os"

	"github.com/paulmach/osm"
)

type City struct {
}

var cityTypes []string = []string{
	"city",
	"district",
	"county",
	"municipality",
	"town",
	"village",
}

func ScanCities(node *osm.Node, cities *os.File) *models.City {
	var citiesModel *models.City

	for _, cityType := range cityTypes {
		if node.Tags.Find("place") == cityType {
			citiesModel = &models.City{
				ID:              node.ID,
				Name:            node.Tags.Find("name"),
				Latitude:        node.Lat,
				Longitude:       node.Lon,
				Is_capital:      node.Tags.Find("is_capital"),
				Is_in:           utils.StringPointer(node.Tags.Find("is_in")),
				Is_in_continent: utils.StringPointer(node.Tags.Find("is_in:continent")),
				Is_in_country: utils.StringPointer(
					utils.Ternary(
						node.Tags.Find("is_in:country"),
						node.Tags.Find("is_in"),
					),
				),
				Is_in_state: utils.StringPointer(
					utils.Ternary(
						node.Tags.Find("is_in:state"),
						node.Tags.Find("is_in"),
					),
				),
				Population: utils.StringPointer(node.Tags.Find("population")),
			}

			// if *citiesModel.Is_in_state == "" || *citiesModel.Is_in_country == "" {
			// 	url := fmt.Sprintf(
			// 		"https://nominatim.openstreetmap.org/reverse?lat=%v&lon=%v&format=json",
			// 		citiesModel.Latitude,
			// 		citiesModel.Longitude,
			// 	)

			// 	resp, err := http.Get(url)

			// 	if err != nil {
			// 		log.Fatal("err =>", err)
			// 	}

			// 	defer resp.Body.Close()

			// 	fmt.Println(json.NewDecoder(resp.Body).Decode())
			// }

			data, _ := json.Marshal(citiesModel)

			cities.WriteString(string(data))
			cities.WriteString("\n")
		}
	}

	return citiesModel
}
