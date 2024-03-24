package models

import (
	"database/sql"
	"fmt"

	customTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func CreateNewShop(shop customTypes.Shop) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	query := fmt.Sprintf("CALL create_shop('%v', '%v', '%v', '%v', '%v','%v', '%v')",
		shop.OwnerId,
		shop.ShopName,
		shop.Email,
		shop.PhoneNumber,
		shop.MapLocation,
		shop.ShopType,
		shop.ShopDescription)

	rows, err := db.Exec(query)
	return rows, err
}
