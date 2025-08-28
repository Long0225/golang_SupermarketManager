package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/middleware"
	"github.com/supermarketmanager/models"
	"github.com/supermarketmanager/services"
	"net/http"
	"strconv"
)

// ProductController 商品控制器

type ProductController struct {
	productService  *services.ProductService
	categoryService *services.CategoryService
	supplierService *services.SupplierService
}

// NewProductController 创建商品控制器实例
func NewProductController() *ProductController {
	return &ProductController{
		productService:  services.NewProductService(),
		categoryService: services.NewCategoryService(),
		supplierService: services.NewSupplierService(),
	}
}

// SetupRoutes 设置路由
func (c *ProductController) SetupRoutes(r *gin.Engine) {
	product := r.Group("/product")
	product.Use(middleware.AuthMiddleware())
	{
		product.GET("/list", c.List)
		product.GET("/view/:id", c.View)
		product.GET("/toAdd", c.ToAdd)
		product.POST("/add", c.Add)
		product.GET("/toUpdate/:id", c.ToUpdate)
		product.POST("/update", c.Update)
		product.GET("/del/:id", c.Delete)
		product.GET("/query", c.Query)
	}
}

// List 商品列表
func (c *ProductController) List(ctx *gin.Context) {
	// 获取所有商品
	products, err := c.productService.GetAllProducts()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 获取所有分类和供应商用于下拉选择
	categories, _ := c.categoryService.GetAllCategories()
	suppliers, _ := c.supplierService.GetAllSuppliers()
	
	ctx.HTML(http.StatusOK, "product_list.html", gin.H{
		"products":   products,
		"categories": categories,
		"suppliers":  suppliers,
	})
}

// View 查看商品详情
func (c *ProductController) View(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/product/list")
		return
	}
	
	// 获取商品详情
	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "product_view.html", gin.H{
		"product": product,
	})
}

// ToAdd 跳转到添加商品页面
func (c *ProductController) ToAdd(ctx *gin.Context) {
	// 获取所有分类和供应商用于下拉选择
	categories, _ := c.categoryService.GetAllCategories()
	suppliers, _ := c.supplierService.GetAllSuppliers()
	
	ctx.HTML(http.StatusOK, "product_add.html", gin.H{
		"categories": categories,
		"suppliers":  suppliers,
	})
}

// Add 添加商品
func (c *ProductController) Add(ctx *gin.Context) {
	// 绑定表单数据
	var product models.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 创建商品
	if err := c.productService.CreateProduct(&product); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到商品列表
	ctx.Redirect(http.StatusFound, "/product/list")
}

// ToUpdate 跳转到修改商品页面
func (c *ProductController) ToUpdate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/product/list")
		return
	}
	
	// 获取商品详情
	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 获取所有分类和供应商用于下拉选择
	categories, _ := c.categoryService.GetAllCategories()
	suppliers, _ := c.supplierService.GetAllSuppliers()
	
	ctx.HTML(http.StatusOK, "product_update.html", gin.H{
		"product":    product,
		"categories": categories,
		"suppliers":  suppliers,
	})
}

// Update 更新商品
func (c *ProductController) Update(ctx *gin.Context) {
	// 绑定表单数据
	var product models.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 更新商品
	if err := c.productService.UpdateProduct(&product); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到商品列表
	ctx.Redirect(http.StatusFound, "/product/list")
}

// Delete 删除商品
func (c *ProductController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/product/list")
		return
	}
	
	// 删除商品
	if err := c.productService.DeleteProduct(uint(id)); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到商品列表
	ctx.Redirect(http.StatusFound, "/product/list")
}

// Query 查询商品
func (c *ProductController) Query(ctx *gin.Context) {
	// 构建查询条件
	query := make(map[string]interface{})
	if name := ctx.Query("name"); name != "" {
		query["name"] = name
	}
	if categoryID := ctx.Query("category_id"); categoryID != "" {
		if id, err := strconv.ParseUint(categoryID, 10, 32); err == nil {
			query["category_id"] = uint(id)
		}
	}
	if supplierID := ctx.Query("supplier_id"); supplierID != "" {
		if id, err := strconv.ParseUint(supplierID, 10, 32); err == nil {
			query["supplier_id"] = uint(id)
		}
	}
	if minPrice := ctx.Query("min_price"); minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
			query["min_price"] = price
		}
	}
	if maxPrice := ctx.Query("max_price"); maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			query["max_price"] = price
		}
	}
	
	// 查询商品
	products, err := c.productService.QueryProducts(query)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 获取所有分类和供应商用于下拉选择
	categories, _ := c.categoryService.GetAllCategories()
	suppliers, _ := c.supplierService.GetAllSuppliers()
	
	ctx.HTML(http.StatusOK, "product_list.html", gin.H{
		"products":   products,
		"categories": categories,
		"suppliers":  suppliers,
		"query":      query,
	})
}