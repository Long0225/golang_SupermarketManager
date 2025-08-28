package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/middleware"
	"github.com/supermarketmanager/models"
	"github.com/supermarketmanager/services"
	"net/http"
	"strconv"
)

// CategoryController 分类控制器

type CategoryController struct {
	categoryService *services.CategoryService
}

// NewCategoryController 创建分类控制器实例
func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: services.NewCategoryService(),
	}
}

// SetupRoutes 设置路由
func (c *CategoryController) SetupRoutes(r *gin.Engine) {
	category := r.Group("/category")
	category.Use(middleware.AuthMiddleware())
	{
		category.GET("/list", c.List)
		category.GET("/view/:id", c.View)
		category.GET("/toAdd", c.ToAdd)
		category.POST("/add", c.Add)
		category.GET("/toUpdate/:id", c.ToUpdate)
		category.POST("/update", c.Update)
		category.GET("/del/:id", c.Delete)
	}
}

// List 分类列表
func (c *CategoryController) List(ctx *gin.Context) {
	// 获取所有分类
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "category_list.html", gin.H{
		"categories": categories,
	})
}

// View 查看分类详情
func (c *CategoryController) View(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/category/list")
		return
	}
	
	// 获取分类详情
	category, err := c.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "category_view.html", gin.H{
		"category": category,
	})
}

// ToAdd 跳转到添加分类页面
func (c *CategoryController) ToAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "category_add.html", nil)
}

// Add 添加分类
func (c *CategoryController) Add(ctx *gin.Context) {
	// 绑定表单数据
	var category models.Category
	if err := ctx.ShouldBind(&category); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 创建分类
	if err := c.categoryService.CreateCategory(&category); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到分类列表
	ctx.Redirect(http.StatusFound, "/category/list")
}

// ToUpdate 跳转到修改分类页面
func (c *CategoryController) ToUpdate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/category/list")
		return
	}
	
	// 获取分类详情
	category, err := c.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "category_update.html", gin.H{
		"category": category,
	})
}

// Update 更新分类
func (c *CategoryController) Update(ctx *gin.Context) {
	// 绑定表单数据
	var category models.Category
	if err := ctx.ShouldBind(&category); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 更新分类
	if err := c.categoryService.UpdateCategory(&category); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到分类列表
	ctx.Redirect(http.StatusFound, "/category/list")
}

// Delete 删除分类
func (c *CategoryController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/category/list")
		return
	}
	
	// 删除分类
	if err := c.categoryService.DeleteCategory(uint(id)); err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 重定向到分类列表
	ctx.Redirect(http.StatusFound, "/category/list")
}