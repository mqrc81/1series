package api

import (
	"database/sql"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

func NewSessionManager(db *sql.DB) (*scs.SessionManager, error) {
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.NewWithCleanupInterval(db, 1*time.Hour)
	sessionManager.Lifetime = 24 * time.Hour
	return sessionManager, nil
}
