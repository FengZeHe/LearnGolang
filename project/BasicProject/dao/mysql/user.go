package mysql

import (
	"BasicProject/models"
	"fmt"
	"github.com/pkg/errors"
)

// 查询是否已经存在该用户
func CheckUserExist(email string) (err error) {
	var user models.User
	result := db.Find(&user, models.User{Email: email})
	if result.RowsAffected > 0 {
		// 存在 不能注册
		err = errors.New(ErrorUserExit)
		return err
	} else if result.Error == nil {
		// 不存在 可以注册
		return nil
	} else {
		// 查询出错
		err = errors.New(ErrorQueryFailed)
		return err
	}
}

// 创建新用户
func CreateUser(user *models.User) (err error) {
	if err = db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 根据邮箱查询用户信息
func FindByEmail(user *models.User) (result models.User, err error) {
	if err = db.Where("email = ?", user.Email).Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

// 根据用户id查询信息
func FindById(user *models.User) (result models.User, err error) {
	if err = db.Where("id=?", user.Id).Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

// 更新用户信息
func UpdateUserProfile(userid string, user *models.EditUserProfile) (err error) {
	if err = db.Table("users").Where("id=?", userid).Save(&user).Error; err != nil {
		return err
	}
	return nil
}
