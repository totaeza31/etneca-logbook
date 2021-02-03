package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllBoatType() (models.AllBoatType, error) {
	var allboatType models.AllBoatType
	var boatType models.BoatType

	db, err := driver.ConnectMongoBoatType()
	if err != nil {
		return allboatType, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allboatType, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boatType)
		
		allboatType.BoatType = append(allboatType.BoatType, boatType)
	}
	return allboatType, nil
}

func FindBoatType(id primitive.ObjectID) (models.BoatType, error) {
	var boatType models.BoatType
	db, err := driver.ConnectMongoBoatType()
	if err != nil {
		return boatType, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boatType)
	if err != nil {
		return boatType, err
	}
	return boatType, nil
}

func InsertBoatType(boatType models.BoatType) error {
	collection, err := driver.ConnectMongoBoatType()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), boatType)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBoatType(boatType models.BoatType, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatType()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", boatType}}
	_, err = db.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBoatType(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatType()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
