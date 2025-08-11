package router

import (
	"gochatx/internal/user/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	UserHandler handler.UserHandler
}

func (routes UserRoutes) GetAllUses(ctx *gin.Context) {
	if users, err := routes.UserHandler.Uc.GetAll(); err == nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"users": users,
		})
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

}

func (routes UserRoutes) GetByUsername(ctx *gin.Context) {
	// var res dtos.UserResponseDTO
	username := ctx.Param("username")
	user1, err := routes.UserHandler.Uc.GetByUsername(username)
	if err != nil {

		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data": user1,
	})

}
