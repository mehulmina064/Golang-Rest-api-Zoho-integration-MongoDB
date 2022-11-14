package main

import (
	"gin-mongo-api/configs"
	"fmt"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
	// "net/http"
	// "github.com/gin-contrib/static"

	middleware "gin-mongo-api/middleware"
	"log"
	"os"
)



var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	port := os.Getenv("PORT") 

	if port == "" {
		port = "6000"
	} 
	// router := gin.Default()
	router := gin.New()
	fmt.Println(gin.Version)
	router.Use(gin.Logger()) 
	// router.Use(static.Serve("/", static.LocalFile("./../prodo-internal-vue/dist", false)))


	//run database
	configs.ConnectDB()

	//without auth  routes
	routes.UserRoute(router) //add this


	//log test
	InfoLogger.Println("Starting the application...")
    InfoLogger.Println("Something noteworthy happened")
    WarningLogger.Println("There is something you should know about")
    ErrorLogger.Println("Something went wrong")


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
