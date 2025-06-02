package service

import "github.com/6.2/model/dao"

type IOrderService interface {
	Create()
}

type OrderService struct {
	dao *dao.OrderDao
}

func (o *OrderService) Create() {
	o.dao.CreateOrder()
}

func NewOrderService(dao *dao.OrderDao) *OrderService {
	return &OrderService{
		dao: dao,
	}
}
