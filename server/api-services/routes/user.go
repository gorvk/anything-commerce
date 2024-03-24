package routes

import (
	"net/http"

	controllers "github.com/gorvk/anything-commerce/server/api-services/controllers/user"
)

func init() {
	http.HandleFunc("/create-user", controllers.CreateNewUser)
	http.HandleFunc("/get-user", controllers.GetAllUsers)
	http.HandleFunc("/get-user-by-email", controllers.GetUserByEmail)
	http.HandleFunc("/delete-user-by-email", controllers.DeleteUserByEmail)
	http.HandleFunc("/update-user-by-email", controllers.UpdateUserByEmail)
}
