package controllers

import (
	custom_types "github.com/gorvk/anything-commerce/server/api-services/common"
	"github.com/gorvk/anything-commerce/server/api-services/models/userModel"
)

func GetAllUsers() []custom_types.User {
	var users []custom_types.User
	rows, err := userModel.GetAllUsers()

	if err != nil || rows == nil {
		panic("unable to fetch users")
	}

	defer rows.Close()
	for rows.Next() {
		userRow := custom_types.User{}
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

func GetUserByEmail(email string) []custom_types.User {
	var users []custom_types.User
	rows, err := userModel.GetUserByEmail(email)

	if err != nil || rows == nil {
		panic("unable to fetch users")
	}

	defer rows.Close()
	for rows.Next() {
		userRow := custom_types.User{}
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
