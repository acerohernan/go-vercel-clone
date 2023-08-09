package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func AuthGetProvider(c *gin.Context) {

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func AuthGetProviderCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusAccepted, user)
}
