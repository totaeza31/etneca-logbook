package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllTitle() (models.AllTitle, error) {
	var allTitle models.AllTitle
	var title models.Title

	db, err := driver.ConnectMongoTitle()
	if err != nil {
		return allTitle, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allTitle, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&title)

		allTitle.Title = append(allTitle.Title, title)
	}
	return allTitle, nil
}

func FindTitle(id primitive.ObjectID) (models.Title, error) {
	var title models.Title
	db, err := driver.ConnectMongoTitle()
	if err != nil {
		return title, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&title)
	if err != nil {
		return title, err
	}
	return title, nil
}

func InsertTitle(title models.Title) error {
	collection, err := driver.ConnectMongoTitle()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), title)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTitle(title models.Title, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoTitle()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", title}}
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

func DeleteTitle(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoTitle()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
