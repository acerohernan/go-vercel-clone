package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/acerohernan/go-vercel-clone/api-service/config"
	"github.com/acerohernan/go-vercel-clone/api-service/models"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
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
	jwtUsr, exists := ctx.Get("user")

	if !exists {
		ctx.AbortWithStatus(401)

		return
	}

	userId := jwtUsr.(config.JWTUser).Id

	var usr *models.User

	tx := config.DB.First(usr, userId)

	if tx.Error != nil {

	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	githubCtx := context.Background()

	client := github.NewClient(httpClient)

	repos, _, err := client.Repositories.List(githubCtx, "", &github.RepositoryListOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 1,
		},
	})

	if err != nil {
		fmt.Printf("Error at retriving user repositories from github. ERROR: %v", err)

		ctx.AbortWithStatus(401)

		return
	}

	ctx.JSON(http.StatusOK, repos)
}
