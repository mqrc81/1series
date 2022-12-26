package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mqrc81/zeries/controllers"
	"github.com/mqrc81/zeries/jobs"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/registry"
	"github.com/mqrc81/zeries/sql"
	"time"
)

var (
	Config config
)

type config struct {
	Port          string `env:"PORT"`
	BackendUrl    string `env:"BACKEND_URL"`
	DatabaseUrl   string `env:"DATABASE_URL"`
	TmdbKey       string `env:"TMDB_KEY"`
	TraktKey      string `env:"TRAKT_KEY"`
	RunJobsOnInit bool   `env:"RUN_JOBS_ON_INIT" envDefault:"true"`
}

func main() {
	logger.Info("Starting application...")

	// Load environment variables
	_ = godotenv.Load()
	err := env.Parse(&Config, env.Options{RequiredIfNoDef: true})
	checkError(err)

	// Register interface adapters
	database, err := registry.NewDatabase(Config.DatabaseUrl)
	checkError(err)

	tmdbClient, err := registry.NewTmdbClient(Config.TmdbKey)
	checkError(err)

	traktClient, err := registry.NewTraktClient(Config.TraktKey)
	checkError(err)

	scheduler, err := registry.NewScheduler(database, tmdbClient, traktClient)
	checkError(err)

	controller, err := registry.NewController(database, tmdbClient, traktClient, scheduler)
	checkError(err)

	// Start application
	migrateDatabase(database)
	scheduleAndRunJobs(scheduler)
	serveApplication(controller)
}

func migrateDatabase(database *sql.Database) {
	logger.Info("Migrating database")
	err := database.Migrate()
	checkError(err)
}

func scheduleAndRunJobs(scheduler *gocron.Scheduler) {
	logger.Info("Scheduling and running jobs")
	scheduler.StartAsync()
	if Config.RunJobsOnInit {
		err := scheduler.RunByTagWithDelay(jobs.RunOnInitTag, time.Second)
		checkError(err)
	}
}

func serveApplication(controller controllers.Controller) {
	logger.Info("Listening on " + Config.BackendUrl)
	err := controller.Start(":" + Config.Port)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
