package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllTech() (models.AllTech, error) {
	var allTech models.AllTech
	var tech models.Tech

	db, err := driver.ConnectMongoTech()
	if err != nil {
		return allTech, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allTech, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&tech)
		allTech.Tech = append(allTech.Tech, tech)
	}
	return allTech, nil
}

func FindTech(id primitive.ObjectID) (models.Tech, error) {
	db, err := driver.ConnectMongoTech()
	var tech models.Tech
	if err != nil {
		return tech, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&tech)
	if err != nil {
		return tech, err
	}
	return tech, nil
}

func InsertTech(tech models.Tech) error {
	collection, err := driver.ConnectMongoTech()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), tech)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTech(tech models.Tech, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoTech()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", tech}}
	_, err = db.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteTech(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoTech()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
