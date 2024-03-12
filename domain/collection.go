package domain

import (
	"flashcards/constants"
	"time"
)

type Collection struct {
	ID int
	DeckId int
	Name string `gorm:"uniqueIndex:name_user_id_idx"`
	CreatedBy int `gorm:"uniqueIndex:name_user_id_idx"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Visibility constants.Visibility
}