package main

import (
	"log"
	"os"

	"github.com/BBrandude/go-auth-client/configs"
	"github.com/BBrandude/go-auth-client/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	configs.SetSessionTimeToLive(500)

	r := gin.Default()
	r.POST("/create", controllers.CreateAccount)
	r.POST("/login", controllers.Login)
	r.GET("/accountInfo", controllers.GetUserName)

	r.Run(os.Getenv("port"))
}
