package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/bupd/digital-wellbeing/internal/database"
)

type createUserParams struct {
	Name string `json:"name"`
}

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

func AddUser(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createUserParams
		if err := DecodeRequestBody(r, &req); err != nil {
			log.Println(err)
			HandleAppError(w, err)
			return
		}

		createUserParam := database.CreateUserParams{
			Name: req.Name,
		}
		row, err := db.CreateUser(r.Context(), createUserParam)
		if err != nil {
			log.Printf("error in creating User: %v", err)
			err := &AppError{
				Message: "Error: Create User Failed",
				Code:    http.StatusInternalServerError,
			}
			HandleAppError(w, err)
		}

		WriteJSONResponse(w, http.StatusOK, row)
	}
}
