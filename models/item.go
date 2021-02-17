package models

import "time"

// Item struct representing a table in the database for keeping track of objects
type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Quantity  int       `json:"quantity"`
	Added     time.Time `json:"added"`
	Completed time.Time `json:"completed"`
}

//ToDTO converts the item to a Data Transfer Object
func (i *Item) ToDTO() *ItemDTO {
	return &ItemDTO{
		Name:      i.Name,
		Desc:      i.Desc,
		Quantity:  i.Quantity,
		Added:     i.Added,
		Completed: i.Completed,
	}
}

//ItemDTO represents the Data Transfer Object version of Item struct
type ItemDTO struct {
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Quantity  int       `json:"quantity"`
	Added     time.Time `json:"added"`
	Completed time.Time `json:"completed"`
}

type AddItemRequest struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Quantity int    `json:"quantity"`
}

type GetItemRequest struct {
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Quantity  int       `json:"quantity"`
	Added     time.Time `json:"added"`
	Completed time.Time `json:"completed"`
}
