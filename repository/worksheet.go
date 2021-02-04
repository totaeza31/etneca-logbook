package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllWorksheet() ([]models.WorkSheetRespond, error) {
	var allWorkSheet []models.WorkSheetRespond
	var worksheet models.WorkSheetRespond

	db, err := driver.ConnectMongoWorksheet()
	if err != nil {
		return allWorkSheet, err
	}
	techState := bson.D{{"$lookup", bson.D{{"from", "tech"}, {"localField", "company"}, {"foreignField", "company"}, {"as", "techDetail"}}}}
	boatState := bson.D{{"$lookup", bson.D{{"from", "boat"}, {"localField", "deviceNumber"}, {"foreignField", "deviceNumber"}, {"as", "boatDetail"}}}}
	cur, err := db.Aggregate(context.TODO(), mongo.Pipeline{techState, boatState})
	if err != nil {
		return allWorkSheet, err
	}
	for cur.Next(context.Background()) {

		err = cur.Decode(&worksheet)
		worksheet.BoatName = worksheet.BoatDevice[0].Name
		worksheet.Telephone = worksheet.TechDetail[0].Telephone
		address := strings.SplitAfter(worksheet.TechDetail[0].Address, "จ.")
		worksheet.StartDate = worksheet.Time.Format("2006-01-02")
		worksheet.Address = address[1]
		allWorkSheet = append(allWorkSheet, worksheet)
	}
	return allWorkSheet, nil
}

func FindWorksheet(id primitive.ObjectID) (models.WorkSheetRespond, error) {
	var worksheet models.WorkSheetRespond
	db, err := driver.ConnectMongoWorksheet()
	if err != nil {
		return worksheet, err
	}
	techState := bson.D{{"$lookup", bson.D{{"from", "tech"}, {"localField", "company"}, {"foreignField", "company"}, {"as", "techDetail"}}}}
	boatState := bson.D{{"$lookup", bson.D{{"from", "boat"}, {"localField", "deviceNumber"}, {"foreignField", "deviceNumber"}, {"as", "boatDetail"}}}}
	matchStage := bson.D{{"$match", bson.D{{"_id", id}}}}
	cur, err := db.Aggregate(context.TODO(), mongo.Pipeline{techState, boatState, matchStage})
	if err != nil {
		return worksheet, err
	}
	for cur.Next(context.Background()) {

		err = cur.Decode(&worksheet)
		worksheet.BoatName = worksheet.BoatDevice[0].Name
		worksheet.Telephone = worksheet.TechDetail[0].Telephone
		address := strings.SplitAfter(worksheet.TechDetail[0].Address, "จ.")
		worksheet.StartDate = worksheet.Time.Format("2006-01-02")
		worksheet.Address = address[1]
	}
	return worksheet, nil
}

func InsertWorksheet(workSheet models.WorkSheet) error {
	collection, err := driver.ConnectMongoWorksheet()

	workSheet.Time = time.Now()
	_, err = collection.InsertOne(context.Background(), workSheet)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWorksheet(workSheet models.WorkSheet, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoWorksheet()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", workSheet}}
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
