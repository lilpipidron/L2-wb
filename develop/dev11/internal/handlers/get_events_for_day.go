package handlers

import (
	"dev11/internal/models"
	"dev11/internal/parsers"
	"dev11/internal/responses"
	"dev11/internal/storage/postgresql"
	"net/http"
)

func GetEventsForDay(storage *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := parsers.UrlParser(r)
		if err != nil {
			responses.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		var events []models.Event
		query := storage.DB.Where("user_id = ? AND date >= ? AND date < ?::date + INTERVAL '1 day'", event.UserID, event.Date, event.Date)
		if err = query.Find(&events).Error; err != nil {
			responses.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		responses.RespondJSON(w, http.StatusOK, map[string][]models.Event{"result": events})
	}
}
