package entities

import "time"

// Paragraph is part of a story and can be written by different authors
type Paragraph struct {
	Author    Writer
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
