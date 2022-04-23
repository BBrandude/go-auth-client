package main

import (
	"github.com/BBrandude/go-auth-client/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/create", controllers.CreateAccount)
	r.POST("/login", controllers.Login)

	r.Run()
}
