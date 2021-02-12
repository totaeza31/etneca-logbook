package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindAllGender() (models.AllGender, error) {
	var allGender models.AllGender
	var gender models.Gender

	db, client, err := driver.ConnectMongoGender()
	if err != nil {
		return allGender, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allGender, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&gender)

		allGender.Gender = append(allGender.Gender, gender)
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return allGender, err
	}
	return allGender, nil
}
