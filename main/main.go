package main

import (
	"l33ass/crud/database"
	"l33ass/crud/routes"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	setup_bundler()
	routes.SetupRouter()
	database.DBConnect()
	log.Fatal(http.ListenAndServe(":1337", nil))
}
