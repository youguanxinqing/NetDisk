package controller

import (
	"net/http"
	"netdisk/service/user"
	"netdisk/utils/ygerr"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// @Summary 用户注册
// @tags user
// @Produce json
// @Param netdisk_no query string true "网盘号"
// @Param password query string true "密码"
// @Param tel query string true "电话号码"
// @Success 200 {object} SignUpRsp
// @Router /user/sigup [get]
func (*UserController) SignUp(c *gin.Context) {
	srv := new(user.SignUpService)
	if err := c.BindJSON(srv); err != nil {
		c.JSON(http.StatusOK, SignUpRsp{http.StatusUnprocessableEntity, "参数异常", nil})
		return
	}

	if err := srv.Register(); err != nil {
		switch err.Code() {
		case ygerr.ClientError:
			c.JSON(http.StatusOK, SignUpRsp{http.StatusBadRequest, err.Error(), nil})
		case ygerr.ServerError:
			c.JSON(http.StatusOK, SignUpRsp{http.StatusInternalServerError, err.Error(), nil})
		}
		return
	}

	c.JSON(http.StatusOK, SignUpRsp{http.StatusOK, "注册成功", nil})
}

type SignUpRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{}
}

// @Summary 用户登录
// @tags
// @Produce json
// @Param name query string true "姓名"
// @Success 200 {object} your.struct
// @Router url(/xxx/xxx) [get]
func (*UserController) SignIn(c *gin.Context) {

}

// @Summary 获取用户信息
// @tags
// @Produce json
// @Param name query string true "姓名"
// @Success 200 {object} your.struct
// @Router url(/xxx/xxx) [get]
func (*UserController) Info(c *gin.Context) {

}
