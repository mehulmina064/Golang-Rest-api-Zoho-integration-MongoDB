package main

import (
        "gin-mongo-api/configs"
        "gin-mongo-api/routes"
        "github.com/gin-gonic/gin"
)

func main() {
        router := gin.Default()
        
        //run database
        configs.ConnectDB()
    
        //routes
        routes.UserRoute(router) //add this
    
        router.Run("localhost:6000")
    }