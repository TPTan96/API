package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router = SetCityRoutes(router)
	router = SetUserRoutes(router)
	return router
}
