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

<<<<<<< HEAD
	userRepo *UserRepo
=======
	userRepo      *UserRepo
>>>>>>> main
	operationRepo *OperationRepo
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
 
// Создаем экземпляр операции
func (d *Database) OpretaionRepoCreate() *OperationRepo {
	if d.operationRepo == nil {
		d.operationRepo = NewOperationRepo(d.db)
	}
	return d.operationRepo
}

// Создаем экземпляр пользователь
func (d *Database) User() *UserRepo {
	if d.userRepo == nil {
		d.userRepo = NewUserRepo(d.db)
	}
	return d.userRepo
}

func (d *Database) Operation() *OperationRepo {
	if d.operationRepo == nil {
		d.operationRepo = NewOperationRepo(d.db)
	}
	return d.operationRepo
}
