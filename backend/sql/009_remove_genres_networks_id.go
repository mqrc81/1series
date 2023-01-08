package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upRemoveGenresNetworksId, downRemoveGenresNetworksId)
}

//goland:noinspection SqlResolve
func upRemoveGenresNetworksId(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE genres DROP CONSTRAINT genres_pkey;
		ALTER TABLE networks DROP CONSTRAINT networks_pkey;
		ALTER TABLE genres ALTER COLUMN id TYPE INT;
		ALTER TABLE networks ALTER COLUMN id TYPE INT;
		UPDATE genres SET id = tmdb_id;
		UPDATE networks SET id = tmdb_id;
		ALTER TABLE genres DROP COLUMN tmdb_id;
		ALTER TABLE networks DROP COLUMN tmdb_id;
		ALTER TABLE genres ALTER COLUMN id SET NOT NULL;
		ALTER TABLE networks ALTER COLUMN id SET NOT NULL;
		ALTER TABLE genres RENAME COLUMN id TO genre_id;
		ALTER TABLE networks RENAME COLUMN id TO network_id;
		`)
	return err
}

func downRemoveGenresNetworksId(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE genres RENAME COLUMN genre_id TO id;
		ALTER TABLE networks RENAME COLUMN network_id TO id;
		ALTER TABLE genres ALTER COLUMN id TYPE SERIAL;
		ALTER TABLE networks ALTER COLUMN id TYPE SERIAL;
		ALTER TABLE genres ADD CONSTRAINT genres_pkey PRIMARY KEY (id);
		ALTER TABLE networks ADD CONSTRAINT networks_pkey PRIMARY KEY (id);
		ALTER TABLE genres ADD COLUMN tmdb_id INT NOT NULL DEFAULT id;
		ALTER TABLE networks ADD COLUMN tmdb_id INT NOT NULL DEFAULT id;
		`)
	return err
}
