package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllCompany() (models.AllCompany, error) {
	var allcomp models.AllCompany
	var comp models.Company

	db, err := driver.ConnectMongoCompany()
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
	return allcomp, nil
}

func FindCompany(id primitive.ObjectID) (models.Company, error) {
	var comp models.Company
	db, err := driver.ConnectMongoCompany()
	if err != nil {
		return comp, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&comp)
	if err != nil {
		return comp, err
	}
	return comp, nil
}

func InsertCompany(comp models.Company) error {
	collection, err := driver.ConnectMongoCompany()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), comp)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCompany(comp models.Company, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoCompany()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", comp}}
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

func DeleteCompany(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoCompany()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
