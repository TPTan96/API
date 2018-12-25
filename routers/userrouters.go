package routers

import (
	"API_MVC/controllers"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controllers.Login).Methods("GET")
	router.HandleFunc("/signup", controllers.AddNewUser).Methods("POST")
	return router
}
