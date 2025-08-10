package main

import (
	"fmt"
	"gochatx/internal/auth"
	"gochatx/internal/infrastructure/db"
	"gochatx/internal/infrastructure/db/migrations"
	"gochatx/internal/user"
	userRepo "gochatx/internal/user/repo"

	"gochatx/internal/user/handler"

	"gochatx/internal/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func mainRoutes(db *gorm.DB) {
	userRepo := userRepo.UserRepositoryImpl{Db: db}
	userUsecase := user.UserUsecase{Repo: &userRepo}
	userHandler := handler.UserHandler{Uc: &userUsecase}
	userRoute := router.UserRoutes{
		UserHandler: userHandler,
	}
	route := gin.Default()
	// not authenticated
	userAuth := auth.UserAuth{Uc: userUsecase}
	open := route.Group("/api")
	open.POST("/login", userAuth.Login)
	open.POST("/register", userAuth.Register)
	open.GET("/users", userRoute.GetAllUses)
	//authenticated
	protected := route.Group("/api")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/profile", userRoute.GetByUsername)
		protected.GET("/user/:username", userRoute.GetByUsername)
	}
	if err := route.Run("localhost:8000"); err != nil {
		fmt.Print(err.Error())
	}
}

func main() {
	dbcon, err := db.Connect()

	if err != nil {
		fmt.Print(fmt.Errorf("unable to connect database %s", err))
	} else {

		if err := migrations.AutoMigrate(dbcon); err != nil {
			fmt.Print(fmt.Errorf("migrations error %s", err))
		}
	}
	mainRoutes(dbcon)
}
