package main

import (
	"github.com/BBrandude/go-auth-client/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	//controllers.Hello()

	r := gin.Default()
	r.POST("/", controllers.CreateAccount)

	r.Run()
}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://cropduster12:Brandude41@cluster0.tbhgr.mongodb.net/example?retryWrites=true&w=majority"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	quickstartDatabase := client.Database("quickstart")
// 	podcastsCollection := quickstartDatabase.Collection("podcasts")
// 	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
// 		{Key: "title", Value: "The Polyglot Developer Podcast"},
// 		{Key: "author", Value: "Nic Raboy"},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(podcastResult.InsertedID)
// }
