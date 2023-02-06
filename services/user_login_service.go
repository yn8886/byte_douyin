package services

import (
	"errors"
	"go_code/project/byte_douyin/models"
	"go_code/project/byte_douyin/utils"
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

//查询登录用户，得到user_id和token
func PostUserLogin(username, password string) (*LoginResponse, error) {
	return (&LoginRequest{username: username, password: password}).Do()
}

func (r *LoginRequest) Do() (*LoginResponse, error) {
	//从数据库中查询用户
	if ! models.UserIsExistByUsername(r.username) {
		return nil, errors.New("用户名不存在，请先注册")
	}
	//查询数据库验证用户名和密码，生成token
	if err := r.QueryLoginData(); err != nil {
		return nil, err
	}
	//返回user_id和token
	return r.PackLoginResponse(), nil
}

func (r *LoginRequest) QueryLoginData() error {
	userLogin := models.UserLogin{}
	//验证用户名和密码
	err := utils.PwdVerify(r.username, r.password, &userLogin)
	if err != nil {
		return err
	}
	//生成token
	r.token, err = utils.GenerateToken(userLogin)
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