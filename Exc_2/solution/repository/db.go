package repository

import (
	"ordersystem/model"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{Description: "Cocacola zero sugar", Id: 1, Name: "Cocacola", Price: 2},
		{Description: "Orange Fanta", Id: 2, Name: "Fanta", Price: 1.5},
		{Description: "Coffe with milk", Id: 3, Name: "LateMacciato", Price: 1},
	}

	// Init orders slice with some test data
	orders := []model.Order{
		{Amount: 2, CreatedAt: "5 pm", DrinkID: 1},
		{Amount: 1, CreatedAt: "6:30 pm", DrinkID: 2},
		{Amount: 1, CreatedAt: "2 am", DrinkID: 3},
		{Amount: 3, CreatedAt: "12 am", DrinkID: 1},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	totalledOrders := make(map[uint64]uint64)
	for _, o := range db.orders {
		totalledOrders[o.DrinkID] += uint64(o.Amount)
	}

	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}
