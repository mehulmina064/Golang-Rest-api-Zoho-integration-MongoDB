package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"

	middleware "gin-mongo-api/middleware"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "6000"
	}
	// router := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this
	//place authorization routes
	router.Use(middleware.Authentication())

	// API-2
	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	// API-1
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	// router.Run("localhost:6000")
	router.Run(":" + port)

}
