package routes

import (
	"gin-mongo-api/controllers" 

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	// router.POST("/teams/signup", controllers.SignUp())
	// router.POST("/teams/loginEmail", controllers.LoginWithEmail())
	// router.POST("/teams/loginMobile", controllers.LoginWithMobile())
}

func ProductAuthRoute(router *gin.Engine) {
	router.POST("/product", controllers.CreateProduct())
	router.GET("/product/:productId", controllers.GetProduct())
	router.PUT("/product/:productId", controllers.EditProduct())
	router.DELETE("/product/:productId", controllers.DeleteProduct())
	router.GET("/products", controllers.GetAllProduct())
}
