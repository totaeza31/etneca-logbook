package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllPosition() (models.AllPosition, error) {
	var allPosition models.AllPosition
	var position models.Position

	db, err := driver.ConnectMongoPosition()
	if err != nil {
		return allPosition, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allPosition, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&position)

		allPosition.Position = append(allPosition.Position, position)
	}
	return allPosition, nil
}

func FindPosition(id primitive.ObjectID) (models.Position, error) {
	var position models.Position
	db, err := driver.ConnectMongoCompany()
	if err != nil {
		return position, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&position)
	if err != nil {
		return position, err
	}
	return position, nil
}

func InsertPosition(position models.Position) error {
	collection, err := driver.ConnectMongoCompany()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), position)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePosition(position models.Position, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoCompany()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", position}}
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

func DeletePosition(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoCompany()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
