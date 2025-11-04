package repository

import (
	"fmt"
	"os"
	"time"

	"ordersystem/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Store holds the DB handle.
type Store struct {
	DB *gorm.DB
}

// NewStore connects, migrates, and seeds the database.
func NewStore() (*Store, error) {
	host := getenv("DB_HOST", "sbd3-postgres")
	dbname := getenv("POSTGRES_DB", "order")
	user := getenv("POSTGRES_USER", "docker")
	pass := getenv("POSTGRES_PASSWORD", "docker")
	port := getenv("POSTGRES_TCP_PORT", "5432")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, pass, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	// Auto-migrate schema
	if err := db.AutoMigrate(&model.Drink{}, &model.Order{}); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	// Seed drinks if empty
	var n int64
	if err := db.Model(&model.Drink{}).Count(&n).Error; err != nil {
		return nil, err
	}
	if n == 0 {
		seedDrinks := []model.Drink{
			{Name: "Latte", PriceCents: 390},
			{Name: "Espresso", PriceCents: 250},
			{Name: "Tea", PriceCents: 220},
		}
		if err := db.Create(&seedDrinks).Error; err != nil {
			return nil, fmt.Errorf("seed drinks: %w", err)
		}
	}

	// Seed a few orders if empty
	if err := db.Model(&model.Order{}).Count(&n).Error; err != nil {
		return nil, err
	}
	if n == 0 {
		var latte model.Drink
		if err := db.Where("name = ?", "Latte").First(&latte).Error; err == nil {
			o := []model.Order{
				{DrinkID: latte.ID, Qty: 2},
				{DrinkID: latte.ID, Qty: 1},
			}
			if err := db.Create(&o).Error; err != nil {
				return nil, fmt.Errorf("seed orders: %w", err)
			}
		}
	}

	// Tune connections
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return &Store{DB: db}, nil
}

// Drinks returns all drinks.
func (s *Store) Drinks() ([]model.Drink, error) {
	var out []model.Drink
	return out, s.DB.Order("id").Find(&out).Error
}

// CreateOrder inserts an order.
func (s *Store) CreateOrder(drinkID uint, qty int) (model.Order, error) {
	o := model.Order{DrinkID: drinkID, Qty: qty}
	return o, s.DB.Create(&o).Error
}

// TotalledOrders returns joined order lines with totals.
func (s *Store) TotalledOrders() ([]model.TotalledOrder, error) {
	var out []model.TotalledOrder
	err := s.DB.
		Table("orders o").
		Select("o.id as order_id, d.name as drink_name, o.qty, (o.qty * d.price_cents) as line_total_cents, o.created_at").
		Joins("join drinks d on d.id = o.drink_id").
		Order("o.id").
		Scan(&out).Error
	return out, err
}

// getenv returns env var or default.
func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
