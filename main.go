package main

import (
	"github.com/BBrandude/go-auth-client/configs"
	"github.com/BBrandude/go-auth-client/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	configs.SetSessionTimeToLive(500)

	r := gin.Default()
	r.POST("/create", controllers.CreateAccount)
	r.POST("/login", controllers.Login)

	r.Run()
}
