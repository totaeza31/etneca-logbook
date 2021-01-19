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

func FindEmail(email string) (models.Authen, error) {
	var authen models.Authen
	db, err := driver.ConnectMongo()
	if err != nil {
		return authen, err
	}
	err = db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&authen)
	if err != nil {
		return authen, err
	}
	return authen, nil
}
