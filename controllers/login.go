package controllers

import (
	"context"
	"fmt"
	"math/rand"
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

type sessionData struct {
	Cookie   string `json:"cookie"`
	UserName string `json:"userName"`
}

var collection *mongo.Collection = configs.GetCollection(configs.DB, "userAccounts")
var userSessions *mongo.Collection = configs.GetCollection(configs.DB, "sessions")

// func randomToken(length int) {
// 	var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
// 	s := make([]rune, n)

// }
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Login(c *gin.Context) {
	var loginAttemptInfo accountLoginInfo
	var existingAccount Account

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
		c.String(http.StatusCreated, "incorrect password")
	} else {
		fmt.Println(existingAccount)
		newSession := sessionData{Cookie: "325", UserName: existingAccount.Email}
		session, err := userSessions.InsertOne(context.TODO(), newSession)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(session)
		c.SetCookie("cookieName", "name", 10, "/", "yourDomain", true, false)
		c.String(http.StatusCreated, "login successful")
	}

	fmt.Println(RandStringBytes(25))
}
