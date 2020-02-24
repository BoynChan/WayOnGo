package vm

import (
	"log"
	"web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/23

type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

type LoginViewModelOp struct {
}

// 检查用户是否已经登录
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("找不到用户:", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password)
}

func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}
