package huser

import (
	"netdisk/model"
	"netdisk/service/user"
)

// 隐藏用户的关键信息
func HideUserInfo(u *model.UserModel) user.Info {
	return user.Info{
		NetDiskNo: u.Id,
		NickName:  u.Username,
	}
}
