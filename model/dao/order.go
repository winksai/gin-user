package dao

import (
	"gorm.io/gorm"
)

type OrderDao struct {
	db *gorm.DB
}

func (o *OrderDao) CreateOrder() {
	o.db.Create("demos")
}

func NewOrderDao(db *gorm.DB) *OrderDao {
	return &OrderDao{db: db}
}
