package handlers

import (
	"log"
	"net/http"

	"github.com/bupd/digital-wellbeing/internal/database"
)

func AddKey(db *database.Queries) http.HandlerFunc {
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
