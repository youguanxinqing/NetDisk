package router

import (
	"netdisk/controller"
	"netdisk/middleware"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func New() *gin.Engine {
	// user
	gUser := r.Group("/user")
	registerUser(gUser)

	// ...

	// swagger
	registerSwagger()
	return r
}

func registerUser(g *gin.RouterGroup) {
	ctl := new(controller.UserController)
	g.POST("/sigup", ctl.SignUp) // 创建用户
	g.POST("/sigin", ctl.SignIn) // 用户登录

	gauth := g.Use(middleware.AuthMiddleWare)
	gauth.GET("/info", ctl.Info) // 获取用户信息
}

func init() {
	r = gin.Default()
}
