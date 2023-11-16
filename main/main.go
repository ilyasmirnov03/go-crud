package main

import (
	"l33ass/crud/database"
	"l33ass/crud/routes"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Prepare ENV and other variables
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Env variables failed to load", err)
	}

	// Concatenate colon character to port to start the server
	var port_string strings.Builder
	port_string.WriteString(":")
	port_string.WriteString(env["APP_PORT"])

	// Call setup functions
	setup_bundler()
	routes.SetupRouter()
	database.DBConnect()
	log.Fatal(http.ListenAndServe(port_string.String(), nil))
}
