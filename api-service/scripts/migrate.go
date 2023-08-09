package main

import (
	"github.com/acerohernan/go-vercel-clone/api-service/config"
	"github.com/acerohernan/go-vercel-clone/api-service/models"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
}
