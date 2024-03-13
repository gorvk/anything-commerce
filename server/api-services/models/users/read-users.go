package models

import (
	"database/sql"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func GetAllUsers() (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	rows, err := db.Query("SELECT * from read_all_users()")
	return rows, err
}
