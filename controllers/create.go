package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/BBrandude/go-auth-client/configs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var collection *mongo.Collection = configs.GetCollection(configs.DB, "userAccounts")

func CreateAccount(c *gin.Context) {
	//
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingAccount Account
	var newAccount Account

	if err := c.BindJSON(&newAccount); err != nil {
		c.String(http.StatusInternalServerError, "invalid credentials")
		return
	}

	filter := bson.M{"email": newAccount.Email}

	err := collection.FindOne(context.TODO(), filter).Decode(&existingAccount)
	if err == mongo.ErrNoDocuments {

		fmt.Println("does not exist, creating acc")

		insertedAccount, err := collection.InsertOne(ctx, newAccount)
		if err != nil {
			c.String(http.StatusInternalServerError, "internal error")
			return
		}
		fmt.Println(insertedAccount)
		c.String(http.StatusCreated, "Account successfully created")
	} else {
		c.String(http.StatusCreated, "Account under "+existingAccount.Email+" already exists")
	}
}
