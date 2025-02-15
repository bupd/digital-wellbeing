package server

import (
	"net/http"

	h "github.com/bupd/digital-wellbeing/pkg/server/handlers"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/ping", h.Ping).Methods("GET")
	r.HandleFunc("/health", h.Health(s.db)).Methods("GET")

	return r
}
