package domain

import "time"

// User defines structure for the domain User
type User struct{
	ID int
	Username string `gorm:"type:varchar(50);uniqueIndex:username_idx"`
	Password string
	Email string `gorm:"type:varchar(100);uniqueIndex:email_idx"`
	FirstName string
	LastName string
	OAuthMode string
	Birthday time.Time
}
