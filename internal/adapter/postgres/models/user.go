package models

// Tablers is interface of GORM
type Tabler interface {
	TableName() string
}

type User struct {
	ID		string	`gorm:"column:user_id;type:uuid"`
	Name	string	`gorm:"type:text;not null"`
}

func (User) TableName() string {
	return "users"
}