package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mqrc81/zeries/controllers"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/registry"
	"github.com/mqrc81/zeries/sql"
	"time"
)

var (
	config struct {
		Port                string   `env:"PORT"`
		BackendUrl          string   `env:"BACKEND_URL"`
		DatabaseUrl         string   `env:"DATABASE_URL"`
		TmdbKey             string   `env:"TMDB_KEY"`
		TraktKey            string   `env:"TRAKT_KEY"`
		SendGridKey         string   `env:"SENDGRID_KEY"`
		SendGridSenderEmail string   `env:"SENDGRID_SENDER_EMAIL"`
		JobTagsOnInit       []string `env:"JOB_TAGS_ON_INIT" envDefault:""`
	}
)

func main() {
	logger.Info("Starting application...")

	// Load environment variables
	_ = godotenv.Load()
	err := env.Parse(&config, env.Options{RequiredIfNoDef: true})
	checkError(err)

	// Register interface adapters
	database, err := registry.NewDatabase(config.DatabaseUrl)
	checkError(err)

	tmdbClient, err := registry.NewTmdbClient(config.TmdbKey)
	checkError(err)

	traktClient, err := registry.NewTraktClient(config.TraktKey)
	checkError(err)

	emailClient, err := registry.NewEmailClient(config.SendGridKey, config.SendGridSenderEmail)
	checkError(err)

	scheduler, err := registry.NewScheduler(database, tmdbClient, traktClient, emailClient)
	checkError(err)

	controller, err := registry.NewController(database, tmdbClient, traktClient, emailClient, scheduler)
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
	for _, tag := range config.JobTagsOnInit {
		err := scheduler.RunByTagWithDelay(tag, time.Second)
		checkError(err)
	}
}

func serveApplication(controller controllers.Controller) {
	logger.Info("Listening on " + config.BackendUrl)
	err := controller.Start(":" + config.Port)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
