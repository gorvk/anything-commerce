package models

import (
	"database/sql"
	"fmt"

	dbTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	"github.com/gorvk/anything-commerce/server/api-services/initializers"
)

func CreateNewOrder(Order dbTypes.Order) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	query := fmt.Sprintf("CALL create_order('%v', '%v', '%v', '%v', '%v','%v', '%v', '%v')",
		Order.FromMapLocation,
		Order.ToMapLocation,
		Order.LastStopMapLocation,
		Order.OrderStatus,
		Order.PaymentStatus,
		Order.ProductId,
		Order.BuyerId,
		Order.ShopId)

	rows, err := db.Exec(query)
	return rows, err
}
