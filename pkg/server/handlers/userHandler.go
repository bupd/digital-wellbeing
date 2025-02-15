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

		row, err := db.CreateUser(r.Context(), req.Name)
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

func DeleteUser(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createUserParams
		if err := DecodeRequestBody(r, &req); err != nil {
			log.Println(err)
			HandleAppError(w, err)
			return
		}

		err := db.DelteUser(r.Context(), req.Name)
		if err != nil {
			log.Printf("error in deleting User: %v", err)
			err := &AppError{
				Message: "Error: Delete User Failed",
				Code:    http.StatusNotFound,
			}
			HandleAppError(w, err)
			return
		}

		log.Printf("Delete User success: %v \n", r.UserAgent())
		WriteJSONResponse(w, http.StatusNoContent, "{}")
	}
}
