package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mqrc81/zeries/jobs"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/registry"
	"os"
	"time"
)

func main() {
	logger.Info("Starting application...")

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
	logger.Info("Scheduling and running jobs")
	scheduler.StartAsync()
	err = scheduler.RunByTagWithDelay(jobs.RunOnInitTag, time.Second)
	checkError(err)

	logger.Info("Listening on " + os.Getenv("BACKEND_URL"))
	err = controller.Start(":" + os.Getenv("PORT"))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
