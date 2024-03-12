package domain

import "time"

type Card struct {
	ID int
	DeckId int
	Title string
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
}