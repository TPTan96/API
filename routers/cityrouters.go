package routers

import (
	"API_MVC/controllers"

	"github.com/gorilla/mux"
)

func SetCityRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/city", controllers.GetAllDatas).Methods("GET")
	router.HandleFunc("/city/{id}", controllers.GetAData).Methods("GET")
	router.HandleFunc("/city", controllers.CreateAData).Methods("POST")
	router.HandleFunc("/city/{id}", controllers.DeleteAData).Methods("DELETE")
	router.HandleFunc("/city/{id}", controllers.UpdateAData).Methods("PATCH")
	return router
}
