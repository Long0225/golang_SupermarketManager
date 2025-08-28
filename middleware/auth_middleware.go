package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

// AdminSessionKey 会话中的管理员信息键名
const AdminSessionKey = "admin_in_session"

// AuthMiddleware 身份验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取会话
		session := sessions.Default(c)
		adminID := session.Get(AdminSessionKey)
		
		// 如果没有登录，则重定向到登录页面
		if adminID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		
		// 继续处理请求
		c.Next()
	}
}