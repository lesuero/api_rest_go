package main

import (
	"github.com/gin-gonic/gin"
	"../controllers"

)
const (
	port= ":8084"
)

var (
	router = gin.Default()

)
func main() {

	router.GET("/users/:userId",controllers.GetUserFromAPI)
	router.GET("/countries/:countryId",controllers.GetCountryFromAPI)
	router.GET("/sites/:siteId",controllers.GetSiteFromAPI)
	router.GET("/results/:userId",controllers.GetResultFromAPI)
	//router.waitgroup
	//router.chanel
	router.Run(port)
}
