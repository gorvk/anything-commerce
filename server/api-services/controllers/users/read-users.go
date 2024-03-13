package controllers

import (
	modelsTypes "github.com/gorvk/anything-commerce/server/api-services/models"
	models "github.com/gorvk/anything-commerce/server/api-services/models/users"
)

func GetAllUsers() []modelsTypes.User {

	var users []modelsTypes.User
	rows, err := models.GetAllUsers()

	if err != nil || rows == nil {
		panic("unable to fetch users")
	}

	defer rows.Close()
	for rows.Next() {
		userRow := modelsTypes.User{}
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
