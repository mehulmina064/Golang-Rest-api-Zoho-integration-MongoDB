package main

import (
	"gin-mongo-api/configs"
	"fmt"
	"gin-mongo-api/routes"
	logger "gin-mongo-api/log"


	"github.com/gin-gonic/gin"
	// "net/http"
	// "github.com/gin-contrib/static"

	middleware "gin-mongo-api/middleware"
	"log"
	"os"
)

func main() {

	port := os.Getenv("PORT") 

	if port == "" {
		port = "6000"
	} 
	router := gin.New()
	fmt.Println(gin.Version)
	router.Use(gin.Logger()) 
	// router.Use(static.Serve("/", static.LocalFile("./../prodo-internal-vue/dist", false)))


	//run database
	configs.ConnectDB()

	//log test
	// log.Println(" started server ")
	logger.InfoLogger.Println("--------------------...Starting the application...--------------------")
    // logger.WarningLogger.Println("There is something you should know about")
    // logger.ErrorLogger.Println("Something went wrong")


	//without auth  routes 
	routes.UserRoute(router) //add this

	// API-1
	router.GET("/", func(c *gin.Context) {

		// c.JSON(200, gin.H{"success": "Welcome to Prodo"})
		c.String(200, " Welcome to Prodo :)")

	})

	//after this all  authorization routes
	router.Use(middleware.Authentication())

	//auth  routes
	routes.UserAuthRoute(router) 


	// API-2
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for Prodo API"})
	})

	// router.Run("localhost:6000")
	router.Run(":" + port)

}
