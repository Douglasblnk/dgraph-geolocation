download:
	go run osm/cmd/osm-download/main.go

scanner:
	go run osm/cmd/osm-scanner/main.go

clear:
	go run osm/cmd/clear-dgraph/main.go

insert:
	go run database/cmd/main.go

schemas:
	cd database/schemas && curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'