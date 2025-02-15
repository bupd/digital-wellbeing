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
	r.HandleFunc("/user", h.AddUser(s.dbQueries)).Methods("POST")
	r.HandleFunc("/user", h.ListUsers(s.dbQueries)).Methods("GET")
	r.HandleFunc("/user", h.DeleteUser(s.dbQueries)).Methods("DELETE")

	// keys event handler
	r.HandleFunc("/keys/all", h.ListAllKeys(s.dbQueries)).Methods("GET")
	r.HandleFunc("/keys/1hr", h.ListKeysPastHour(s.dbQueries)).Methods("GET")
	r.HandleFunc("/keys/1day", h.ListKeysPastDay(s.dbQueries)).Methods("GET")

	// keys event handler
	r.HandleFunc("/mouse/all", h.ListAllKeys(s.dbQueries)).Methods("GET")
	r.HandleFunc("/mouse/1hr", h.ListKeysPastHour(s.dbQueries)).Methods("GET")
	return r
}
