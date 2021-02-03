package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllBoatBeam() (models.AllBoatBeam, error) {
	var allboatBeam models.AllBoatBeam
	var boatBeam models.BoatBeam

	db, err := driver.ConnectMongoBoatBeam()
	if err != nil {
		return allboatBeam, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allboatBeam, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boatBeam)

		allboatBeam.BoatBeam = append(allboatBeam.BoatBeam, boatBeam)
	}
	return allboatBeam, nil
}

func FindBoatBeam(id primitive.ObjectID) (models.BoatBeam, error) {
	var boatBeam models.BoatBeam
	db, err := driver.ConnectMongoBoatBeam()
	if err != nil {
		return boatBeam, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boatBeam)
	if err != nil {
		return boatBeam, err
	}
	return boatBeam, nil
}

func InsertBoatBeam(boatBeam models.BoatBeam) error {
	collection, err := driver.ConnectMongoBoatBeam()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), boatBeam)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBoatBeam(boatBeam models.BoatBeam, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatBeam()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", boatBeam}}
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

func DeleteBoatBeam(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatBeam()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
