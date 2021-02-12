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

	db, client, err := driver.ConnectMongoTech()
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
	err = client.Disconnect(context.Background())

	if err != nil {
		return allTech, err
	}
	return allTech, nil
}

func FindTech(id primitive.ObjectID) (models.Tech, error) {
	db, client, err := driver.ConnectMongoTech()
	var tech models.Tech
	if err != nil {
		return tech, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&tech)
	if err != nil {
		return tech, err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return tech, err
	}
	return tech, nil
}

func FindTechName(company string) (models.Tech, error) {
	db, client, err := driver.ConnectMongoTech()
	var tech models.Tech
	if err != nil {
		return tech, err
	}
	err = db.FindOne(context.TODO(), bson.M{"company": company}).Decode(&tech)
	if err != nil {
		return tech, err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return tech, err
	}
	return tech, nil
}

func InsertTech(tech models.Tech) error {
	collection, client, err := driver.ConnectMongoTech()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), tech)
	if err != nil {
		return err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func UpdateTech(tech models.Tech, id primitive.ObjectID) error {
	db, client, err := driver.ConnectMongoTech()
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
		err = client.Disconnect(context.Background())

		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteTech(id primitive.ObjectID) error {
	db, client, err := driver.ConnectMongoTech()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return err
}
