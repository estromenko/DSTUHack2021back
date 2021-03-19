package server

import (
	"dstuhack/internal/db"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Server struct {
	logger *zerolog.Logger
	db     *db.Database
}

func NewServer(db *db.Database, logger *zerolog.Logger) *Server {
	return &Server{
		db:     db,
		logger: logger,
	}
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.baseMiddleware(s.handler()))
	return http.ListenAndServe(":"+viper.GetString("port"), nil)
}
