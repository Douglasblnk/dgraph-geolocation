geo:
	go run github.com/cosmtrek/air -c geolocation/cmd/.air.toml

download:
	go run osm/cmd/osm-download/main.go

schemas:
	go run osm/cmd/dgraph-schema/main.go

scanner:
	go run osm/cmd/osm-scanner/main.go

clear:
	go run osm/cmd/clear-dgraph/main.go