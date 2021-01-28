package database

import (
	"shoppinglist/models"
	"time"
)

type db []models.Item

//Items contains the items of the database
var Items db = db{
	{Name: "Paprika", Desc: "skal ha fin fasong", Quantity: 2, Added: time.Now(), Completed: time.Now().Add(time.Hour * 4)},
	{Name: "is", Desc: "ben & jerries", Quantity: 1, Added: time.Now()},
}

func (db *db) Find(id string) (interface{}, error) {
	for _, v := range Items {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, ErrNotFound
}

func (db *db) FindAll() (interface{}, error) {
	return Items, nil
}
