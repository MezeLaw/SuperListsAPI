package main

import (
	"SuperListsAPI/cmd/api/middleware"
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
			lists.POST("/", middleware.ValidateJWTOnRequest, listsHandler.Create)
			lists.GET("/:id", middleware.ValidateJWTOnRequest, listsHandler.Get)
			lists.PUT("/:id", middleware.ValidateJWTOnRequest, listsHandler.Update)
			lists.DELETE("/:id", middleware.ValidateJWTOnRequest, listsHandler.Delete)
			lists.POST("/joinList/:listID", middleware.ValidateJWTOnRequest, listsHandler.JoinList)
		}

		userLists := v1.Group("/userLists")
		{
			userLists.POST("/", middleware.ValidateJWTOnRequest, userListHandler.Create)
			userLists.GET("/:id", middleware.ValidateJWTOnRequest, userListHandler.Get)
			userLists.GET("/", middleware.ValidateJWTOnRequest, userListHandler.GetUserListsByUserID)
			userLists.DELETE("/:id", middleware.ValidateJWTOnRequest, userListHandler.Delete)
		}

	}

	router.Run()
}
