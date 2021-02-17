package database

import (
	"fmt"
	"shoppinglist/models"
	"time"
)

//ItemListDB represents a simple in-memory database (a list)
type ItemListDB struct {
	items []models.Item
}

//NewItemListDB returns an empty ItemListDB
func NewItemListDB() *ItemListDB {
	return &database
}

//Initialize the database abd a identifier for id
var (
	database ItemListDB = ItemListDB{
		items: []models.Item{
			{ID: "0", Name: "Paprika", Desc: "skal ha fin fasong", Quantity: 2, Added: time.Now(), Completed: time.Now().Add(time.Hour * 4)},
			{ID: "1", Name: "is", Desc: "ben & jerries", Quantity: 1, Added: time.Now()},
		},
	}
	identifier = 2
)

//Find a item using the id as criteria
func (db ItemListDB) Find(id string) (interface{}, error) {

	for _, v := range db.items {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, ErrNotFound
}

//FindAll retrieves all items from the database
func (db ItemListDB) FindAll() (interface{}, error) {
	return db.items, nil
}

//Filter retrieves items based on a predicate
func (db ItemListDB) Filter(predicate func(interface{}) bool) (interface{}, error) {
	tmpItems := []models.Item{}
	for _, v := range db.items {
		if predicate(v) {
			tmpItems = append(tmpItems, v)
		}
	}
	return tmpItems, nil
}

//Update an entity
func (db ItemListDB) Update(entity interface{}) (interface{}, error) {

	if item, ok := entity.(models.Item); ok {
		db.items = append(db.items, item)
	}
	return nil, ErrOperationFailed

}

//Create an entity
func (db ItemListDB) Create(entity interface{}) (interface{}, error) {
	if item, ok := entity.(models.Item); ok {
		item.ID = fmt.Sprint(identifier)
		db.items = append(db.items, item)
		identifier++
		return item, nil
	}
	return nil, ErrOperationFailed
}

//Delete an item from the database
func (db ItemListDB) Delete(id string) (interface{}, error) {
	length := len(db.items)
	for i := 0; i < length; i++ {
		if db.items[i].ID == id {
			lastItem, currentItem := db.items[i], db.items[length-1]
			db.items[len(db.items)-1] = currentItem
			db.items[i] = lastItem
			db.items = db.items[:length-1]
			return currentItem, nil
		}
	}
	return nil, ErrNotFound
}
