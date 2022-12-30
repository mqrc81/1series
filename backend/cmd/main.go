package main

import (
	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
	"github.com/mqrc81/zeries/controllers"
	"github.com/mqrc81/zeries/env"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/registry"
	"github.com/mqrc81/zeries/sql"
	"time"
)

func main() {
	logger.Info("Starting application...")

	// Register interface adapters
	database := registry.NewDatabase()

	tmdbClient := registry.NewTmdbClient()

	traktClient := registry.NewTraktClient()

	emailClient := registry.NewEmailClient()

	scheduler := registry.NewScheduler(database, tmdbClient, traktClient, emailClient)

	controller := registry.NewController(database, tmdbClient, traktClient, emailClient, scheduler)

	// Start application
	migrateDatabase(database)

	scheduleAndRunJobs(scheduler)

	serveApplication(controller)
}

func migrateDatabase(database *sql.Database) {
	logger.Info("Migrating database")
	logger.FatalOnError(database.Migrate())
}

func scheduleAndRunJobs(scheduler *gocron.Scheduler) {
	logger.Info("Scheduling and running jobs")
	scheduler.StartAsync()
	for _, tag := range env.Config().JobTagsOnInit {
		logger.FatalOnError(scheduler.RunByTagWithDelay(tag, time.Second))
	}
}

func serveApplication(controller controllers.Controller) {
	logger.Info("Listening on " + env.Config().BackendUrl)
	logger.FatalOnError(controller.Start(":" + env.Config().Port))
}
