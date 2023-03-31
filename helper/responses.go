package helper

import (
	"github.com/gin-gonic/gin"
)

type ResponseErrorStruct struct {
	Message string
	Status  int
}

func ResponseError(ctx *gin.Context, msg string, status int) {
	ctx.AbortWithStatusJSON(status, gin.H{
		"message": msg,
	})
}
