package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/middleware"
	"github.com/supermarketmanager/models"
	"github.com/supermarketmanager/services"
	"net/http"
	"strconv"
)

// SupplierController 供应商控制器

type SupplierController struct {
	supplierService *services.SupplierService
}

// NewSupplierController 创建供应商控制器实例
func NewSupplierController() *SupplierController {
	return &SupplierController{
		supplierService: services.NewSupplierService(),
	}
}

// SetupRoutes 设置路由
func (c *SupplierController) SetupRoutes(r *gin.Engine) {
	supplier := r.Group("/supplier")
	supplier.Use(middleware.AuthMiddleware())
	{
		supplier.GET("/list", c.List)
		supplier.GET("/view/:id", c.View)
		supplier.GET("/toAdd", c.ToAdd)
		supplier.POST("/add", c.Add)
		supplier.GET("/toUpdate/:id", c.ToUpdate)
		supplier.POST("/update", c.Update)
		supplier.GET("/del/:id", c.Delete)
	}
}

// List 供应商列表
func (c *SupplierController) List(ctx *gin.Context) {
	// 获取所有供应商
	suppliers, err := c.supplierService.GetAllSuppliers()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "supplier_list.html", gin.H{
		"suppliers": suppliers,
	})
}

// View 查看供应商详情
func (c *SupplierController) View(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/supplier/list")
		return
	}
	
	// 获取供应商详情
	supplier, err := c.supplierService.GetSupplierByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "supplier_view.html", gin.H{
		"supplier": supplier,
	})
}

// ToAdd 跳转到添加供应商页面
func (c *SupplierController) ToAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "supplier_add.html", nil)
}

// Add 添加供应商
func (c *SupplierController) Add(ctx *gin.Context) {
	// 绑定表单数据
	var supplier models.Supplier
	if err := ctx.ShouldBind(&supplier); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 创建供应商
	if err := c.supplierService.CreateSupplier(&supplier); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到供应商列表
	ctx.Redirect(http.StatusFound, "/supplier/list")
}

// ToUpdate 跳转到修改供应商页面
func (c *SupplierController) ToUpdate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/supplier/list")
		return
	}
	
	// 获取供应商详情
	supplier, err := c.supplierService.GetSupplierByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "supplier_update.html", gin.H{
		"supplier": supplier,
	})
}

// Update 更新供应商
func (c *SupplierController) Update(ctx *gin.Context) {
	// 绑定表单数据
	var supplier models.Supplier
	if err := ctx.ShouldBind(&supplier); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 更新供应商
	if err := c.supplierService.UpdateSupplier(&supplier); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到供应商列表
	ctx.Redirect(http.StatusFound, "/supplier/list")
}

// Delete 删除供应商
func (c *SupplierController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/supplier/list")
		return
	}
	
	// 删除供应商
	if err := c.supplierService.DeleteSupplier(uint(id)); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到供应商列表
	ctx.Redirect(http.StatusFound, "/supplier/list")
}