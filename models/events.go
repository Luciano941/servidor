package models

import "server/db"

type Event struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AllEvents []Event

func MigrateEvent() {
	db.Database().AutoMigrate(Event{})
}
