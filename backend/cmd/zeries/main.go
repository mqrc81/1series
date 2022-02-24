// Package cmd/zeries is the entry-point which starts the application
// and initializes database, external clients, list of genres/networks
package main

import (
	"os"

	"github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
	"github.com/mqrc81/zeries/api"
	"github.com/mqrc81/zeries/postgres"
	"github.com/mqrc81/zeries/trakt"
	"github.com/mqrc81/zeries/util"
)

// TODO: initialize genres & networks
func main() {
	util.LogInfo("Starting application...")
	// Environment variables need to be initialized from .env file first when ran locally
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		err := godotenv.Load()
		checkError(err)
	}

	store, sessionStore, err := postgres.NewStore(os.Getenv("DATABASE_URL"))
	checkError(err)

	tmdbClient, err := tmdb.Init(os.Getenv("TMDB_KEY"))
	checkError(err)

	traktClient, err := trakt.Init(os.Getenv("TRAKT_KEY"))
	checkError(err)

	handler, err := api.NewHandler(*store, sessionStore, tmdbClient, traktClient)
	checkError(err)

	util.LogInfo("Listening on " + os.Getenv("BACKEND_URL"))
	err = handler.Start(":" + os.Getenv("PORT"))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		util.LogPanic(err)
	}
}
