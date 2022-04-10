package main

import (
	"SuperListsAPI/cmd/api/middleware"
	"SuperListsAPI/cmd/auth/handler"
	"SuperListsAPI/cmd/auth/repository"
	"SuperListsAPI/cmd/auth/service"
	listItemHandler "SuperListsAPI/cmd/listItems/handler"
	listItemRepository "SuperListsAPI/cmd/listItems/repository"
	listItemService "SuperListsAPI/cmd/listItems/service"
	listHandler "SuperListsAPI/cmd/lists/handler"
	listRepository "SuperListsAPI/cmd/lists/repository"
	listService "SuperListsAPI/cmd/lists/service"
	userListHandler "SuperListsAPI/cmd/userLists/handler"
	userListRepository "SuperListsAPI/cmd/userLists/repository"
	userListService "SuperListsAPI/cmd/userLists/service"
	userHandler "SuperListsAPI/cmd/users/handler"
	userRepository "SuperListsAPI/cmd/users/repository"
	userService "SuperListsAPI/cmd/users/service"
	"SuperListsAPI/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	database.InitDatabase()

	router := gin.Default()
	//router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders: []string{"*", "*"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}))
	gin.ForceConsoleColor()

	authRepository := repository.NewAuthRepository(database.AppDatabase)
	authService := service.NewAuthService(&authRepository)
	authHandler := handler.NewAuthHandler(&authService)

	userRepository := userRepository.NewUsersRepository(database.AppDatabase)
	userService := userService.NewUserService(&userRepository)
	userHandler := userHandler.NewUserHandler(&userService)

	userListRepository := userListRepository.NewUserListRepository(database.AppDatabase)
	userListService := userListService.NewUserListService(&userListRepository)
	userListHandler := userListHandler.NewUserListHandler(&userListService)

	listItemRepository := listItemRepository.NewListItemRepository(database.AppDatabase)
	listItemService := listItemService.NewListItemService(&listItemRepository)
	listItemHandler := listItemHandler.NewListItemHandler(&listItemService)

	listRepository := listRepository.NewListRepository(database.AppDatabase)
	listService := listService.NewListService(&listRepository)
	listsHandler := listHandler.NewListHandler(&listService, &userListService, &listItemService)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{ //TODO cambiar los tests de login por POST
			auth.POST("/login", authHandler.Login)
			auth.POST("/signup", authHandler.SignUp)
		}

		users := v1.Group("/users")
		{
			users.GET("/:email", middleware.ValidateJWTOnRequest, userHandler.Get)
		}

		lists := v1.Group("/lists")
		{
			lists.POST("/", middleware.ValidateJWTOnRequest, listsHandler.Create)
			lists.GET("/:id", middleware.ValidateJWTOnRequest, listsHandler.Get)
			lists.GET("/", middleware.ValidateJWTOnRequest, listsHandler.GetLists)
			lists.PUT("/:id", middleware.ValidateJWTOnRequest, listsHandler.Update)
			lists.DELETE("/:id", middleware.ValidateJWTOnRequest, listsHandler.Delete)
			lists.POST("/joinList/:inviteCode", middleware.ValidateJWTOnRequest, listsHandler.JoinList)
		}

		userLists := v1.Group("/userLists")
		{
			userLists.POST("/", middleware.ValidateJWTOnRequest, userListHandler.Create)
			userLists.GET("/:id", middleware.ValidateJWTOnRequest, userListHandler.Get)
			userLists.GET("/", middleware.ValidateJWTOnRequest, userListHandler.GetUserListsByUserID)
			userLists.DELETE("/:id", middleware.ValidateJWTOnRequest, userListHandler.Delete)
		}

		listItems := v1.Group("/listItems")
		{
			listItems.POST("/", middleware.ValidateJWTOnRequest, listItemHandler.Create)
			listItems.GET("/:id", middleware.ValidateJWTOnRequest, listItemHandler.Get)
			listItems.PUT("/:id", middleware.ValidateJWTOnRequest, listItemHandler.Update)
			listItems.DELETE("/:id", middleware.ValidateJWTOnRequest, listItemHandler.Delete)
			listItems.POST("/bulkDelete", middleware.ValidateJWTOnRequest, listItemHandler.BulkDelete)
		}

	}

	err := router.Run()
	if err != nil {
		panic("Error running sv!")
	}
}
