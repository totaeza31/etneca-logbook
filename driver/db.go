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

func ConnectMongoProfile() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("USER_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("User").Collection("profile")

	return collection, nil
}

func ConnectMongoPackage() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("USER_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("User").Collection("package")

	return collection, nil
}

func ConnectMongoBO() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("owner")

	return collection, nil
}

func ConnectMongoTech() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("tech")

	return collection, nil
}

func ConnectMongoHuman() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("human")

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
