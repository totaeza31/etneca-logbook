package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllBoatFinance() (models.AllBoatFinance, error) {
	var allboatFinance models.AllBoatFinance
	var boatFinance models.BoatFinance

	db, err := driver.ConnectMongoBoatFinance()
	if err != nil {
		return allboatFinance, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allboatFinance, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boatFinance)

		allboatFinance.BoatFinance = append(allboatFinance.BoatFinance, boatFinance)
	}
	return allboatFinance, nil
}

func FindBoatFinance(id primitive.ObjectID) (models.BoatFinance, error) {
	var boatFinance models.BoatFinance
	db, err := driver.ConnectMongoBoatFinance()
	if err != nil {
		return boatFinance, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boatFinance)
	if err != nil {
		return boatFinance, err
	}
	return boatFinance, nil
}

func InsertBoatFinance(boatFinance models.BoatFinance) error {
	collection, err := driver.ConnectMongoBoatFinance()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), boatFinance)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBoatFinance(boatFinance models.BoatFinance, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatFinance()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", boatFinance}}
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

func DeleteBoatFinance(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatFinance()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
