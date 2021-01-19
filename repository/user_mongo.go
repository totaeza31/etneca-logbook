package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUser(id primitive.ObjectID) (models.User, error) {
	var user models.User
	db, err := driver.ConnectMongo()
	if err != nil {
		return user, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
