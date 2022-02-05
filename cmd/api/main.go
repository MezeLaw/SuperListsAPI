package main

import (
	"SuperListsAPI/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()

	router := gin.Default()
	gin.ForceConsoleColor()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/login", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "logged",
				})
			})
			auth.POST("/signup", nil)
		}

	}

	router.Run()
}
