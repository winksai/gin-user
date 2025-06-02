 # gin 单体框架搭建
    本项目是基于 **Gin + GORM** 构建的后端服务框架，采用经典的 **MVC 分层架构**，并结合接口抽象、模块化设计，为后续微服务化演进提供良好基础。
    
    适合中小型项目快速搭建，也可作为大型项目微服务架构的原型。

目录结构
├── api # API 层
│ ├── controller # 控制器：处理 HTTP 请求
│ │ └── order.go
│ └── request # 请求参数封装结构体
│ └── order_req.go
│
├── global # 全局变量，如数据库连接等
│ └── global.go
│
├── initialize # 项目初始化逻辑（数据库、服务等）
│ └── init.go
│
├── model # 模型层（可含 DAO 与数据库模型）
│ └── dao
│ └── order.go
│
├── router # 路由注册模块
│ ├── router.go # 主路由初始化
│ └── order_router.go # 订单模块路由
│
├── service # 服务层：业务逻辑封装
│ └── order_service.go
│
├── go.mod # Go 依赖管理
├── main.go # 应用程序入口