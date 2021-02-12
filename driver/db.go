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

func ConnectMongoProfile() (*mongo.Collection, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("USER_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("User").Collection("profile")

	return collection, client, nil
}

func ConnectMongoPackage() (*mongo.Collection, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("USER_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("User").Collection("package")

	return collection, client, nil
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

func ConnectMongoReport() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("report")

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

func ConnectMongoBoat() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("boat")

	return collection, nil
}

func ConnectMongoBoatType() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("boatType")

	return collection, nil
}

func ConnectMongoBoatBeam() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("boatBeamStatus")

	return collection, nil
}

func ConnectMongoBoatDevice() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("boatDeviceStatus")

	return collection, nil
}

func ConnectMongoBoatFinance() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("boatFinancialStatus")

	return collection, nil
}

func ConnectMongoBoatGateway() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("boatGateway")

	return collection, nil
}

func ConnectMongoBoatVgm() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("boatVmsGen")

	return collection, nil
}

func ConnectMongoCompany() (*mongo.Collection, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("back_office").Collection("company")

	return collection, client, nil
}

func ConnectMongoGender() (*mongo.Collection, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("back_office").Collection("gender")

	return collection, client, nil
}

func ConnectMongoPosition() (*mongo.Collection, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("back_office").Collection("position")

	return collection, client, nil
}

func ConnectMongoEmp() (*mongo.Collection, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("back_office").Collection("employee")

	return collection, client, nil
}

func ConnectMongoGoods() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("goods")

	return collection, nil
}

func ConnectMongoWorksheet() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("back_office").Collection("worksheet")

	return collection, nil
}

func ConnectMongoTitle() (*mongo.Collection, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("BO_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("back_office").Collection("titlename")

	return collection, client, nil
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
