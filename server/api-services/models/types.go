package models

type User struct {
	Id          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	UserAddress string
	IsSeller    bool
}
