package queue

import "time"

type Ticket struct {
	ID      string
	UserID  string
	EventID string
	Date    time.Time
}

type User struct {
	ID   string
	Name string
}
