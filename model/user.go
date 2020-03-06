package model

type UserModel struct {
	ygModel
	Id       string `gorm:"column:id;primary_key"`
	Username string `gorm:"column:username;min=4"`
	Password string `gorm:"column:password"`
	Tel      string `gorm:"column:tel"`
}

func (UserModel) TableName() string {
	return "user"
}
