package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func UpdateUserByEmail(email string, columnName string, newValue string) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("CALL update_user('%v', '%v', '%v', '%v')", columnName, "email", newValue, email)
	rows, err := db.Exec(query)
	return rows, err
}
