package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BadRequestError : Bad request response agar tidak repeat
func BadRequestError(ctx *gin.Context, msg string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": msg,
	})
}
