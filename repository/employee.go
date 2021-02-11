package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllEmployee() (models.AllGetEmployee, error) {
	var allEmp models.AllGetEmployee
	var employee models.GetEmployee
	db, err := driver.ConnectMongoEmp()
	titleState := bson.D{{"$lookup", bson.D{{"from", "titlename"}, {"localField", "title"}, {"foreignField", "_id"}, {"as", "title"}}}}
	positionState := bson.D{{"$lookup", bson.D{{"from", "position"}, {"localField", "position"}, {"foreignField", "_id"}, {"as", "position"}}}}
	companyState := bson.D{{"$lookup", bson.D{{"from", "company"}, {"localField", "company"}, {"foreignField", "_id"}, {"as", "company"}}}}
	genderState := bson.D{{"$lookup", bson.D{{"from", "gender"}, {"localField", "gender"}, {"foreignField", "_id"}, {"as", "gender"}}}}
	cur, err := db.Aggregate(context.TODO(), mongo.Pipeline{titleState, positionState, companyState, genderState})
	if err != nil {
		return allEmp, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&employee)
		employee.TitleName = employee.Title[0]
		employee.Com = employee.Company[0]
		employee.Pst = employee.Position[0]
		employee.Gd = employee.Gender[0]
		allEmp.GetEmployee = append(allEmp.GetEmployee, employee)
	}
	return allEmp, nil
}

func FindEmployee(id string) (models.GetEmployee, error) {
	var employee models.GetEmployee
	db, err := driver.ConnectMongoEmp()
	titleState := bson.D{{"$lookup", bson.D{{"from", "titlename"}, {"localField", "title"}, {"foreignField", "_id"}, {"as", "title"}}}}
	positionState := bson.D{{"$lookup", bson.D{{"from", "position"}, {"localField", "position"}, {"foreignField", "_id"}, {"as", "position"}}}}
	companyState := bson.D{{"$lookup", bson.D{{"from", "company"}, {"localField", "company"}, {"foreignField", "_id"}, {"as", "company"}}}}
	genderState := bson.D{{"$lookup", bson.D{{"from", "gender"}, {"localField", "gender"}, {"foreignField", "_id"}, {"as", "gender"}}}}
	matchStage := bson.D{{"$match", bson.D{{"_id", id}}}}
	cur, err := db.Aggregate(context.TODO(), mongo.Pipeline{titleState, positionState, companyState, genderState, matchStage})
	if err != nil {
		return employee, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&employee)
		employee.TitleName = employee.Title[0]
		employee.Com = employee.Company[0]
		employee.Pst = employee.Position[0]
		employee.Gd = employee.Gender[0]
	}
	return employee, nil
}

func InsertEmployee(emp models.Employee) error {
	collection, err := driver.ConnectMongoEmp()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), emp)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(emp models.Employee, id string) error {
	db, err := driver.ConnectMongoEmp()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", emp}}
	db.CountDocuments(context.TODO(), bson.D{{}})
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

func DeleteEmployee(id string) error {
	db, err := driver.ConnectMongoEmp()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}

func LastEmployee() (models.Employee, error) {
	db, err := driver.ConnectMongoEmp()
	var emp models.Employee
	opts := options.FindOne().SetSort(bson.D{{"_id", -1}})
	err = db.FindOne(context.TODO(), bson.M{}, opts).Decode(&emp)
	if err != nil {
		return emp, err
	}
	return emp, nil
}
