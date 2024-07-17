package models

import "time"

type Event struct {
	ID   int       `json:"user_id"`
	Date time.Time `json:"date" gorm:"not null"`
}
