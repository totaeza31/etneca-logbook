package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllBoatGateway() (models.AllBoatGateway, error) {
	var allboatGateway models.AllBoatGateway
	var boatGateway models.BoatGateway

	db, err := driver.ConnectMongoBoatGateway()
	if err != nil {
		return allboatGateway, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allboatGateway, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boatGateway)

		allboatGateway.BoatGateway = append(allboatGateway.BoatGateway, boatGateway)
	}
	return allboatGateway, nil
}

func FindBoatGateway(id primitive.ObjectID) (models.BoatGateway, error) {
	var boatGateway models.BoatGateway
	db, err := driver.ConnectMongoBoatGateway()
	if err != nil {
		return boatGateway, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boatGateway)
	if err != nil {
		return boatGateway, err
	}
	return boatGateway, nil
}

func InsertBoatGateway(boatGateway models.BoatGateway) error {
	collection, err := driver.ConnectMongoBoatGateway()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), boatGateway)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBoatGateway(boatGateway models.BoatGateway, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatGateway()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", boatGateway}}
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

func DeleteBoatGateway(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoatGateway()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
