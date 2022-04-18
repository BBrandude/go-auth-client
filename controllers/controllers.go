package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Account struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateAccount(c *gin.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("mongoURI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	userData := client.Database("userData")
	userAccounts := userData.Collection("userAccounts")

	//

	var existingAccount Account
	var newAccount Account

	if err := c.BindJSON(&newAccount); err != nil {
		fmt.Println(err)
	}

	filter := bson.M{"email": newAccount.Email}

	err = userAccounts.FindOne(context.TODO(), filter).Decode(&existingAccount)
	if err == mongo.ErrNoDocuments {

		fmt.Println("does not exist, creating acc")

		insertedAccount, err := userAccounts.InsertOne(ctx, newAccount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedAccount)
		c.String(http.StatusCreated, "Account successfully created")
	} else {
		c.String(http.StatusCreated, "Account under "+existingAccount.Email+" already exists")
	}
}
