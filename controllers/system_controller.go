package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/supermarketmanager/middleware"
	"github.com/supermarketmanager/services"
	"net/http"
)

// SystemController 系统控制器

type SystemController struct {
	adminService *services.AdminService
}

// NewSystemController 创建系统控制器实例
func NewSystemController() *SystemController {
	return &SystemController{
		adminService: services.NewAdminService(),
	}
}

// SetupRoutes 设置路由
func (c *SystemController) SetupRoutes(r *gin.Engine) {
	r.GET("/login", c.GoLogin)
	r.POST("/login", c.Login)
	r.GET("/logout", c.Logout)

	// 需要身份验证的路由组
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/index", c.GoIndex)
		auth.GET("/password", c.ChangePasswordForm)
		auth.POST("/password", c.ChangePassword)
	}
}

// GoLogin 跳转到登录页面
func (c *SystemController) GoLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

// Login 登录处理
func (c *SystemController) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// 调用服务层进行登录验证
	admin, err := c.adminService.Login(username, password)
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	// 登录成功，将管理员信息存入会话
	session := sessions.Default(ctx)
	session.Set(middleware.AdminSessionKey, admin.ID)
	session.Save()

	// 重定向到首页
	ctx.Redirect(http.StatusFound, "/index")
}

// Logout 登出处理
func (c *SystemController) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete(middleware.AdminSessionKey)
	session.Save()
	ctx.Redirect(http.StatusFound, "/login")
}

// GoIndex 跳转到首页
func (c *SystemController) GoIndex(ctx *gin.Context) {
	session := sessions.Default(ctx)
	adminID := session.Get(middleware.AdminSessionKey).(uint)

	admin, err := c.adminService.GetAdminByID(adminID)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"admin": admin,
	})
}

// ChangePasswordForm 跳转到修改密码页面
func (c *SystemController) ChangePasswordForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "password.html", nil)
}

// ChangePassword 修改密码处理
func (c *SystemController) ChangePassword(ctx *gin.Context) {
	session := sessions.Default(ctx)
	adminID := session.Get(middleware.AdminSessionKey).(uint)

	oldPassword := ctx.PostForm("oldPassword")
	newPassword := ctx.PostForm("newPassword")
	confirmPassword := ctx.PostForm("confirmPassword")

	// 验证新密码和确认密码是否一致
	if newPassword != confirmPassword {
		ctx.HTML(http.StatusOK, "password.html", gin.H{
			"msg": "新密码与确认密码不一致",
		})
		return
	}

	// 调用服务层更新密码
	err := c.adminService.UpdatePassword(adminID, oldPassword, newPassword)
	if err != nil {
		ctx.HTML(http.StatusOK, "password.html", gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 密码修改成功，清除会话，需要重新登录
	session.Delete(middleware.AdminSessionKey)
	session.Save()
	ctx.Redirect(http.StatusFound, "/login")
}
