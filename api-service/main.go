package main

import (
	"os"

	"github.com/acerohernan/go-vercel-clone/api-service/config"
	"github.com/acerohernan/go-vercel-clone/api-service/controllers"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback", "user:email"),
	)

	r := gin.Default()

	// Auth routes
	r.GET("/auth/:provider", controllers.AuthGetProvider)
	r.GET("/auth/:provider/callback", controllers.AuthGetProviderCallback)

	// User routes

	r.Run()
}
