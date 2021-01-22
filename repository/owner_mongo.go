package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllOwner() (models.AllOwner, error) {
	var allOwner models.AllOwner
	var owner models.Owner

	db, err := driver.ConnectMongoBO()
	if err != nil {
		return allOwner, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allOwner, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&owner)
		owner.Birthday = owner.Birthday_date.Format("2006-01-02")
		allOwner.Owner = append(allOwner.Owner, owner)
	}
	return allOwner, nil
}

func FindOwner(id primitive.ObjectID) (models.Owner, error) {
	db, err := driver.ConnectMongoBO()
	var owner models.Owner
	if err != nil {
		return owner, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&owner)
	if err != nil {
		return owner, err
	}
	return owner, nil
}

func InsertOwner(ower models.Owner) error {
	collection, err := driver.ConnectMongoBO()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), ower)
	if err != nil {
		return err
	}
	return nil
}
