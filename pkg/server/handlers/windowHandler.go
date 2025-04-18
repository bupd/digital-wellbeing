package handlers

import (
	"log"
	"net/http"

	"github.com/bupd/digital-wellbeing/internal/database"
)

func ListAllWindowEvents(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListAllWmclass(r.Context())
		if err != nil {
			log.Printf("error in listing all the Window Events in DB: %v", err)
			err := &AppError{
				Message: "Error: List All Window Events Failed should not use this",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List All Window Events success, AVOID THIS, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func ListWindowPastHour(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListLastHourWmClass(r.Context())
		if err != nil {
			log.Printf("error in listing all the Window Events in DB: %v", err)
			err := &AppError{
				Message: "Error: List All Window Events Failed should not use this",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List All Window Events success, AVOID THIS, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func ListWindowPastDay(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListLastDayWmClass(r.Context())
		if err != nil {
			log.Printf("error in listing all the Window Events in DB: %v", err)
			err := &AppError{
				Message: "Error: List All Window Events Failed should not use this",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List All Window Events success, AVOID THIS, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}
