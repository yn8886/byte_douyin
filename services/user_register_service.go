package services

import (
	"errors"
	"go_code/project/byte_douyin/middlewares"
	"go_code/project/byte_douyin/models"
	"unicode/utf8"
)

type RegisterResponse struct {
	UserId int32  `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterRequest struct {
	username string
	password string
	response *RegisterResponse
	userid 	 int64
	token    string
}

//注册用户，并得到token和user_id
func PostUserRegister(username, password string) (*RegisterResponse, error) {
	return (&RegisterRequest{username: username, password: password}).Do()
}


func (r *RegisterRequest) Do() (*RegisterResponse, error){
	//验证用户名规范性
	if err := r.CheckUserName(); err != nil {
		return nil, err
	}
	//注册用户到数据库,并生成token
	if err := r.SaveRegisterData(); err != nil {
		return nil, err
	}
	//返回user_id和token
	return r.PackRegisterResponse(), nil	
}

func (r *RegisterRequest) CheckUserName() error {
	if r.username == "" {
		return errors.New("用户名为空")
	}
	if models.UserIsExistByUsername(r.username) {
		return errors.New("用户已存在")
	}
	if utf8.RuneCountInString(r.username) > 32 {
		return errors.New("用户名长度超出限制")
	}
	return nil
}

func (r *RegisterRequest) SaveRegisterData() error {
	userlogin := models.UserLogin{
		UserName: r.username,
		Password: r.password,
	}
	user := models.User{
		Name: r.username,
		UserLogin: &userlogin,
	}
	//注册用户
	err := models.RegisterUser(&user)
	if err != nil {
		return err
	}
	//生成token
	r.token, err = middlewares.GenerateToken(userlogin)
	if err != nil {
		return err
	}
	r.userid = userlogin.UserId
	return nil
}

func (r *RegisterRequest) PackRegisterResponse() *RegisterResponse {
	r.response = &RegisterResponse{
		UserId: int32(r.userid),
		Token: r.token,
	}
	return r.response
}


	
