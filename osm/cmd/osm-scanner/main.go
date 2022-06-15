package main

import (
	"dgraph-osm/osm/internal/services"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/douglasblnk/Documentos/Projects/dgraph-go/osm/data/pbf/sul-latest.osm.pbf")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	services.Scanner(file)
}
