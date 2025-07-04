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

func TopWinLastDay(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.TopDurationWinLastDay(r.Context())
		if err != nil {
			log.Printf("error in listing active Window Events Duration by day in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Day Active Windows Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List Active Last Day Window Events success: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func TopWinLastHour(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.TopDurationWinLastHour(r.Context())
		if err != nil {
			log.Printf("error in listing active Window Events Duration by hour in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Hour Active Windows Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List Active Last Hour Window Events success: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func TopWinLastDayActive(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.TopActiveDurationWinLastDay(r.Context())
		if err != nil {
			log.Printf("error in listing active Window Events Duration by day in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Day Active Windows Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List Active Last Day Window Events success: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func TopWinLastHourActive(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.TopActiveDurationWinLastHour(r.Context())
		if err != nil {
			log.Printf("error in listing active Window Events Duration by hour in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Hour Active Windows Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List Active Last Hour Window Events success: %v, %v \n", len(rows), r.UserAgent())
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
