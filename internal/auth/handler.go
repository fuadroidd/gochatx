package auth

import (
	"fmt"
	user "gochatx/internal/user"
	"gochatx/internal/user/dtos"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuth struct {
	Uc user.UserUsecase
}

// Login
func (uc UserAuth) Login(ctx *gin.Context) {
	//get username and password send to the authenticator
	var user dtos.UserLoginRequestDTO
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	ctx.Bind(&user)
	// 🔒 You should verify user credentials from DB here
	user4mdb, _ := uc.Uc.GetByUsername(user.Username)
	print(user4mdb.Username)
	if user.Username != user4mdb.Username || user.Password != "password1" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// Register as a NewUser
func (Ua UserAuth) Register(ctx *gin.Context) {
	var usr dtos.UserRequestDTO
	if err := ctx.ShouldBind(&usr); err != nil {
		ctx.IndentedJSON(http.StatusNoContent, gin.H{
			"error": fmt.Errorf("not created %s: ", err.Error()),
		})
	}

	ue := user.UserEntity{
		Username:    usr.Username,
		Displayname: usr.Displayname,
		Email:       usr.Email,
	}

	if err := Ua.Uc.Register(ue, usr.Password); err != nil {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": fmt.Sprintf("user not created %s", err),
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "successfully created",
	})
}
