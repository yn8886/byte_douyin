package services

import (
	"errors"
	"go_code/project/byte_douyin/middlewares"
	"go_code/project/byte_douyin/models"
	"regexp"
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

//注册用户并得到token和user_id
func GetUserRegisterResponse(username, password string) (*RegisterResponse, error) {
	return (&RegisterRequest{username: username, password: password}).Do()
}


func (r *RegisterRequest) Do() (*RegisterResponse, error){
	//检验用户名是否规范
	if err := r.CheckUserName(); err != nil {
		return nil, err
	}
	//保存注册的用户名和密码,添加成功则生成token
	if err := r.SaveRegisterData(); err != nil {
		return nil, err
	}
	//返回用户id和权限token
	return r.PackRegisterResponse(), nil	
}

func (r *RegisterRequest) CheckUserName() error {
	if r.username == "" {
		return errors.New("用户名为空")
	}
	if models.UserIsExistByUsername(r.username) {
		return errors.New("用户已存在")
	}
	if regexp.MustCompile(`^.{1,32}$`).FindString(r.username) == "" {
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
	//更新数据库
	err := models.RegisterUser(&user)
	if err != nil {
		return err
	}
	//获取token
	token, err := middlewares.GenerateToken(userlogin)
	if err != nil {
		return err
	}
	r.userid = user.Id
	r.token = token
	return nil
}

func (r *RegisterRequest) PackRegisterResponse() *RegisterResponse {
	r.response = &RegisterResponse{
		UserId: int32(r.userid),
		Token: r.token,
	}
	return r.response
}


	
