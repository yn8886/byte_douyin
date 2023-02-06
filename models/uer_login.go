package models

import (
	"errors"
)

type UserLogin struct {
	Id 		 int64  `gorm:"id"`
	UserId   int64  `gorm:"user_id"`
	UserName string `json:"username" gorm:"user_name"` //用户名
	Password string `json:"password" gorm:"password"`  //密码
}

func (UserLogin) TableName() string {
	return "user_login"
}

func UserIsExistByUsername(username string) bool {
	var userlogin UserLogin
	err := DB.Where("user_name=?", username).First(&userlogin).Error
	return err == nil
}

func QueryUserByUsername(username string, userLogin *UserLogin) (*UserLogin, error) {
	if userLogin == nil {
		return nil, errors.New("空指针异常：*UserLogin")
	}
	DB.Where("user_name=?", username).First(userLogin)
	return userLogin, nil
}
