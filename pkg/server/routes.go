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

	// mouse events handler
	r.HandleFunc("/mouse/all", h.ListAllMouseEvents(s.dbQueries)).Methods("GET")
	r.HandleFunc("/mouse/1hr", h.ListMousePastHour(s.dbQueries)).Methods("GET")
	r.HandleFunc("/mouse/1day", h.ListMouseEventsLastDay(s.dbQueries)).Methods("GET")

	// window events handler
	r.HandleFunc("/window/all", h.ListAllWindowEvents(s.dbQueries)).Methods("GET")
	r.HandleFunc("/window/1hr", h.ListWindowPastHour(s.dbQueries)).Methods("GET")
	r.HandleFunc("/window/1day", h.ListWindowPastDay(s.dbQueries)).Methods("GET")
	r.HandleFunc("/window/1day/active/top", h.TopWinLastDayActive(s.dbQueries)).Methods("GET")
	r.HandleFunc("/window/1hr/active/top", h.TopWinLastHourActive(s.dbQueries)).Methods("GET")
	r.HandleFunc("/window/1day/top", h.TopWinLastDay(s.dbQueries)).Methods("GET")
	r.HandleFunc("/window/1hr/top", h.TopWinLastHour(s.dbQueries)).Methods("GET")

	return r
}
