package repository

import (
	"etneca-logbook/driver"
	"time"
)

func SetToken(id string, token string, ex time.Time, sub time.Time) error {
	client, err := driver.ConnectRedis()
	if err != nil {
		return err
	}
	err = client.Set(id, token, ex.Sub(sub)).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteToken(id string) error {
	client, err := driver.ConnectRedis()
	if err != nil {
		return err
	}
	client.Del(id).Result()
	return nil
}

func GetToken(id string) (string, error) {
	client, err := driver.ConnectRedis()
	val, err := client.Get(id).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
