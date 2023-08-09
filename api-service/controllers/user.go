package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInformation(ctx *gin.Context) {

	usr, exists := ctx.Get("user")

	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.JSON(http.StatusOK, usr)
}
