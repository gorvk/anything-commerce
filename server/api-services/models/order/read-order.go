package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func GetAllOrders() (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	rows, err := db.Query("SELECT * from read_all_orders()")
	return rows, err
}

func GetOrderById(id int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * from read_order_by_column('%v', '%v')", "id", id)
	rows, err := db.Query(query)
	return rows, err
}
