package routes

import (
	"net/http"

	controllers "github.com/gorvk/anything-commerce/server/api-services/controllers/product"
)

func init() {
	http.HandleFunc("/create-product", controllers.CreateNewProduct)
	http.HandleFunc("/get-product", controllers.GetAllProducts)
	http.HandleFunc("/get-product-by-id", controllers.GetProductById)
	http.HandleFunc("/delete-product-by-id", controllers.DeleteProductById)
	http.HandleFunc("/update-product-by-id", controllers.UpdateProductById)
}
