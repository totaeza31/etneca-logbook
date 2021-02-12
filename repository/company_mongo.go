package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindAllCompany() (models.AllCompany, error) {
	var allcomp models.AllCompany
	var comp models.Company

	db, client, err := driver.ConnectMongoCompany()
	if err != nil {
		return allcomp, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allcomp, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&comp)

		allcomp.Company = append(allcomp.Company, comp)
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return allcomp, err
	}
	return allcomp, nil
}
