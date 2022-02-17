// Package cmd/zeries
// Entry-point which starts the application and initializes database, external clients, list of genres/networks
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
	"github.com/mqrc81/zeries/api"
	"github.com/mqrc81/zeries/trakt"
)

// TODO:
//  - initialize DB
//  - initialize Genres & Networks
func main() {
	log.Println("Starting application...")

	// Environment variables need to be initialized from local file first when ran locally
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}

	tmdbClient, err := tmdb.Init(os.Getenv("TMDB_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	traktClient, err := trakt.Init(os.Getenv("TRAKT_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	handler := api.Init(tmdbClient, traktClient)

	log.Println("Listening on " + os.Getenv("BACKEND_URL"))
	if err = http.ListenAndServe(":"+os.Getenv("PORT"), handler); err != nil {
		log.Fatalln(err)
	}
}
