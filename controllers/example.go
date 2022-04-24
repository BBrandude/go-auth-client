package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/BBrandude/go-auth-client/configs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var existingSessions *mongo.Collection = configs.GetCollection(configs.DB, "sessions")
var foundSession SessionData

func GetUserName(c *gin.Context) {
	reqCookie, err := c.Cookie("session")
	if err == http.ErrNoCookie {

		c.String(http.StatusCreated, "invalid cookie header")
		return
	}
	fmt.Println(reqCookie)
	sessionLookupFilter := bson.M{"cookie": reqCookie}
	sessionLookup := existingSessions.FindOne(context.TODO(), sessionLookupFilter).Decode(&foundSession)
	if sessionLookup == mongo.ErrNoDocuments {
		c.String(http.StatusCreated, "invalid sesison")
		return
	}
	c.String(http.StatusCreated, foundSession.UserName)
	fmt.Println(foundSession)
}
