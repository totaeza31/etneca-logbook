package repository

import (
	"context"
	"etneca-logbook/driver"
	"etneca-logbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllGoods() (models.AllGoods, error) {
	var allGoods models.AllGoods
	var goods models.Goods

	db, err := driver.ConnectMongoGoods()
	if err != nil {
		return allGoods, err
	}
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return allGoods, err
	}
	for cur.Next(context.Background()) {
		err = cur.Decode(&goods)
		allGoods.Goods = append(allGoods.Goods, goods)
	}
	return allGoods, nil
}

func FindGoods(id primitive.ObjectID) (models.Goods, error) {
	var goods models.Goods
	db, err := driver.ConnectMongoGoods()
	if err != nil {
		return goods, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&goods)
	if err != nil {
		return goods, err
	}
	return goods, nil
}

func InsertGoods(goods models.Goods) error {
	collection, err := driver.ConnectMongoGoods()
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), goods)
	if err != nil {
		return err
	}
	return nil
}

func UpdateGoods(goods models.Goods, id primitive.ObjectID) error {
	db, err := driver.ConnectMongoGoods()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", goods}}
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

func DeleteGoods(id primitive.ObjectID) error {
	db, err := driver.ConnectMongoGoods()
	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return err
}
