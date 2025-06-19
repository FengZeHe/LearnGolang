package bcrypt

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrGenPasswd = errors.New("ErrGenPasswd")
)

// 密码加密与密码对比功能的实现
func GetPwd(password string) (hashStr string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashStr = string(hash)
	if err != nil {
		return "", ErrGenPasswd
	}
	return hashStr, nil
}

func ComparePwd(PwdBeforeEncryption, EncryptedPasswd string) (result bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(PwdBeforeEncryption), []byte(EncryptedPasswd)); err != nil {
		return false
	} else {
		return true
	}
}
