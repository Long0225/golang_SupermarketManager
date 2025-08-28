package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/controllers"
	"github.com/supermarketmanager/database"
	"log"
	"net/http"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println("Database connected successfully")

	// 创建 Gin 引擎实例
	router := gin.Default()

	// 设置模板目录
	router.LoadHTMLGlob("templates/*")

	// 设置静态文件目录
	router.Static("/static", "./static")

	// 初始化并设置控制器路由
	systemController := controllers.NewSystemController()
	systemController.SetupRoutes(router)

	productController := controllers.NewProductController()
	productController.SetupRoutes(router)

	categoryController := controllers.NewCategoryController()
	categoryController.SetupRoutes(router)

	supplierController := controllers.NewSupplierController()
	supplierController.SetupRoutes(router)

	customerController := controllers.NewCustomerController()
	customerController.SetupRoutes(router)

	orderController := controllers.NewOrderController()
	orderController.SetupRoutes(router)

	systemLogController := controllers.NewSystemLogController()
	systemLogController.SetupRoutes(router)

	// 添加根路径重定向到登录页面
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})

	// 启动服务器
	fmt.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
