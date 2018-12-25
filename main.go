package main

import (
	"API_MVC/common"
	"API_MVC/controllers"
	"API_MVC/data"
	"API_MVC/routers"

	"github.com/codegangsta/negroni"
)

func main() {
	//Create a session connect to MongoDB
	clcity, cluser := common.StartUp()
	datacity := data.Col{clcity}
	datauser := data.Col{cluser}

	//Get connection for controller
	controllers.GetCitySession(datacity)
	controllers.GetUserSession(datauser)

	//create router
	router := routers.InitRouters()

	//Add middleware
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(common.Validate))
	n.UseHandler(router)
	n.Run(":8000")
}
