package initializers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client

func ConnectToDb() {
	var err error
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("MONGODB_URI not set.")
	}

	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic("cannot connect to the monogodb cluster.")
	}

	//defer func() {
	//	if err = Client.Disconnect(c); err != nil {
	//		panic(err)
	//	}
	//}()
}
