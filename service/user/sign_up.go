package user

import (
	"netdisk/model"
	"netdisk/utils/ygerr"

	"github.com/jinzhu/gorm"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// 用户注册

type SignUpService struct {
	NetDiskNo string `json:"netdisk_no"` // 网盘号
	Password  string `json:"password"`
	Tel       string `json:"tel"`
}

func (srv *SignUpService) Register() ygerr.YgError {
	// 1. 用户是否存在
	var user model.UserModel
	if d := db.First(&user, "id=?", srv.NetDiskNo); d.Error == nil {
		return ygerr.NewWebCtl(ygerr.ClientError, "用户已存在")
	} else if !gorm.IsRecordNotFoundError(d.Error) {
		log.Println(d.Error)
		return ygerr.NewWebCtl(ygerr.ServerError, "未知错误")
	}

	// 2. 创建用户
	user.Id = srv.NetDiskNo
	user.Password = srv.Password
	user.Password = encPassword(srv.Password)
	if d := db.Create(&user); d.Error != nil {
		log.Println(d.Error)
		return ygerr.NewWebCtl(ygerr.ServerError, "创建用户失败")
	}
	return nil
}

// 加密密码
func encPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	return string(hashPassword)
}
