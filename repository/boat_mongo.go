package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllBoat() (models.AllBoats, error) {
	var allboat models.AllBoats
	var boat models.Boat

	db, err := driver.ConnectMongoBoat()
	if err != nil {
		return allboat, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allboat, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boat)
		boat.Anniversary = boat.Anniversary_date.Format("2006-01-02")
		boat.WarrantyExp = boat.WarrantyExp_date.Format("2006-01-02")
		boat.ReportDate = boat.ReportDate_date.Format("2006-01-02")
		allboat.Boat = append(allboat.Boat, boat)
	}
	return allboat, nil
}

func FindBoat(id primitive.ObjectID) (models.Boat, error) {
	var boat models.Boat
	db, err := driver.ConnectMongoBoat()
	if err != nil {
		return boat, err
	}
	
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boat)
	if err != nil {
		return boat, err
	}
	boat.Anniversary = boat.Anniversary_date.Format("2006-01-02")
	boat.WarrantyExp = boat.WarrantyExp_date.Format("2006-01-02")
	boat.ReportDate = boat.ReportDate_date.Format("2006-01-02")
	return boat, nil
}

func InsertBoat(boat models.Boat) error {
	collection, err := driver.ConnectMongoBoat()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), boat)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBoat(boat models.Boat, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoat()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", boat}}
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

func DeleteBoat(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoBoat()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}

func FindBoatByName(text string) (models.AllBoats, error) {
	var allboat models.AllBoats
	var boat models.Boat

	db, err := driver.ConnectMongoBoat()
	if err != nil {
		return allboat, err

	}
	text = "(.*)" + text + "(.*)"
	searchStage := bson.D{{"$search", bson.D{{"regex", bson.D{{"query", text}, {"path",
		bson.A{"boatName.th", "boatName.en", "boatName.en", "encBoxNumber", "boatReg", "shipTrackingReport", "boatType", "deviceNumber"}}, {
		"allowAnalyzedField", true}}}}}}

	cur, err := db.Aggregate(context.TODO(), mongo.Pipeline{searchStage})

	if err != nil {
		return allboat, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&boat)
		boat.Anniversary = boat.Anniversary_date.Format("2006-01-02")
		boat.WarrantyExp = boat.WarrantyExp_date.Format("2006-01-02")
		boat.ReportDate = boat.ReportDate_date.Format("2006-01-02")
		allboat.Boat = append(allboat.Boat, boat)
	}

	return allboat, nil
}

func FindBoatDeviceNumber(deviceNumber string) (models.Boat, error) {
	var boat models.Boat
	db, err := driver.ConnectMongoBoat()
	if err != nil {
		return boat, err
	}
	err = db.FindOne(context.TODO(), bson.M{"deviceNumber": deviceNumber}).Decode(&boat)
	if err != nil {
		return boat, err
	}
	return boat, nil
}
