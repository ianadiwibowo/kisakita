package entity

import "time"

// Story is a collection of paragraphs
type Story struct {
	ID         int
	Title      string
	Synopsis   string
	Paragraphs []Paragraph
	Admins     []Writer
	Authors    []Writer
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
