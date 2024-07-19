package handlers

import (
	"dev11/internal/models"
	"dev11/internal/responses"
	"dev11/internal/storage/postgresql"
	"net/http"
	"strconv"
)

func DeleteEvent(storage *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			responses.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err = storage.DB.Delete(&models.Event{}, id).Error; err != nil {
			responses.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		responses.RespondJSON(w, http.StatusNoContent, map[string]string{"result": "success"})
	}
}
