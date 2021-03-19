package server

import (
	"net/http"
)

func (s *Server) baseMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")

		s.logger.Info().Msg(r.URL.Path + ", " + r.Method)
		handler(w, r)
	}
}
