package handlers

import (
	"dev11/internal/parsers"
	"dev11/internal/responses"
	"dev11/internal/storage/postgresql"
	"net/http"
)

func CreateEvent(storage *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := parsers.ParseBody(r)
		if err != nil {
			responses.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err = storage.DB.Create(&event).Error; err != nil {
			responses.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		responses.RespondJSON(w, http.StatusCreated, map[string]string{"result": "event created successfully"})
	}
}
