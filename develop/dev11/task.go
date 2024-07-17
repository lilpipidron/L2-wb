package main

import (
	"dev11/internal/storage/postgresql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")

	storage, err := postgresql.NewPostgresDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/create_event", nil)
	http.Handle("/update_event", nil)
	http.Handle("/delete_event", nil)
	http.Handle("/events_for_day", nil)
	http.Handle("/events_for_week", nil)
	http.Handle("/events_for_month", nil)

	log.Print("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
