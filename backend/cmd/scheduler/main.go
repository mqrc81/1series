// Package cmd/jobs is the entry-point for a daily jobs
// which notifies users of new releases & recommendations
// and updates list of upcoming releases stored in DB
package main

import (
	"log"
	"os"

	"github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
	"github.com/mqrc81/zeries/jobs"
	"github.com/mqrc81/zeries/postgres"
	"github.com/mqrc81/zeries/trakt"
)

func main() {
	// Environment variables need to be initialized from .env file first when ran locally
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}

	store, _, err := postgres.Init(os.Getenv("DATABASE_URL"))
	checkError(err)

	tmdbClient, err := tmdb.Init(os.Getenv("TMDB_KEY"))
	checkError(err)

	traktClient, err := trakt.Init(os.Getenv("TRAKT_KEY"))
	checkError(err)

	_ = jobs.NewUpdateReleasesJob(*store, tmdbClient, traktClient).Execute()
	_ = jobs.NewNotifyUsersJob().Execute()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
