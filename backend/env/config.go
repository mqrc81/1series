package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/mqrc81/1series/logger"
	_ "github.com/mqrc81/1series/sql"
	"sync"
)

var (
	once   = new(sync.Once)
	config EnvironmentVariables
)

type EnvironmentVariables struct {
	Port                string   `env:"PORT"`
	BackendUrl          string   `env:"BACKEND_URL"`
	FrontendUrl         string   `env:"FRONTEND_URL"`
	DatabaseUrl         string   `env:"DATABASE_URL"`
	TmdbKey             string   `env:"TMDB_KEY"`
	TraktKey            string   `env:"TRAKT_KEY"`
	SendGridKey         string   `env:"SENDGRID_KEY"`
	SendGridSenderEmail string   `env:"SENDGRID_SENDER_EMAIL"`
	JobTagsOnInit       []string `env:"JOB_TAGS_ON_INIT" envDefault:""`
	Admins              []string `env:"ADMINS" envDefault:""`
}

func Config() EnvironmentVariables {
	once.Do(func() {
		_ = godotenv.Load()
		err := env.Parse(&config, env.Options{RequiredIfNoDef: true})
		if err != nil {
			logger.FatalOnError(err)
		}
	})
	return config
}
