package handlers

import (
	"log"
	"net/http"

	"github.com/bupd/digital-wellbeing/internal/database"
)

func ListAllKeys(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListAllKeys(r.Context())
		if err != nil {
			log.Printf("error in listing all the keys in DB: %v", err)
			err := &AppError{
				Message: "Error: List All Keys Failed should not use this",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List All Keys success, AVOID THIS, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func ListKeysPastHour(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListKeysPressedLastHour(r.Context())
		if err != nil {
			log.Printf("error in listing Past hour keys in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Hour Keys Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List Past Hour Keys success, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}

func ListKeysPastDay(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListKeysPressedLast24Hours(r.Context())
		if err != nil {
			log.Printf("error in listing Last Day keys in DB: %v", err)
			err := &AppError{
				Message: "Error: List Past Day Keys Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List Past Day Keys success, Length: %v, %v \n", len(rows), r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}
