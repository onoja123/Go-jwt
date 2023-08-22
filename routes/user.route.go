package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/onoja123/Go-jwt/controller"
	"github.com/onoja123/Go-jwt/middleware"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.Login())
}
