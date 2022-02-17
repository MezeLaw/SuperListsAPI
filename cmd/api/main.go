package main

import (
	"SuperListsAPI/cmd/auth/handler"
	"SuperListsAPI/cmd/auth/repository"
	"SuperListsAPI/cmd/auth/service"
	listHandler "SuperListsAPI/cmd/lists/handler"
	listRepository "SuperListsAPI/cmd/lists/repository"
	listService "SuperListsAPI/cmd/lists/service"
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

	listRepository := listRepository.NewListRepository(database.AppDatabase)
	listService := listService.NewListService(&listRepository)
	listsHandler := listHandler.NewListHandler(&listService, &userListService)

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
			lists.POST("/", listsHandler.Create)
			lists.GET("/:id", listsHandler.Get)
			lists.PUT("/:id", listsHandler.Update)
			lists.DELETE("/:id", listsHandler.Delete)
			lists.POST("/joinList/:listID", listsHandler.JoinList)
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
