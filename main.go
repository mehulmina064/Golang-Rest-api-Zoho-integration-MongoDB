package main

import (
	"fmt"
	"gin-mongo-api/configs"
	logger "gin-mongo-api/log"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	// "net/http"
	// "github.com/gin-contrib/static"

	middleware "gin-mongo-api/middleware"
	"io"
	"os"
	"time"
)

	
func main() {

	port := os.Getenv("PORT") 

	if port == "" {
		port = "6000"
	} 

	//set log for debug
	f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout) 


	router := gin.New()
	fmt.Println(gin.Version) 

	//default logger format
	// router.Use(gin.Logger()) 

	//custom logger format
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
    router.Use(gin.Recovery())  
	router.Use(cors.Default())

	// router.Use(middleware.CORSMiddleware())

	 
	//run database
	configs.ConnectDB()

	//log test
	// log.Println(" started server ")
	logger.InfoLogger.Println("--------------------...Starting the application...--------------------")
    // logger.WarningLogger.Println("There is something you should know about")
    // logger.ErrorLogger.Println("Something went wrong")

	//set to v1 api version

	//without auth  routes 
	routes.UserRoute(router) //add this
	

	// Server Test Home Route
	router.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{"success": "Welcome to Prodo"})
		c.String(200, " Welcome to Prodo 55 :)")

	})

	//after this all  authorization routes
	router.Use(middleware.Authentication())
	//auth  routes
	routes.UserAuthRoute(router) 
	routes.TeamAuthRoute(router)
	routes.ProductAuthRoute(router)

	


	// API-2
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for Prodo API"})
	})

	// router.Run("localhost:6000")
	router.Run(":" + port)

}
