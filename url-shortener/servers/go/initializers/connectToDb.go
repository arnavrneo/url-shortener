package initializers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectToDb for connecting to MongoDB
func ConnectToDb(uri string) error {
	var err error

	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	return err

	//defer func() {
	//	if err = Client.Disconnect(c); err != nil {
	//		panic(err)
	//	}
	//}()
}
