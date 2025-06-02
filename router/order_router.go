package router

import (
	"github.com/6.2/api/controller"
	"github.com/6.2/global"
	"github.com/6.2/model/dao"
	"github.com/6.2/service"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (o *OrderRouter) InitOrderRouter(r *gin.RouterGroup) {

	order := r.Group("order")
	order.Use() //jwt
	orderController := controller.NewOrderController(service.NewOrderService(dao.NewOrderDao(global.DB)))
	{
		order.POST("/create", orderController.CreateOrder)
	}
}
