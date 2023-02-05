package services

import (
	"errors"
	"go_code/project/byte_douyin/middlewares"
	"go_code/project/byte_douyin/models"
)

type LoginResponse struct {
	UserId int32  `json:"user_id"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	username string
	password string
	response *LoginResponse
	userid 	 int64
	token 	 string
}


func GetUserLoginResponse(username, password string) (*LoginResponse, error) {
	return (&LoginRequest{username: username, password: password}).Do()
}


func (r *LoginRequest) Do() (*LoginResponse, error) {
	//查询用户是否存在
	if ! models.UserIsExistByUsername(r.username) {
		return nil, errors.New("用户名不存在，请先注册")
	}
	//查询用户名和密码是否匹配，若匹配则生成token
	if err := r.QueryLoginData(); err != nil {
		return nil, err
	}
	//返回用户id和权限token
	return r.PackLoginResponse(), nil
}

func (r *LoginRequest) QueryLoginData() error {
	userLogin := models.UserLogin{}
	//查询数据库数值是否一致
	err := models.CheckNameAndPwd(r.username, r.password, &userLogin)
	if err != nil {
		return err
	}
	//生成token
	r.token, err = middlewares.GenerateToken(userLogin)
	if err != nil {
		return err
	}
	r.userid = userLogin.UserId
	return nil
}

func (r *LoginRequest) PackLoginResponse() (*LoginResponse) {
	r.response = &LoginResponse{
		UserId: int32(r.userid),
		Token: r.token,
	}
	return r.response
}