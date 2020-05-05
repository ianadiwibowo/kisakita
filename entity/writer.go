package entity

import "time"

// Writer is the author of a single paragraph
type Writer struct {
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
