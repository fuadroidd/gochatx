package router

import (
	"fmt"
	"gochatx/internal/user"
	"gochatx/internal/user/dtos"
	"gochatx/internal/user/handler"
	"gochatx/internal/user/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(engine *gin.Engine, db *gorm.DB) {
	useRepo := repo.UserRepositoryImpl{Db: db}
	userHandler := handler.UserHandler{
		Uc: &user.UserUsecase{Repo: &useRepo},
	}

	engine.POST("register/", func(ctx *gin.Context) {
		var usr dtos.UserRequestDTO
		if err := ctx.ShouldBind(&usr); err != nil {
			ctx.IndentedJSON(http.StatusNoContent, gin.H{
				"error": fmt.Errorf("not created %s: ", err.Error()),
			})
			print("what the heck", err.Error)
		}
		print(usr.Password)
		if err := userHandler.Register(usr); err != nil {
			ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("user not created %s", err),
			})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"message": "successfully created",
		})
	})

	engine.GET("users/", func(ctx *gin.Context) {
		// var res dtos.UserResponseDTO
		user1, err := userHandler.Uc.GetByUsername("user1")
		if err != nil {

			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"data": user1,
		})

	})
}
