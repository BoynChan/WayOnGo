package vm

import "web/go-mega/model"

// Author:Boyn
// Date:2020/2/25

type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

type ResetPasswordViewModelOp struct {
}

func (ResetPasswordViewModelOp) GetVM(token string) ResetPasswordViewModel {
	v := ResetPasswordViewModel{}
	v.SetTitle("重置密码")
	v.Token = token
	return v
}

// 检查Token是否合法
func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

// 重置密码
func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}
