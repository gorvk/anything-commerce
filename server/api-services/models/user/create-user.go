package models

import (
	"database/sql"
	"fmt"

	customTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func CreateNewUser(user customTypes.User) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	query := fmt.Sprintf("CALL create_user('%v', '%v', '%v', '%v', '%v')",
		user.Email,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.UserAddress)

	rows, err := db.Exec(query)
	return rows, err
}
