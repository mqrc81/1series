package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mqrc81/zeries/registry"
	. "github.com/mqrc81/zeries/util"
)

func main() {
	// Environment variables need to be initialized from .env file first when ran locally
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		err := godotenv.Load()
		checkError(err)
	}

	database, err := registry.NewDatabase(os.Getenv("DATABASE_URL"))
	checkError(err)

	tmdbClient, err := registry.NewTmdbClient(os.Getenv("TMDB_KEY"))
	checkError(err)

	traktClient, err := registry.NewTraktClient(os.Getenv("TRAKT_KEY"))
	checkError(err)

	err = registry.NewUpdateReleasesJob(database, tmdbClient, traktClient).Execute()
	checkError(err)

	err = registry.NewNotifyUsersJob().Execute()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		LogPanic(err)
	}
}
