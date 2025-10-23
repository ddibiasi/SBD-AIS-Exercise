package model

import "time"

type Drink struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
