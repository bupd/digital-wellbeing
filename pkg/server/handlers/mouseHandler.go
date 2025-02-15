package handlers

import (
	"log"
	"net/http"

	"github.com/bupd/digital-wellbeing/internal/database"
)

func ListAllMouseEvents(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListAllMouse(r.Context())
		if err != nil {
			log.Printf("error in listing all the Mouse Events in DB: %v", err)
			err := &AppError{
				Message: "Error: List All Mouse Events Failed should not use this",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List All Mouse Events success, AVOID THIS, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func ListMousePastHour(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListMouseEventsLastHour(r.Context())
		if err != nil {
			log.Printf("error in listing Past hour mouse events in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Hour Mouse Events Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List past hour mouse events success, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func ListMouseEventsLastDay(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListMouseEventsLast24Hours(r.Context())
		if err != nil {
			log.Printf("error in listing Last Day keys in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Day Mouse Events Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List past day mouse events success, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}
