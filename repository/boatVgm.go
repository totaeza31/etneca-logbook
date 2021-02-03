package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllBoatVgm() (models.AllBoatVgm, error) {
	var allboatVgm models.AllBoatVgm
	var boatVgm models.BoatVgm

	db, err := driver.ConnectMongoBoatVgm()
	if err != nil {
		return allboatVgm, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allboatVgm, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boatVgm)

		allboatVgm.BoatVgm = append(allboatVgm.BoatVgm, boatVgm)
	}
	return allboatVgm, nil
}

func FindBoatVgm(id primitive.ObjectID) (models.BoatVgm, error) {
	var boatVgm models.BoatVgm
	db, err := driver.ConnectMongoBoatVgm()
	if err != nil {
		return boatVgm, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boatVgm)
	if err != nil {
		return boatVgm, err
	}
	return boatVgm, nil
}

func InsertBoatVgm(boatVgm models.BoatVgm) error {
	collection, err := driver.ConnectMongoBoatVgm()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), boatVgm)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBoatVgm(boatVgm models.BoatVgm, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatVgm()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", boatVgm}}
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

func DeleteBoatVgm(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatVgm()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
