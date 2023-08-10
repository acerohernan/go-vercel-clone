package controllers

import (
	"fmt"
	"net/http"

	"github.com/acerohernan/go-vercel-clone/api-service/config"
	"github.com/acerohernan/go-vercel-clone/api-service/models"
	"github.com/gin-gonic/gin"
)

func UserGetInformation(ctx *gin.Context) {

	usr, exists := ctx.Get("user")

	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	usrId := usr.(config.JWTUser).Id

	var dbUser models.User

	tx := config.DB.First(&dbUser, usrId)

	if tx.Error != nil {
		fmt.Printf("Error at retriving user from database. ERROR: %v", tx.Error)

		ctx.AbortWithStatus(401)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       dbUser.ID,
		"email":    dbUser.Email,
		"username": dbUser.Username,
	})
}

func UserGetGithubRepositories(ctx *gin.Context) {
	usr, exists := ctx.Get("user")

	if !exists {
		ctx.AbortWithStatus(401)
	}

	ctx.JSON(http.StatusOK, usr)
}
