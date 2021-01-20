package driver

import (
	"context"
	"os"

	"github.com/go-redis/redis"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	return collection, nil
}

func ConnectRedis() (*redis.Client, error) {

	addr := os.Getenv("ADDR_REDIS")
	opt, err := redis.ParseURL(addr)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)
	return client, nil
}
