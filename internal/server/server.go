package server

import (
	"dstuhack/internal/db"
	"dstuhack/internal/services"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Server struct {
	logger *zerolog.Logger
	db     *db.Database

	userService *services.UserService
}

func NewServer(db *db.Database, logger *zerolog.Logger) *Server {
	return &Server{
		db:     db,
		logger: logger,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/reg", s.baseMiddleware(s.RegisterUser())).Methods("POST")

	return http.ListenAndServe(":"+viper.GetString("port"), router)
}

func (s *Server) User() *services.UserService {
	if s.userService == nil {
		s.userService = services.NewUserService(s.db)
	}
	return s.userService
}
