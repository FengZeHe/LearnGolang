package logic

import (
	"BasicProject/dao/mysql"
	"BasicProject/models"
	"BasicProject/pkg/bcrypt"
	"BasicProject/pkg/snowflake"
	"BasicProject/sms"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 处理注册逻辑
func SignIn(user *models.RegisterFormByEmail) (err error) {
	err = mysql.CheckUserExist(user.Email)
	if err == nil {
		// 用户不存在 允许注册
		encipherPassword, _ := bcrypt.GetPwd(user.Password)
		objemail := mysql.StringToSqlNullString(user.Email)
		user := models.User{Email: &objemail, Password: encipherPassword, Id: snowflake.GenId(), Ctime: time.Now().Unix()}
		if err = mysql.CreateUser(&user); err != nil {
			return err
		}
	}
	// 返回注册失败
	return err
}

// 处理登录逻辑
func Login(user *models.LoginForm) (result bool, tempUser models.User, err error) {
	objemail := mysql.StringToSqlNullString(user.Email)
	tempuser := models.User{Email: &objemail, Password: user.Password}
	dbuser, err := mysql.FindByEmail(&tempuser)
	log.Println(dbuser)
	result = bcrypt.ComparePwd(dbuser.Password, tempuser.Password)
	return result, dbuser, err
}

// 处理SMS登录逻辑 生成验证码，并调用发送验证码的方法
func SMSLogin(phone string) (code string, err error) {
	// 1. 生成随机验证码
	source := rand.NewSource(time.Now().UnixNano())
	code = fmt.Sprintf("%06d", rand.New(source).Intn(1000000))
	// 2. 发送验证码
	if err := sms.SendSMS(phone, code); err != nil {
		fmt.Println("SMS Send Code ERROR", err)
		return code, err
	}
	return code, nil
}

// 根据邮箱查询用户信息
func GetUserProfileByEmail(email sql.NullString) (user models.User, err error) {
	//tempUser := models.User{Email: email}
	//user, err = mysql.FindByEmail(&tempUser)
	return user, err
}

// 根据手机号查询用户细腻系
func GetUserProfileByPhone(phone string) (user models.User, err error) {
	tempUser := models.User{Phone: phone}
	user, err = mysql.FindByPhone(&tempUser)
	return user, err
}

// 根据手机号注册用户
func CreateUserByPhone(phone string) (err error) {
	tempUser := models.User{Phone: phone, Id: snowflake.GenId(), Ctime: time.Now().Unix()}
	if err = mysql.CreateUser(&tempUser); err != nil {
		fmt.Println("create user error", err)
		return err
	}
	return err
}

func GetUserProfileById(id string) (user models.User, err error) {
	tempUser := models.User{Id: id}
	user, err = mysql.FindById(&tempUser)
	return user, err
}

// 修改用户信息
func EditUserProfile(userid string, user *models.EditUserProfile) (err error) {
	if err = mysql.UpdateUserProfile(userid, user); err != nil {
		return err
	} else {
		return nil
	}
}
