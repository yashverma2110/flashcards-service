package domain

import "time"

type Deck struct {
	ID int
	TopicId int
	Name string `gorm:"uniqueIndex:deck_name_user_id_idx"`
	Description string
	CreatedBy int `gorm:"uniqueIndex:deck_name_user_id_idx"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Views int
}