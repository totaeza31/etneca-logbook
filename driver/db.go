package driver

import (
	"context"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	gotenv.Load()
}

func ConnectMongo() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("USER_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("User").Collection("profile")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Ping(ctx, readpref.Primary())
	client.ListDatabaseNames(ctx, bson.M{})

	return collection, nil
}
