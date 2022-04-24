package controllers

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

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

type SessionData struct {
	Cookie    string    `json:"cookie"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
	Timestamp time.Time `json:"timestamp"`
}

var collection *mongo.Collection = configs.GetCollection(configs.DB, "userAccounts")
var userSessions *mongo.Collection = configs.GetCollection(configs.DB, "sessions")

// func randomToken(length int) {
// 	var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
// 	s := make([]rune, n)

// }

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

		newCookie := randomCookieString(32)
		createSession(existingAccount, newCookie)
		c.SetCookie("session", newCookie, 10, "/", "yourDomain", true, false)
		c.String(http.StatusCreated, "login successful")
	}
}

func createSession(accSession Account, sessionCookie string) {
	newSession := SessionData{Cookie: sessionCookie, UserName: accSession.Email, Timestamp: time.Now(), CreatedAt: time.Now()}

	res, err := userSessions.InsertOne(context.TODO(), newSession)
	if err != nil {
		log.Fatal(err)
	}
	_ = res
}

var bytes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randomCookieString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(b)
}
