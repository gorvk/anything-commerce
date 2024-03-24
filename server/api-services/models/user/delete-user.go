package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func DeleteUserByEmail(email string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("CALL delete_user('%v', '%v'::text)", "email", email)
	rows, err := db.Query(query)
	return rows, err
}
