package model

import (
	"crypto/md5"
	"encoding/hex"
)

// Author:Boyn
// Date:2020/2/24

// 将明文密码转化为MD5
func GeneratePasswordHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	pwdHash := hex.EncodeToString(hasher.Sum(nil))
	return pwdHash
}
