package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string     `json:"name" gorm:"name"`
	UserLogin     *UserLogin `json:"-" gorm:"foreignkey:UserId"`
	FollowCount   int64      `json:"follow_count" gorm:"follow_count"`
	FollowerCount int64      `json:"follower_count" gorm:"follower_count"`
	IsFollow      bool       `json:"is_follow" gorm:"is_follow"`
}

func (User) TableName() string {
	return "user_info"
}

func RegisterUser(user *User) error {
	err := DB.Create(user).Error
	return err
}

func QueryUserInfoByUserId(userId int64, user *User) error {
	if user == nil {
		return errors.New("空指针异常：*User")
	}
	err := DB.Where("id=?", userId).Select([]string{"id","name","follow_count","follower_count", "is_follow"}).First(&user).Error
	if err != nil {
		return errors.New("用户不存在")
	}
	return nil
}