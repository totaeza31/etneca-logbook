package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindAllEmployee() (models.AllEmp, error) {
	var allEmp models.AllEmp
	var emp models.Employee

	db, err := driver.ConnectMongoEmp()
	if err != nil {
		return allEmp, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allEmp, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&emp)

		allEmp.Employee = append(allEmp.Employee, emp)
	}
	return allEmp, nil
}

func FindEmployee(id string) (models.Employee, error) {
	var emp models.Employee
	db, err := driver.ConnectMongoEmp()
	if err != nil {
		return emp, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&emp)
	if err != nil {
		return emp, err
	}
	return emp, nil
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

func CountEmployee() (int64, error) {
	db, err := driver.ConnectMongoEmp()
	doc, err := db.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return doc, err
	}
	return doc, nil
}
