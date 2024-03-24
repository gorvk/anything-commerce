package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func GetAllShops() (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	rows, err := db.Query("SELECT * from read_all_shops()")
	return rows, err
}

func GetShopByEmail(email string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * from read_shop_by_column('%v', '%v'::TEXT)", "email", email)
	rows, err := db.Query(query)
	return rows, err
}
