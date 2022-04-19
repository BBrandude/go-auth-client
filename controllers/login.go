package controllers

import (
	"context"
	"net/http"

	"github.com/BBrandude/go-auth-client/configs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type accountLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var collection *mongo.Collection = configs.GetCollection(configs.DB, "userAccounts")

func Login(c *gin.Context) {
	var loginAttemptInfo accountLoginInfo
	var existingAccount accountLoginInfo

	if err := c.BindJSON(&loginAttemptInfo); err != nil {
		c.String(http.StatusInternalServerError, "invalid data submitted")
		return
	}

	filter := bson.M{"email": loginAttemptInfo.Email}

	doesAccountExist := collection.FindOne(context.TODO(), filter).Decode(&existingAccount)
	if doesAccountExist == mongo.ErrNoDocuments {
		c.String(http.StatusCreated, "account does not exist")
		return
	}
	doesPasswordMatch := bcrypt.CompareHashAndPassword([]byte(existingAccount.Password), []byte(loginAttemptInfo.Password))
	if doesPasswordMatch == bcrypt.ErrMismatchedHashAndPassword {
		c.String(http.StatusCreated, "account does not exist")
	} else {
		c.String(http.StatusCreated, "login successful")
	}

}
