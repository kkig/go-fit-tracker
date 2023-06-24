package models

type Excersize struct {
	ID		string		`gorm:"column:exc_id;type:uuid"`
	Name	string		`gorm:"type:text;not null"`
}