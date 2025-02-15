package handlers

import (
	"net/http"

	"github.com/bupd/digital-wellbeing/internal/database"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func Health(dbQueries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("healthy"))
	}
}
