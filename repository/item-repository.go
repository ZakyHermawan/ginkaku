package repository

import (
	"fmt"
	"github.com/ZakyHermawan/ginkaku/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ItemRepository interface {
	CreateOrder(order *entity.Order) error
	UpdateOrder(order *entity.Order) error
	DeleteOrder(order *entity.Order) error
	GetAllOrder() ([]entity.Order, error)
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewItemRepository() ItemRepository {
	fmt.Println("new repo")
	db, err := gorm.Open("sqlite3", "orders_by.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Order{}, &entity.Order{})

	db.AutoMigrate(&entity.Item{}, &entity.Item{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) CreateOrder(order *entity.Order) error {
	return db.connection.Model(order).Create(order).Error
}

func (db *database) UpdateOrder(order *entity.Order) error {
	return db.connection.Model(order).Update(order).Error
}

func (db *database) DeleteOrder(order *entity.Order) error {
	return db.connection.Model(order).Delete(order).Error
}

func (db *database) GetAllOrder() ([]entity.Order, error) {
	var orders []entity.Order
	err := db.connection.Set("gorm:auto_preload", true).Find(&orders).Error
	return orders, err
}
