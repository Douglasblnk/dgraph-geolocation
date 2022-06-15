package services

import (
	"dgraph-osm/osm/internal/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

const baseUrl string = "https://nominatim.openstreetmap.org/"

func ScanCities(node *osm.Node, cities *os.File) *models.City {
	var citiesModel *models.City

	for _, cityType := range cityTypes {
		if node.Tags.Find("place") == cityType {
			url := fmt.Sprintf(baseUrl+"reverse?lat=%f&lon=%f&format=json", node.Lat, node.Lon)

			response, err := http.Get(url)

			if err != nil {
				log.Fatal("err =>", err)
			}

			body, _ := ioutil.ReadAll(response.Body)

			cities.WriteString(string(body))
			cities.WriteString("\n")

			// citiesModel = &models.City{
			// 	ID:              node.ID,
			// 	Name:            node.Tags.Find("name"),
			// 	Latitude:        node.Lat,
			// 	Longitude:       node.Lon,
			// 	Is_capital:      node.Tags.Find("is_capital"),
			// 	Is_in:           utils.StringPointer(node.Tags.Find("is_in")),
			// 	Is_in_continent: utils.StringPointer(node.Tags.Find("is_in:continent")),
			// 	Is_in_country: utils.StringPointer(
			// 		utils.Ternary(
			// 			node.Tags.Find("is_in:country"),
			// 			node.Tags.Find("is_in"),
			// 		),
			// 	),
			// 	Is_in_state: utils.StringPointer(
			// 		utils.Ternary(
			// 			node.Tags.Find("is_in:state"),
			// 			node.Tags.Find("is_in"),
			// 		),
			// 	),
			// 	Population: utils.StringPointer(node.Tags.Find("population")),
			// }

			// data, _ := json.Marshal(citiesModel)

			// cities.WriteString(string(data))
			// cities.WriteString("\n")
		}
	}

	return citiesModel
}
