package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllWorksheet() (models.AllWorkSheet, error) {
	var allWorkSheet models.AllWorkSheet
	var worksheet models.WorkSheet

	db, err := driver.ConnectMongoWorksheet()
	if err != nil {
		return allWorkSheet, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allWorkSheet, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&worksheet)
		allWorkSheet.WorkSheet = append(allWorkSheet.WorkSheet, worksheet)
	}
	return allWorkSheet, nil
}

func FindWorksheet(id primitive.ObjectID) (models.Tech, error) {
	db, err := driver.ConnectMongoWorksheet()
	var tech models.Tech
	if err != nil {
		return tech, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&tech)
	if err != nil {
		return tech, err
	}
	return tech, nil
}

func InsertWorksheet(workSheet models.WorkSheet) error {
	collection, err := driver.ConnectMongoWorksheet()
	tech, err := FindTechName(workSheet.Company)
	if err != nil {
		return err
	}
	fmt.Println(tech.ID)
	_, err = collection.InsertOne(context.Background(), workSheet)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWorksheet(tech models.Tech, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoWorksheet()
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
		return nil
	}
}

func DeleteWorksheet(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoWorksheet()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
