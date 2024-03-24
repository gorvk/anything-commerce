package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func DeleteProductById(id int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("CALL delete_product('%v', '%v')", "id", id)
	rows, err := db.Query(query)
	return rows, err
}
