package utils

import (
	"errors"
	"go_code/project/byte_douyin/models"

	"golang.org/x/crypto/bcrypt"
)

func PwdHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PwdVerify(username, password string, userLogin *models.UserLogin) error {
	userLogin, err := models.QueryUserByUsername(username, userLogin)
	if err != nil {
		return err
	}
	//核对密码,比较用户输入的明文和和数据库取出的的密码解析后是否匹配
	err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password),[]byte(password))
	if err != nil {
		return errors.New("用户名或密码错误")
	}
	return nil
}