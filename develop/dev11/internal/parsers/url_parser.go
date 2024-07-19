package parsers

import (
	"dev11/internal/models"
	"net/http"
	"strconv"
	"time"
)

func UrlParser(r *http.Request) (*models.Event, error) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		return nil, err
	}

	return &models.Event{UserID: userID, Date: date}, nil
}
