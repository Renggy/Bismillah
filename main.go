package main

import (
	"bismillah/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/registration", handler.HandleRegistration)
		api.POST("/status_inqury", handler.HandlerStatusInquiry)
		api.POST("/payment", handler.HandlePayment)
	}

	router.Run(":8080")
}
