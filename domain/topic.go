package domain

type Topic struct {
	ID   int
	Name string `gorm:"type:varchar(50);uniqueIndex:name_idx"`
}