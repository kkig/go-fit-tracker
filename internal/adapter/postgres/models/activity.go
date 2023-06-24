package models

type Activity struct {
	ID			string		`gorm:"column:activity_id;type:uuid"`
	UserID		string		`gorm:"type:uuid"`
	ExcID		string		`gorm:"type:text;not null"`
	Duration	int			
	Date		string		`gorm:"type:text;not null"`
}