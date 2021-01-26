package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllHuman() (models.AllHuman, error) {
	var allHuman models.AllHuman
	var human models.Human

	db, err := driver.ConnectMongoHuman()
	if err != nil {
		return allHuman, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allHuman, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&human)
		allHuman.Human = append(allHuman.Human, human)
	}
	return allHuman, nil
}

func FindHuman(id primitive.ObjectID) (models.Human, error) {
	db, err := driver.ConnectMongoHuman()
	var human models.Human
	if err != nil {
		return human, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&human)
	if err != nil {
		return human, err
	}
	return human, nil
}

func InsertHuman(human models.Human) error {
	collection, err := driver.ConnectMongoHuman()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), human)
	if err != nil {
		return err
	}
	return nil
}

func UpdateHuman(human models.Human, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoHuman()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", human}}
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

func DeleteHuman(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoHuman()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
