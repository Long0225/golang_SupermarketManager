package controllers

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/middleware"
	"github.com/supermarketmanager/models"
	"github.com/supermarketmanager/services"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// OrderController 订单控制器

type OrderController struct {
	orderService    *services.OrderService
	customerService *services.CustomerService
	productService  *services.ProductService
}

// NewOrderController 创建订单控制器实例
func NewOrderController() *OrderController {
	return &OrderController{
		orderService:    services.NewOrderService(),
		customerService: services.NewCustomerService(),
		productService:  services.NewProductService(),
	}
}

// SetupRoutes 设置路由
func (c *OrderController) SetupRoutes(r *gin.Engine) {
	order := r.Group("/order")
	order.Use(middleware.AuthMiddleware())
	{
		order.GET("/list", c.List)
		order.GET("/view/:id", c.View)
		order.GET("/toAdd", c.ToAdd)
		order.POST("/add", c.Add)
		order.GET("/toUpdate/:id", c.ToUpdate)
		order.POST("/update", c.Update)
		order.GET("/del/:id", c.Delete)
	}
}

// List 订单列表
func (c *OrderController) List(ctx *gin.Context) {
	// 获取所有订单
	orders, err := c.orderService.GetAllOrders()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "order_list.html", gin.H{
		"orders": orders,
	})
}

// View 查看订单详情
func (c *OrderController) View(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/order/list")
		return
	}
	
	// 获取订单详情
	order, err := c.orderService.GetOrderByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "order_view.html", gin.H{
		"order": order,
	})
}

// ToAdd 跳转到添加订单页面
func (c *OrderController) ToAdd(ctx *gin.Context) {
	// 获取所有客户和商品用于下拉选择
	customers, _ := c.customerService.GetAllCustomers()
	products, _ := c.productService.GetAllProducts()
	
	ctx.HTML(http.StatusOK, "order_add.html", gin.H{
		"customers": customers,
		"products":  products,
	})
}

// Add 添加订单
func (c *OrderController) Add(ctx *gin.Context) {
	// 绑定表单数据
	var order models.Order
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 生成订单号
	order.OrderNo = uuid.New().String()[:10]
	// 设置订单时间
	order.OrderTime = time.Now()
	// 设置订单状态
	order.Status = "pending"
	
	// 创建订单
	if err := c.orderService.CreateOrder(&order); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到订单列表
	ctx.Redirect(http.StatusFound, "/order/list")
}

// ToUpdate 跳转到修改订单页面
func (c *OrderController) ToUpdate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/order/list")
		return
	}
	
	// 获取订单详情
	order, err := c.orderService.GetOrderByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 获取所有客户和商品用于下拉选择
	customers, _ := c.customerService.GetAllCustomers()
	products, _ := c.productService.GetAllProducts()
	
	ctx.HTML(http.StatusOK, "order_update.html", gin.H{
		"order":     order,
		"customers": customers,
		"products":  products,
	})
}

// Update 更新订单
func (c *OrderController) Update(ctx *gin.Context) {
	// 绑定表单数据
	var order models.Order
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 更新订单
	if err := c.orderService.UpdateOrder(&order); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到订单列表
	ctx.Redirect(http.StatusFound, "/order/list")
}

// Delete 删除订单
func (c *OrderController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/order/list")
		return
	}
	
	// 删除订单
	if err := c.orderService.DeleteOrder(uint(id)); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到订单列表
	ctx.Redirect(http.StatusFound, "/order/list")
}