package entity

// Story is a collection of paragraphs
type Story struct {
	Title      string
	Synopsis   string
	Paragraphs []Paragraph
	Admins     []Writer
	Authors    []Writer
}
