package user

import (
	"netdisk/model"
	"netdisk/utils/ygerr"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SignInService struct {
	NetDiskNo string `json:"netdisk_no"`
	Password  string `json:"password"`
}

func (srv *SignInService) Login() (*Info, ygerr.YgError) {
	// 1. 用户是否存在
	var user model.UserModel
	if d := db.First(&user, "id=?", srv.NetDiskNo); d.RecordNotFound() {
		return nil, ygerr.NewWebCtl(ygerr.ClientError, "用户不存在")
	} else if d.Error != nil {
		log.Error(d.Error)
		return nil, ygerr.NewWebCtl(ygerr.ServerError, "未知错误")
	}
	// 2. 密码是否正确
	if !isEqualPassword(user.Password, srv.Password) {
		return nil, ygerr.NewWebCtl(ygerr.ClientError, "密码错误")
	}

	// 3. 生成用户信息
	return enInfo(&user), nil
}

func isEqualPassword(hashPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

// 影响部分不可公开信息
func enInfo(user *model.UserModel) *Info {
	return &Info{
		NetDiskNo: user.Id,
		NickName:  user.Username,
	}
}
