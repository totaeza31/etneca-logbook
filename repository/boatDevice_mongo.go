package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllBoatDevice() (models.AllBoatDevice, error) {
	var allboatDevice models.AllBoatDevice
	var boatDevice models.BoatDevice

	db, err := driver.ConnectMongoBoatDevice()
	if err != nil {
		return allboatDevice, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allboatDevice, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boatDevice)

		allboatDevice.BoatDevice = append(allboatDevice.BoatDevice, boatDevice)
	}
	return allboatDevice, nil
}

func FindBoatDevice(id primitive.ObjectID) (models.BoatDevice, error) {
	var boatDevice models.BoatDevice
	db, err := driver.ConnectMongoBoatDevice()
	if err != nil {
		return boatDevice, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boatDevice)
	if err != nil {
		return boatDevice, err
	}
	return boatDevice, nil
}

func InsertBoatDevice(boatDevice models.BoatDevice) error {
	collection, err := driver.ConnectMongoBoatDevice()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), boatDevice)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBoatDevice(boatDevice models.BoatDevice, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatDevice()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", boatDevice}}
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

func DeleteBoatDevice(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatDevice()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
