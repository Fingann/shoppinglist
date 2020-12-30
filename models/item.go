package models

import "time"

// Item struct representing a table in the database for keeping track of objects
type Item struct {
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Quantity  int       `json:"quantity"`
	Added     time.Time `json:"added"`
	Completed time.Time `json:"completed"`
}
