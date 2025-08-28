package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/middleware"
	"github.com/supermarketmanager/models"
	"github.com/supermarketmanager/services"
	"net/http"
	"strconv"
)

// CustomerController 客户控制器

type CustomerController struct {
	customerService *services.CustomerService
}

// NewCustomerController 创建客户控制器实例
func NewCustomerController() *CustomerController {
	return &CustomerController{
		customerService: services.NewCustomerService(),
	}
}

// SetupRoutes 设置路由
func (c *CustomerController) SetupRoutes(r *gin.Engine) {
	customer := r.Group("/customer")
	customer.Use(middleware.AuthMiddleware())
	{
		customer.GET("/list", c.List)
		customer.GET("/view/:id", c.View)
		customer.GET("/toAdd", c.ToAdd)
		customer.POST("/add", c.Add)
		customer.GET("/toUpdate/:id", c.ToUpdate)
		customer.POST("/update", c.Update)
		customer.GET("/del/:id", c.Delete)
	}
}

// List 客户列表
func (c *CustomerController) List(ctx *gin.Context) {
	// 获取所有客户
	customers, err := c.customerService.GetAllCustomers()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "customer_list.html", gin.H{
		"customers": customers,
	})
}

// View 查看客户详情
func (c *CustomerController) View(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/customer/list")
		return
	}
	
	// 获取客户详情
	customer, err := c.customerService.GetCustomerByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "customer_view.html", gin.H{
		"customer": customer,
	})
}

// ToAdd 跳转到添加客户页面
func (c *CustomerController) ToAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "customer_add.html", nil)
}

// Add 添加客户
func (c *CustomerController) Add(ctx *gin.Context) {
	// 绑定表单数据
	var customer models.Customer
	if err := ctx.ShouldBind(&customer); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 创建客户
	if err := c.customerService.CreateCustomer(&customer); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到客户列表
	ctx.Redirect(http.StatusFound, "/customer/list")
}

// ToUpdate 跳转到修改客户页面
func (c *CustomerController) ToUpdate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/customer/list")
		return
	}
	
	// 获取客户详情
	customer, err := c.customerService.GetCustomerByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "customer_update.html", gin.H{
		"customer": customer,
	})
}

// Update 更新客户
func (c *CustomerController) Update(ctx *gin.Context) {
	// 绑定表单数据
	var customer models.Customer
	if err := ctx.ShouldBind(&customer); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 更新客户
	if err := c.customerService.UpdateCustomer(&customer); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到客户列表
	ctx.Redirect(http.StatusFound, "/customer/list")
}

// Delete 删除客户
func (c *CustomerController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/customer/list")
		return
	}
	
	// 删除客户
	if err := c.customerService.DeleteCustomer(uint(id)); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到客户列表
	ctx.Redirect(http.StatusFound, "/customer/list")
}