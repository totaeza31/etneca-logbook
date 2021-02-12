package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindAllPosition() (models.AllPosition, error) {
	var allPosition models.AllPosition
	var position models.Position

	db, client, err := driver.ConnectMongoPosition()
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
	err = client.Disconnect(context.Background())

	if err != nil {
		return allPosition, err
	}
	return allPosition, nil
}
