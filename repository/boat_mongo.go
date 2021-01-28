package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	db, err := driver.ConnectMongoBoat()
	var boat models.Boat
	if err != nil {
		return boat, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&boat)
	if err != nil {
		return boat, err
	}
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

func UpdateBoat(human models.Human, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoHuman()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", human}}
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
