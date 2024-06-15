package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"` // required post req fields
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e *Event) Save() {
	// TODO: add to DB
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}
