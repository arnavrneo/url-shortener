package initializers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectToDb for connecting to MongoDB
func ConnectToDb(uri string) {
	var err error

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
