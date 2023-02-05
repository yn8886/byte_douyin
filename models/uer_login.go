package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Id       int64  `json:"id,omitempty"`
	UserId   int64  `json:"user_id,omitempty"`
	UserName string `json:"username,omitempty" gorm:"user_name"` //用户名
	Password string `json:"password,omitempty"`                  //密码
}

func (UserLogin) TableName() string {
	return "user_login"
}

func UserIsExistByUsername(username string) bool {
	var userlogin UserLogin
	err := DB.Where("user_name=?", username).First(&userlogin).Error
	return err == nil
}

func CheckNameAndPwd(username, password string, userLogin *UserLogin) error {
	if userLogin == nil {
		return errors.New("空指针异常：*UserLogin")
	}
	DB.Where("user_name=?", username).First(userLogin)
	// 验证
	if !PwdVerify(password, userLogin.Password) {
		return errors.New("用户名或密码错误")
	}
	return nil
}

func PwdVerify(password, hash string) bool {
	// 比较用户输入的明文和和数据库取出的的密码解析后是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}