package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllReport() (models.AllReport, error) {
	db, err := driver.ConnectMongoReport()
	var reports models.AllReport
	var report models.Report
	if err != nil {
		return reports, err
	}

	lookupState := bson.D{{"$lookup", bson.D{{"from", "boat"}, {"localField", "boatId"}, {"foreignField", "_id"}, {"as", "boatName"}}}}
	// projectState :=
	cur, err := db.Aggregate(context.TODO(), mongo.Pipeline{lookupState})
	if err != nil {
		return reports, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&report)
		report.Start = report.StartDate.Format("2006-01-02")
		report.Name = report.BoatName[0].BoatName
		fmt.Println(report.BoatName)
		report.Lastest = report.LastestDate.Format("2006-01-02")
		reports.Report = append(reports.Report, report)
	}

	return reports, nil
}

func FindReport(id primitive.ObjectID) (models.Report, error) {
	db, err := driver.ConnectMongoReport()
	var report models.Report
	if err != nil {
		return report, err
	}

	lookupState := bson.D{{"$lookup", bson.D{{"from", "boat"}, {"localField", "boatId"}, {"foreignField", "_id"}, {"as", "boatName"}}}}
	matchStage := bson.D{{"$match", bson.D{{"_id", id}}}}

	cur, err := db.Aggregate(context.TODO(), mongo.Pipeline{lookupState, matchStage})
	if err != nil {
		return report, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&report)
		report.Start = report.StartDate.Format("2006-01-02")
		report.Name = report.BoatName[0].BoatName
		report.Lastest = report.LastestDate.Format("2006-01-02")

	}

	return report, nil
}

func InsertReport(report models.Report) error {
	collection, err := driver.ConnectMongoReport()
	if err != nil {
		return err
	}
	report.StartDate = time.Now()
	report.LastestDate = time.Now()
	_, err = collection.InsertOne(context.Background(), report)
	if err != nil {
		return err
	}
	return nil
}

func UpdateReport(report models.Report, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoReport()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", report}}
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

func DeleteReport(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoReport()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
