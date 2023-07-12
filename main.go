package main

import (
	registration "bismillah/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/registration", registration.HandleRegistration)
	}

	router.Run(":8080")
}
