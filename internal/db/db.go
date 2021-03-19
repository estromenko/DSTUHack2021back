package db

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Database struct {
	logger *zerolog.Logger
	db     *sql.DB
}

func NewDatabase(logger *zerolog.Logger) *Database {
	return &Database{
		logger: logger,
	}
}

func (d *Database) Open() error {
	db, err := sql.Open("postgres", viper.GetString("dsn"))
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	d.db = db
	d.logger.Info().Msg("Database connection opened successfully.")
	return nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Migrate(migrationsPath string) error {
	d.logger.Info().Msg("Running migrations ...")
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		d.logger.Error().Msg(err.Error())
		return err
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(migrationsPath + "/" + file.Name())

		if err != nil {
			d.logger.Error().Msg(err.Error())
			return err
		}

		if _, err := d.db.Exec(string(data)); err != nil {
			d.logger.Error().Msg(err.Error())
			return err
		}

		d.logger.Info().Msg("- " + file.Name() + ": done.")
	}
	d.logger.Info().Msg("Migrated successfully.")
	return nil
}
