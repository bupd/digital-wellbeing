package handlers

import (
	"log"
	"net/http"

	"github.com/bupd/digital-wellbeing/internal/database"
)

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
				Code:    http.StatusBadRequest,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("Create User success: %v \n", r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, row)
	}
}

func ListUsers(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.ListUsers(r.Context())
		if err != nil {
			log.Printf("error in listing User: %v", err)
			err := &AppError{
				Message: "Error: List User Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("List User success: %v \n", r.UserAgent())
		WriteJSONResponse(w, http.StatusOK, rows)
	}
}
