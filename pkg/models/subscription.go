package models

type Subscription struct {
	User   Users
	Course Course
	Price  int
}
