package models

type ActInput struct {
	UserID		string	`gorm:"type:text;not null"`
	ExcID		string	`gorm:"type:text;not null"`
	Date		string	`gorm:"type:text;not null"`
	Duration	string	`gorm:"type:text;not null"`
}