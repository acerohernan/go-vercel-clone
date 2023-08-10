package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/acerohernan/go-vercel-clone/api-service/config"
	"github.com/acerohernan/go-vercel-clone/api-service/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth/gothic"
	"gorm.io/gorm"
)

func AuthGetProvider(ctx *gin.Context) {
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func AuthGetProviderCallback(ctx *gin.Context) {
	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)

	if err != nil {
		fmt.Printf("Error at authenticating user with goth. ERROR: %v", err)

		ctx.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	// Lookup of the user exists in database
	var usr *models.User

	tx := config.DB.First(&usr, models.User{Email: user.Email})

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			// Create the user in the database
			newUsr := &models.User{Email: user.Email, Username: user.Email}
			result := config.DB.Create(newUsr)

			if result.Error != nil {
				fmt.Printf("Error at creating new user in the database. ERROR: %v", result.Error)

				ctx.AbortWithStatus(http.StatusInternalServerError)

				return
			}

			usr = newUsr
		} else {
			fmt.Printf("Error at querying user in the database. ERROR: %v", tx.Error)

			ctx.AbortWithStatus(http.StatusInternalServerError)

			return
		}
	}

	// Create the jwt session
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
		"sub":   usr.ID,
		"email": usr.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Printf("Error at creating jwt token. ERROR: %v", err)

		ctx.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
