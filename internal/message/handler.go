package message

import (
	msgDTOs "gochatx/internal/message/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	UseCase MessageUsecase
}

func (mh MessageHandler) SendMessage(ctx *gin.Context) {
	var msgRequestDTO msgDTOs.MessageRequestDTO
	if err := ctx.Bind(&msgRequestDTO); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "",
			"error":   err,
		})
	} else {
		if err := mh.UseCase.SendMessage(msgRequestDTO); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "",
				"error":   err,
			})
		} else {
			ctx.IndentedJSON(http.StatusCreated, gin.H{
				"message": "message successfully sent",
				"error":   "",
			})
		}
	}
}
func (mh MessageHandler) GetMessageByID(ctx *gin.Context) {
	if ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "wrong id type",
			"message": "",
		})
	} else {
		if message, err := mh.UseCase.GetByID(uint(ID)); err != nil {
			ctx.IndentedJSON(404, gin.H{
				"message": "",
				"error":   err,
			})
		} else {
			ctx.IndentedJSON(200, gin.H{
				"message": message,
				"error":   err,
			})
		}
	}
}
