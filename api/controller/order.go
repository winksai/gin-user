package controller

import (
	"github.com/6.2/api/request"
	"github.com/6.2/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	service service.IOrderService
}

func NewOrderController(service service.IOrderService) *OrderController {

	return &OrderController{service: service}
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req request.OrderReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "参数获取失败" + err.Error()})
		return
	}
	o.service.Create()
}
