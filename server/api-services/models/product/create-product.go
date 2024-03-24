package models

import (
	"database/sql"
	"fmt"

	dbTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func CreateNewProduct(Product dbTypes.Product) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	query := fmt.Sprintf("CALL create_product('%v', '%v', '%v', '%v', '%v','%v', '%v', '%v', '%v')",
		Product.ProductName,
		Product.SellerId,
		Product.ShopId,
		Product.ProductType,
		Product.ProductCondition,
		Product.Price,
		Product.OriginalPurchasedDate,
		Product.OriginalPurchaisingRecieptNo,
		Product.ProductDescription)

	rows, err := db.Exec(query)
	return rows, err
}
