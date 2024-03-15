package user

import (
	"database/sql"
	"fmt"

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

func GetUserByEmail(email string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * from read_user_by_column('%v', '%v'::TEXT)", "email", email)
	rows, err := db.Query(query)
	return rows, err
}
