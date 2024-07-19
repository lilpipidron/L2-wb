package main

import (
	"dev11/internal/handlers"
	"dev11/internal/middleware"
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
	http.Handle("/create_event", middleware.LoggingMiddleware(handlers.CreateEvent(storage)))
	http.Handle("/update_event", middleware.LoggingMiddleware(handlers.UpdateEvent(storage)))
	http.Handle("/delete_event", middleware.LoggingMiddleware(handlers.DeleteEvent(storage)))
	http.Handle("/events_for_day", middleware.LoggingMiddleware(nil))
	http.Handle("/events_for_week", middleware.LoggingMiddleware(nil))
	http.Handle("/events_for_month", middleware.LoggingMiddleware(nil))

	log.Print("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
