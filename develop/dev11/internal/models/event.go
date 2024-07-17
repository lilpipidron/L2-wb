package models

import "time"

type Event struct {
	ID   int
	Date time.Time `gorm:"not null"`
}
