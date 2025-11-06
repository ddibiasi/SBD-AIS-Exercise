package repository

import (
<<<<<<< HEAD
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
=======
	"errors"
	"fmt"
	"log/slog"
	"ordersystem/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseHandler struct {
	dbConn *gorm.DB
}

func NewDatabaseHandler() (*DatabaseHandler, error) {
	slog.Info("Connecting to database")
	// connect to db
	dsn, err := getDsn()
	if err != nil {
		return nil, err
	}
	dbConn, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// create tables and migrate
	err = dbConn.AutoMigrate(&model.Drink{}, &model.Order{})
	if err != nil {
		return nil, err
	}
	// add test data to db
	err = prepopulate(dbConn)
	if err != nil {
		return nil, err
	}
	return &DatabaseHandler{dbConn: dbConn}, nil
}

func getDsn() (string, error) {
	dbUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_USER' is not set")
	}
	dbPw, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_PASSWORD' is not set")
	}
	dbName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_DB' is not set")
	}
	dbPort, ok := os.LookupEnv("POSTGRES_TCP_PORT")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_TCP_PORT' is not set")
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return "", errors.New("environment variable 'DB_HOST' is not set")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPw, dbName, dbPort)
	return dsn, nil
}

func prepopulate(dbConn *gorm.DB) error {
	// check if prepopulate has already run once
	var exists bool
	err := dbConn.Model(&model.Drink{}).
		Select("count(*) > 0").
		Find(&exists).
		Error
	if err != nil {
		return err
	}
	if exists {
		// don't prepopulate if has already run
		return nil
	}
	// create drink menu
	// todo create drinks
	// todo create orders
	// GORM documentation can be found here: https://gorm.io/docs/index.html

	return nil
}

func (db *DatabaseHandler) GetDrinks() (drinks []model.Drink, err error) {
	err = db.dbConn.Find(&drinks).Error
	if err != nil {
		return nil, err
	}
	return drinks, nil
}

func (db *DatabaseHandler) GetOrders() (orders []model.Order, err error) {
	err = db.dbConn.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

const totalledStmt = `SELECT drink_id, SUM(amount) AS total_amount_ordered FROM orders WHERE deleted_at IS NULL GROUP BY drink_id ORDER BY drink_id;`

func (db *DatabaseHandler) GetTotalledOrders() (totals []model.DrinkOrderTotal, err error) {
	err = db.dbConn.Raw(totalledStmt).Scan(&totals).Error
	if err != nil {
		return nil, err
	}
	return totals, nil
}

func (db *DatabaseHandler) AddOrder(order *model.Order) error {
	err := db.dbConn.Create(order).Error
	if err != nil {
		return err
	}
	return nil
>>>>>>> ed41b50cc2fd92dbd0df12eafd134d95e2bbbd93
}
