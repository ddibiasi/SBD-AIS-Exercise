package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serve as the order history
	orders []model.Order
}

// NewDatabaseHandler initializes an in-memory database
func NewDatabaseHandler() *DatabaseHandler {
	// Initialize the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Cola", Price: 2, Description: "Classic soda drink", CreatedAt: time.Now()},
		{ID: 2, Name: "Lemonade", Price: 3, Description: "Fresh homemade lemonade", CreatedAt: time.Now()},
		{ID: 3, Name: "Water", Price: 1, Description: "Still mineral water", CreatedAt: time.Now()},
		{ID: 4, Name: "Orange Juice", Price: 3, Description: "100% pure orange juice", CreatedAt: time.Now()},
		{ID: 5, Name: "Apple Juice", Price: 3, Description: "Natural apple juice", CreatedAt: time.Now()},
		{ID: 6, Name: "Iced Tea", Price: 2, Description: "Refreshing lemon iced tea", CreatedAt: time.Now()},
		{ID: 7, Name: "Ginger Ale", Price: 2, Description: "Spicy and sweet ginger soda", CreatedAt: time.Now()},
		{ID: 8, Name: "Energy Drink", Price: 4, Description: "High caffeine energy booster", CreatedAt: time.Now()},
		}
	// Initialize the orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, Amount: 3, CreatedAt: time.Now().Add(-48 * time.Hour)},
		{DrinkID: 2, Amount: 1, CreatedAt: time.Now().Add(-46 * time.Hour)},
		{DrinkID: 3, Amount: 5, CreatedAt: time.Now().Add(-44 * time.Hour)},
		{DrinkID: 4, Amount: 2, CreatedAt: time.Now().Add(-42 * time.Hour)},
		{DrinkID: 5, Amount: 1, CreatedAt: time.Now().Add(-40 * time.Hour)},
		{DrinkID: 6, Amount: 4, CreatedAt: time.Now().Add(-38 * time.Hour)},
		{DrinkID: 7, Amount: 2, CreatedAt: time.Now().Add(-36 * time.Hour)},
		{DrinkID: 8, Amount: 3, CreatedAt: time.Now().Add(-34 * time.Hour)},
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

// GetTotalledOrders calculates total ordered amounts per DrinkID
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)

	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += uint64(order.Amount)
	}

	return totalledOrders
}

// AddOrder adds a new order to the database
func (db *DatabaseHandler) AddOrder(order *model.Order) {
	order.CreatedAt = time.Now()
	db.orders = append(db.orders, *order)
}
