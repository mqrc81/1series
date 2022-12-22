package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mqrc81/zeries/job"
	"github.com/mqrc81/zeries/registry"
	. "github.com/mqrc81/zeries/util"
	"os"
	"time"
)

func main() {
	LogInfo("Starting application...")

	// Initialize local environment variables
	if os.Getenv("ENVIRONMENT") != "RAILWAY" {
		err := godotenv.Load()
		checkError(err)
	}

	// Register interface adapters
	database, err := registry.NewDatabase(os.Getenv("DATABASE_URL"))
	checkError(err)

	tmdbClient, err := registry.NewTmdbClient(os.Getenv("TMDB_KEY"))
	checkError(err)

	traktClient, err := registry.NewTraktClient(os.Getenv("TRAKT_KEY"))
	checkError(err)

	scheduler, err := registry.NewScheduler(database, tmdbClient, traktClient)
	checkError(err)

	controller, err := registry.NewController(database, tmdbClient, traktClient)
	checkError(err)

	// Start application
	LogInfo("Scheduling and running jobs")
	scheduler.StartAsync()
	err = scheduler.RunByTagWithDelay(job.RunOnInitTag, time.Second)
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
