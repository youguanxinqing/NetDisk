package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// @Summary
// @tags user
// @Produce json
// @Param name query string true "昵称"
// @Param name query string true "密码"
// @Param name query string true "电话号码"
// @Success 200 {object} your.struct
// @Router /user/sigup [get]
func (*UserController) SignUp(c *gin.Context) {
	name := c.PostForm("name")
	c.JSON(http.StatusOK, gin.H{"name": name})
}
