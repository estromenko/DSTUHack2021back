package server

import (
	"dstuhack/internal/db"
	"dstuhack/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Server struct {
	logger *zerolog.Logger
	db     *db.Database

	userService *services.UserService
	apiService  *services.APIService
}

func NewServer(db *db.Database, logger *zerolog.Logger) *Server {
	return &Server{
		db:     db,
		logger: logger,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"ok": "ok",
		})
	})

	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/reg", s.baseMiddleware(s.RegisterUser())).Methods("POST")
	auth.HandleFunc("/login", s.baseMiddleware(s.LoginUser())).Methods("POST")

	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/info", s.baseMiddleware(s.AuthenticationMiddleware(s.GetUserInfo()))).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/tickers", s.baseMiddleware(s.AuthenticationMiddleware(s.GetAllTickers()))).Methods("GET")
	api.HandleFunc("/tickers/stocks", s.baseMiddleware(s.AuthenticationMiddleware(s.GetAllSymbolStocks()))).Methods("GET")
	api.HandleFunc("/operation", s.baseMiddleware(s.AuthenticationMiddleware(s.BuyOrSellStoke()))).Methods("POST")

	return http.ListenAndServe(":"+viper.GetString("port"), router)
}

func (s *Server) User() *services.UserService {
	if s.userService == nil {
		s.userService = services.NewUserService(s.db)
	}
	return s.userService
}

func (s *Server) API() *services.APIService {
	if s.apiService == nil {
		s.apiService = services.NewAPIService(viper.GetString("api_access_key"))
	}
	return s.apiService
}
