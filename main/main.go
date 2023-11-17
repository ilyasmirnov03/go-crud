package main

import (
	"l33ass/crud/database"
	"l33ass/crud/routes"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	// Prepare ENV and other variables
	sql_host := os.Getenv("POSTGRES_HOST")
	sql_user := os.Getenv("POSTGRES_USER")
	sql_password := os.Getenv("POSTGRES_PASSWORD")
	app_port := os.Getenv("APP_PORT")

	// Concatenate colon character to port to start the server
	var port_string strings.Builder
	port_string.WriteString(":")
	port_string.WriteString(app_port)

	// Call setup functions
	setup_bundler()
	routes.SetupRouter()
	database.DBConnect(sql_user, sql_password, sql_host)
	log.Fatal(http.ListenAndServe(port_string.String(), nil))
}
