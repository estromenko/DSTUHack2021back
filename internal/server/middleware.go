package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func (s *Server) baseMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")

		s.logger.Info().Msg(r.URL.Path + ", " + r.Method + ", " + r.RemoteAddr)
		handler(w, r)
	}
}

func (s *Server) AuthenticationMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unprocessedToken string

		head := string(r.Header.Get("Authorization"))
		splited := strings.Split(head, " ")
		if len(splited) != 2 {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Token is not provided",
			})
			return
		}
		unprocessedToken = splited[1]

		token, err := jwt.Parse(unprocessedToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method, got " + t.Header["alg"].(string))
			}
			return []byte(viper.GetString("jwt_secret")), nil
		})

		if err != nil {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		user, err := s.db.User().FindByID(int(id))
		if err != nil {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		handler(w, r.WithContext(ctx))
	}
}
