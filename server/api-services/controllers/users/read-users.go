package controllers

import (
	customTypes "github.com/gorvk/anything-commerce/server/api-services/common"
	"github.com/gorvk/anything-commerce/server/api-services/models/user"
)

func GetAllUsers() []customTypes.User {
	var users []customTypes.User
	rows, err := user.GetAllUsers()

	if err != nil || rows == nil {
		panic("unable to fetch users")
	}

	defer rows.Close()
	for rows.Next() {
		userRow := customTypes.User{}
		rows.Scan(
			&userRow.Id,
			&userRow.FirstName,
			&userRow.LastName,
			&userRow.Email,
			&userRow.PhoneNumber,
			&userRow.UserAddress,
			&userRow.IsSeller,
		)
		users = append(users, userRow)
	}

	return users
}

func GetUserByEmail(email string) []customTypes.User {
	var users []customTypes.User
	rows, err := user.GetUserByEmail(email)

	if err != nil || rows == nil {
		panic("unable to fetch users")
	}

	defer rows.Close()
	for rows.Next() {
		userRow := customTypes.User{}
		rows.Scan(
			&userRow.Id,
			&userRow.FirstName,
			&userRow.LastName,
			&userRow.Email,
			&userRow.PhoneNumber,
			&userRow.UserAddress,
			&userRow.IsSeller,
		)
		users = append(users, userRow)
	}

	return users
}
