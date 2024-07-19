package models

import "time"

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id" gorm:"not null"`
	Date   time.Time `json:"date" gorm:"not null"`
}
