package model

import (
	"crypto/md5"
	"encoding/hex"
)

// Author:Boyn
// Date:2020/2/24

// 将明文密码转化为MD5
func GeneratePasswordHash(pwd string) string {
	return Md5(pwd)
}

func Md5(origin string) string {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	return hex.EncodeToString(hasher.Sum(nil))
}
