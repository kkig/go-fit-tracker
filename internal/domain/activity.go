package domain

type Activity struct {
	ID			string
	User		User
	Excersize	Excersize
	Duration	int
	Date		string
}