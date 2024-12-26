package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/lwmacct/241224-go-template-gin/app/api/handler"
	"github.com/lwmacct/241224-go-template-gin/app/api/middleware"
	"github.com/lwmacct/241224-go-template-gin/app/api/service"
)

func SetupRouter(userSrv *service.UserService, jwtSecret []byte, enforcer *casbin.Enforcer) *gin.Engine {
	r := gin.Default()

	userHandler := handler.NewUserHandler(userSrv)

	// 登录接口（无需任何中间件）
	r.POST("/login", userHandler.Login)

	// 需要登录才能访问的接口
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware(jwtSecret))

	// 例如 /user/info
	// 再结合 Casbin 中间件来做访问控制
	authGroup.GET("/user/info", middleware.CasbinMiddleware(enforcer), userHandler.GetUserInfo)

	// 也可以对某些路由限制只有 admin / manager 可访问
	// e.g., authGroup.POST("/users", ...)

	return r
}
