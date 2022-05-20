package main

import (
	"dgraph-osm/osm/internal/services"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/douglasblnk/Documentos/Multiplier/dgraph-geolocation/osm/data/pbf/south-america-latest.osm.pbf")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	services.Scanner(file)
}
