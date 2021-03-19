package app

import (
	"dstuhack/internal/db"
	"dstuhack/internal/server"
	"dstuhack/pkg/logger"
	"log"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Application interface {
	Run() error
	ReadConfig(path string) error
}

type app struct {
	logger *zerolog.Logger
	db     *db.Database
	server *server.Server
}

func NewApplication() Application {
	app := &app{}

	logg, err := logger.CreateLogger(viper.GetString("logLevel"))
	if err != nil {
		log.Fatal(err.Error())
	}
	app.logger = logg

	database := db.NewDatabase(app.logger)
	app.db = database

	server := server.NewServer(app.db, app.logger)
	app.server = server

	return app
}

func (a *app) ReadConfig(path string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		a.logger.Fatal().Msg(err.Error())
	}
	return err
}

func (a *app) Run() error {

	if err := a.db.Open(); err != nil {
		a.logger.Fatal().Msg(err.Error())
	}
	defer a.db.Close()

	a.db.Migrate(viper.GetString("migrationsPath"))

	a.logger.Info().Msg("Server started at port " + viper.GetString("port"))
	return a.server.Run()
}
