package bcrypt

import "golang.org/x/crypto/bcrypt"

// 密码加密
func GetPwd(password string) (hashStr string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashStr = string(hash)
	return hashStr, err
}

// 密码比对
func ComparePwd(PwdBeforeEncryption, EncryptedPassword string) (result bool) {
	// pwd1 for the databases
	if err := bcrypt.CompareHashAndPassword([]byte(PwdBeforeEncryption), []byte(EncryptedPassword)); err != nil {
		return false
	} else {
		return true
	}

}
