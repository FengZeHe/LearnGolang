package logic

import (
	"BasicProject/dao/mysql"
	"BasicProject/models"
	"BasicProject/pkg/bcrypt"
	"BasicProject/pkg/snowflake"
	"time"
)

// 处理注册逻辑
func SignIn(user *models.RegisterForm) (err error) {
	err = mysql.CheckUserExist(user.Email)
	if err == nil {
		// 用户不存在 允许注册
		encipherPassword, _ := bcrypt.GetPwd(user.Password)
		user := models.User{Email: user.Email, Password: encipherPassword, Id: snowflake.GenId(), Ctime: time.Now().Unix()}
		if err = mysql.CreateUser(&user); err != nil {
			return err
		}
	}
	// 返回注册失败
	return err
}

// 处理登录逻辑
func Login(user *models.LoginForm) (result bool, err error) {
	tempuser := models.User{Email: user.Email, Password: user.Password}
	dbuser, err := mysql.FindByEmail(&tempuser)
	result = bcrypt.ComparePwd(dbuser.Password, tempuser.Password)
	return result, err
}

// 获取用户信息
func GetUserProfile(email string) (user models.User, err error) {
	tempUser := models.User{Email: email}
	user, err = mysql.FindByEmail(&tempUser)
	return user, err
}

// 修改用户信息
func EditUserProfile(user *models.User) (err error) {
	if err = mysql.UpdateUserProfile(user); err != nil {
		return err
	} else {
		return nil
	}

}