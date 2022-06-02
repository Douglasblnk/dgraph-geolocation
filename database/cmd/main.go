package main

import "dgraph-osm/database/internal/services"

func main() {
	services.InsertCountries()
	services.InsertStates()
	services.InsertCities()
}
