package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mqrc81/zeries/registry"
	. "github.com/mqrc81/zeries/util"
)

// TODO: initialize genres & networks
func main() {
	LogInfo("Starting application...")
	// Environment variables need to be initialized from .env file first when ran locally
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		err := godotenv.Load()
		checkError(err)
	}

	database, err := registry.NewDatabase(os.Getenv("DATABASE_URL"))
	checkError(err)

	sessionManager, err := registry.NewSessionManager(database)
	checkError(err)

	tmdbClient, err := registry.NewTmdbClient(os.Getenv("TMDB_KEY"))
	checkError(err)

	traktClient, err := registry.NewTraktClient(os.Getenv("TRAKT_KEY"))
	checkError(err)

	controller, err := registry.NewController(database, sessionManager, tmdbClient, traktClient)
	checkError(err)

	LogInfo("Listening on " + os.Getenv("BACKEND_URL"))
	err = controller.Start(":" + os.Getenv("PORT"))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		LogPanic(err)
	}
}
