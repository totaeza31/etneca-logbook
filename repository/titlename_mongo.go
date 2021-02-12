package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindAllTitle() (models.AllTitle, error) {
	var allTitle models.AllTitle
	var title models.Title

	db, client, err := driver.ConnectMongoTitle()
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

	err = client.Disconnect(context.Background())

	if err != nil {
		return allTitle,err
	}
	return allTitle, nil
}
