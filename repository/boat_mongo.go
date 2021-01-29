package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/helpers"
	"etneca-logbook/models"
	"fmt"

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

func FindBoat(id primitive.ObjectID) (models.Boats, error) {
	var boat models.Boats
	db, err := driver.ConnectMongoBoat()
	boatBeamStage := bson.D{{"$lookup", bson.D{{"from", "boatBeamStatus"}, {"localField", "boatBeamStatus"}, {"foreignField", "_id"}, {"as", "boatBeamStatus"}}}}
	vmsGenStage := bson.D{{"$lookup", bson.D{{"from", "boatVmsGen"}, {"localField", "vmsGen"}, {"foreignField", "_id"}, {"as", "vmsGen"}}}}
	gatewayStage := bson.D{{"$lookup", bson.D{{"from", "boatGateway"}, {"localField", "gateway"}, {"foreignField", "_id"}, {"as", "gateway"}}}}
	boatTypeStage := bson.D{{"$lookup", bson.D{{"from", "boatType"}, {"localField", "boatType"}, {"foreignField", "_id"}, {"as", "boatType"}}}}
	deviceStatusStage := bson.D{{"$lookup", bson.D{{"from", "boatDeviceStatus"}, {"localField", "deviceStatus"}, {"foreignField", "_id"}, {"as", "deviceStatus"}}}}
	financialStatusStage := bson.D{{"$lookup", bson.D{{"from", "boatFinancialStatus"}, {"localField", "financialStatus"}, {"foreignField", "_id"}, {"as", "financialStatus"}}}}
	LoadedStructCursor, err := db.Aggregate(context.TODO(), mongo.Pipeline{boatBeamStage, vmsGenStage, gatewayStage, boatTypeStage, deviceStatusStage, financialStatusStage})
	fmt.Println(LoadedStructCursor)
	var showsLoadedStruct []bson.M
	if err = LoadedStructCursor.All(context.TODO(), &showsLoadedStruct); err != nil {
		return boat, err
	}
	fmt.Println(showsLoadedStruct)
	bsonBytes, _ := bson.Marshal(showsLoadedStruct[0])
	boats := <-helpers.UnmarshalData(bsonBytes, boat)
	fmt.Println(boats)
	return boats, nil
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
