package parsers

import (
	"dev11/internal/models"
	"net/http"
	"strconv"
	"time"
)

func ParseBody(r *http.Request) (*models.Event, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", r.Form.Get("date"))
	if err != nil {
		return nil, err
	}
	return &models.Event{ID: userID, Date: date}, nil
}
