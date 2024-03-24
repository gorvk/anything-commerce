package routes

import (
	"net/http"

	controllers "github.com/gorvk/anything-commerce/server/api-services/controllers/shop"
)

func init() {
	http.HandleFunc("/create-shop", controllers.CreateNewShop)
	http.HandleFunc("/get-shop", controllers.GetAllShops)
	http.HandleFunc("/get-shop-by-email", controllers.GetShopByEmail)
	http.HandleFunc("/delete-shop-by-email", controllers.DeleteShopByEmail)
	http.HandleFunc("/update-shop-by-email", controllers.UpdateShopByEmail)
}
