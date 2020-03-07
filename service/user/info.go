package user

import (
	"netdisk/utils/ygerr"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type InfoService struct {
	ctx *gin.Context
}

type Info struct {
	NetDiskNo string `json:"netdisk_no"`
	NickName  string `json:"nickname"`
}

func (srv *InfoService) Info() (*Info, ygerr.YgError) {
	user, exists := srv.ctx.Get("user")
	if !exists {
		log.Error("未登录却被访问 /user/info")
		return nil, ygerr.NewWebCtl(ygerr.ServerError, ygerr.SerErrStr)
	}

	info, ok := user.(Info)
	if !ok {
		log.Error("类型转换异常 interface{} -> service.Info")
		return nil, ygerr.NewWebCtl(ygerr.ServerError, ygerr.SerErrStr)
	}

	return &info, nil
}

func (srv *InfoService) SetContext(c *gin.Context) {
	srv.ctx = c
}
