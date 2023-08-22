package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/onoja123/Go-jwt/controller"
)

func AuthRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signup", controller.Signup()) // Changed to controller.Signup
	incomingRoutes.POST("/user/login", controller.Login())   // Changed to controller.Login
}
