package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func UpdateOrderById(id int, columnName string, newValue string) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("CALL update_order('%v', '%v', '%v', '%v')", columnName, "id", newValue, id)
	rows, err := db.Exec(query)
	return rows, err
}
