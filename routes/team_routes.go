package routes

import (
	"gin-mongo-api/controllers" 

	"github.com/gin-gonic/gin"
)

func TeamRoute(router *gin.Engine) {
	// router.POST("/teams/signup", controllers.SignUp())
	// router.POST("/teams/loginEmail", controllers.LoginWithEmail())
	// router.POST("/teams/loginMobile", controllers.LoginWithMobile())
}

func TeamAuthRoute(router *gin.Engine) {
	router.POST("/team", controllers.CreateTeam())
	router.GET("/team/:teamId", controllers.GetTeam())
	router.PUT("/team/:teamId", controllers.EditTeam())
	router.DELETE("/team/:teamId", controllers.DeleteTeam())
	router.GET("/teams", controllers.GetAllTeams())
}
