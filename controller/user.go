package controller

import (
	"net/http"
	"netdisk/service/user"
	"netdisk/utils/ygerr"
	"netdisk/utils/ygjwt"

	log "github.com/sirupsen/logrus"

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
// @Router /user/sigup [post]
func (*UserController) SignUp(c *gin.Context) {
	srv := new(user.SignUpService)
	if err := c.ShouldBind(srv); err != nil {
		log.Error(err)
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
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// @Summary 用户登录
// @tags user
// @Produce json
// @Param netdisk_no query string true "网盘号"
// @Param password query string true "密码"
// @Header 200 {string} string "Set-Authorization"
// @Success 200 {object} SignInRsp
// @Router /user/sigin [post]
func (*UserController) SignIn(c *gin.Context) {
	srv := new(user.SignInService)
	if err := c.ShouldBind(srv); err != nil {
		log.Info(err)
		c.JSON(http.StatusOK, SignInRsp{http.StatusUnprocessableEntity, "参数异常", nil})
		return
	}

	if info, err := srv.Login(); err != nil {
		switch err.Code() {
		case ygerr.ClientError:
			c.JSON(http.StatusOK, SignInRsp{http.StatusBadRequest, err.Error(), nil})
		case ygerr.ServerError:
			c.JSON(http.StatusOK, SignInRsp{http.StatusInternalServerError, err.Error(), nil})
		}
		return
	} else {
		// 发放 token
		tokenStr, err := ygjwt.ReleseToken(info.NetDiskNo)
		if err != nil {
			c.JSON(http.StatusOK, SignInRsp{http.StatusInternalServerError, "生成 token 异常", nil})
			return
		}
		c.Header("Set-Authorization", tokenStr)
		c.JSON(http.StatusOK, SignInRsp{http.StatusOK, "登录成功", nil})
		return
	}
}

type SignInRsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// @Summary 获取用户信息
// @tags user
// @Produce json
// @Param Authorization header string true "权限验证"
// @Success 200 {object} InfoRsp
// @Router /user/info [get]
func (*UserController) Info(c *gin.Context) {
	srv := new(user.InfoService)
	srv.SetContext(c)

	info, err := srv.Info()
	if err != nil {
		switch err.Code() {
		case ygerr.ClientError:
			c.JSON(http.StatusOK, InfoRsp{http.StatusBadRequest, err.Error(), nil})
		case ygerr.ServerError:
			c.JSON(http.StatusOK, InfoRsp{http.StatusInternalServerError, err.Error(), nil})
		}
		return
	}
	c.JSON(http.StatusOK, InfoRsp{http.StatusOK, "", info})
	return
}

type InfoRsp struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data *user.Info `json:"data"`
}
