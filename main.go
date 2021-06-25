package main

import (
	"mock-server/controller"
	"mock-server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	controller := controller.NewController()
	e := gin.Default()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	controller.ApplyRoutes(e)
	e.Run(":9000")
}
