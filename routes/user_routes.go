package routes

import (
	"gin-mongo-api/controllers" //add this

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
	router.POST("/users/signup", controllers.SignUp())
	// router.POST("/users/login", controllers.Login())
	router.POST("/users/loginEmail", controllers.LoginWithEmail())
	router.POST("/users/loginMobile", controllers.LoginWithMobile())
}
