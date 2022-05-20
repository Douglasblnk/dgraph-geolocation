package services

import (
	"context"
	"log"
	"os"
	"runtime"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
)

func Scanner(file *os.File) {
	scanner := osmpbf.New(context.Background(), file, runtime.GOMAXPROCS(-1))
	defer scanner.Close()

	scanner.SkipWays = true
	scanner.SkipRelations = true

	for scanner.Scan() {
		switch node := scanner.Object().(type) {
		case *osm.Node:
			ScanCountries(node)
		}
	}

	scanErr := scanner.Err()

	if scanErr != nil {
		log.Fatal(scanErr)
	}

}
