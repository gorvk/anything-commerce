package routes

import (
	"net/http"

	controllers "github.com/gorvk/anything-commerce/server/api-services/controllers/order"
)

func init() {
	http.HandleFunc("/create-order", controllers.CreateNewOrder)
	http.HandleFunc("/get-order", controllers.GetAllOrders)
	http.HandleFunc("/get-order-by-id", controllers.GetOrderById)
	http.HandleFunc("/delete-order-by-id", controllers.DeleteOrderById)
	http.HandleFunc("/update-order-by-id", controllers.UpdateOrderById)
}
