package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/middleware"
	"github.com/supermarketmanager/services"
	"net/http"
	"strconv"
)

// SystemLogController 系统日志控制器

type SystemLogController struct {
	systemLogService *services.SystemLogService
}

// NewSystemLogController 创建系统日志控制器实例
func NewSystemLogController() *SystemLogController {
	return &SystemLogController{
		systemLogService: services.NewSystemLogService(),
	}
}

// SetupRoutes 设置路由
func (c *SystemLogController) SetupRoutes(r *gin.Engine) {
	log := r.Group("/log")
	log.Use(middleware.AuthMiddleware())
	{
		log.GET("/list", c.List)
		log.GET("/view/:id", c.View)
	}
}

// List 系统日志列表
func (c *SystemLogController) List(ctx *gin.Context) {
	// 获取所有系统日志
	logs, err := c.systemLogService.GetAllSystemLogs()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "log_list.html", gin.H{
		"logs": logs,
	})
}

// View 查看系统日志详情
func (c *SystemLogController) View(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/log/list")
		return
	}
	
	// 获取系统日志详情
	log, err := c.systemLogService.GetSystemLogByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.HTML(http.StatusOK, "log_view.html", gin.H{
		"log": log,
	})
}