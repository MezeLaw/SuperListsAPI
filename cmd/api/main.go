package main

import (
	"SuperListsAPI/cmd/auth/handler"
	"SuperListsAPI/cmd/auth/repository"
	"SuperListsAPI/cmd/auth/service"
	"SuperListsAPI/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDatabase()

	router := gin.Default()
	gin.ForceConsoleColor()

	authRepository := repository.NewAuthRepository(database.AppDatabase)
	authService := service.NewAuthService(&authRepository)
	authHandler := handler.NewAuthHandler(&authService)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/login", authHandler.Login)
			auth.POST("/signup", authHandler.SignUp)
		}

		lists := v1.Group("/lists")
		{
			lists.POST("/", nil)
			lists.GET("/:id", nil)
			lists.PUT("/:id", nil)
			lists.DELETE("/:id", nil)
		}

	}

	router.Run()
}
