package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var regions []string = []string{
	"south-america-latest.osm.pbf",
}

func download(url string, region string, wg *sync.WaitGroup) {
	out, err := os.Create("osm/data/pbf/" + region)

	if err != nil {
		log.Fatal("err =>", err)
	}

	defer out.Close()

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("err =>", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("resp.Status =>", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal("err =>", err)
	}

	wg.Done()
}

func downloadRegion(region string, wg *sync.WaitGroup) {
	url := "http://download.geofabrik.de/" + region
	download(url, region, wg)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(len(regions))

	for _, region := range regions {
		go downloadRegion(region, wg)
	}

	fmt.Println("Downloading regions...")

	wg.Wait()
}
