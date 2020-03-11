package router

import (
	"net/http"
	"netdisk/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func registerSwagger() {

	docs.SwaggerInfo.Title = "NetDisk Api"
	docs.SwaggerInfo.Description = "一个简单的网盘系统"
	docs.SwaggerInfo.Version = "1.0"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		return
	})
}
