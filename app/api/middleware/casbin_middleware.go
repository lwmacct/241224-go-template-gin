package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// CasbinMiddleware 通过 Role + Path + Method 做权限控制
// 注意：需提前写好 Casbin 的模型 (casbin.conf) & 策略 (policy.csv)。
func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户 Role
		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "role not found in context"})
			c.Abort()
			return
		}
		role := roleVal.(string)

		path := c.Request.URL.Path
		method := c.Request.Method

		// Casbin 的 Enforce 参数顺序要与模型 (conf) 中定义保持一致
		allowed, err := e.Enforce(role, path, method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
