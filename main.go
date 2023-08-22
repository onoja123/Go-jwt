package main

import (
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/onoja123/Go-jwt/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.Use(gin.Logger())

	routes.AuthRoute(router)
	routes.UserRoute(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "server successful 1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "server successful 2"})
	})

	router.Run(":" + port)
}
