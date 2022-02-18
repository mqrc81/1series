// Package cmd/zeries is the entry-point which starts the application
// and initializes database, external clients, list of genres/networks
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
	"github.com/mqrc81/zeries/api"
	"github.com/mqrc81/zeries/postgres"
	"github.com/mqrc81/zeries/trakt"
)

// TODO:
//  - initialize Genres & Networks
func main() {
	log.Println("Starting application...")

	// Environment variables need to be initialized from local file first when ran locally
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}

	store, err := postgres.Init(os.Getenv("DATABASE_URL"))
	checkError(err)

	tmdbClient, err := tmdb.Init(os.Getenv("TMDB_KEY"))
	checkError(err)

	traktClient, err := trakt.Init(os.Getenv("TRAKT_KEY"))
	checkError(err)

	handler, err := api.Init(*store, tmdbClient, traktClient)
	checkError(err)

	log.Println("Listening on " + os.Getenv("BACKEND_URL"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
