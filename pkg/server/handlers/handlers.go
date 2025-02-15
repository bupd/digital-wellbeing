package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func Health(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := db.PingContext(r.Context())
		if err != nil {
			slog.Error(fmt.Sprintf("error in pinging DB: %v", err))
			_, _ = w.Write([]byte("unhealthy"))
			return
		}
		slog.Debug(fmt.Sprintf("ping DB success: %v", r))
		log.Printf("ping DB success: %v \n", r.UserAgent())
		_, _ = w.Write([]byte("healthy"))
	}
}
