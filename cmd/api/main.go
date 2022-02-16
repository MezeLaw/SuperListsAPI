package main

import (
	"SuperListsAPI/cmd/auth/handler"
	"SuperListsAPI/cmd/auth/repository"
	"SuperListsAPI/cmd/auth/service"
	userListHandler "SuperListsAPI/cmd/userLists/handler"
	userListRepository "SuperListsAPI/cmd/userLists/repository"
	userListService "SuperListsAPI/cmd/userLists/service"
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

	userListRepository := userListRepository.NewUserListRepository(database.AppDatabase)
	userListService := userListService.NewUserListService(&userListRepository)
	userListHandler := userListHandler.NewUserListHandler(&userListService)

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

		userLists := v1.Group("/userLists")
		{
			userLists.POST("/", userListHandler.Create)
			userLists.GET("/:id", userListHandler.Get)
			userLists.GET("/", userListHandler.GetUserListsByUserID)
			userLists.DELETE("/:id", userListHandler.Delete)
		}

	}

	router.Run()
}
