package services

import (
	"context"
	"dgraph-osm/database/internal/connection"
	"dgraph-osm/database/internal/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

type Country struct {
	Name           string `json:"name"`
	Capital        string `json:"capital"`
	Iso2           string `json:"iso2"`
	Iso3           string `json:"iso3"`
	Currency       string `json:"currency"`
	CurrencySymbol string `json:"currency_symbol"`
	Emoji          string `json:"emoji"`
}

func getCountries() []*models.Country {
	countryFile, _ := os.Open("/home/douglasblnk/Documentos/Projects/dgraph-go/database/data/countries.json")
	defer countryFile.Close()

	countriesByteValue, _ := ioutil.ReadAll(countryFile)

	countryModel := []*Country{}
	json.Unmarshal(countriesByteValue, &countryModel)

	var countries []*models.Country

	for _, country := range countryModel {
		c := &models.Country{
			Name:           country.Name,
			Capital:        country.Capital,
			Iso2:           country.Iso2,
			Iso3:           country.Iso3,
			Currency:       country.Currency,
			CurrencySymbol: country.CurrencySymbol,
			Emoji:          country.Emoji,
		}

		countries = append(countries, c)
	}

	return countries
}

func InsertCountries() {
	dgcon, close := connection.DgraphConnection()
	defer close()

	countries := getCountries()

	json, _ := json.Marshal(&countries)

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   json,
	}

	response, err := dgcon.NewTxn().Mutate(context.Background(), mu)

	if err != nil {
		log.Fatal("err", err)
	}

	println(response)
}
