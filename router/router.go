package router

import (
	"github.com/6.2/initialize"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	OrderRouter
}

var AllRouter = new(RouterGroup)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	initialize.Mysql()
	group := engine.Group("api/v1")
	AllRouter.OrderRouter.InitOrderRouter(group)
	return engine
}
