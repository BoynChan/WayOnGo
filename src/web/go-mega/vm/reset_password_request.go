package vm

import (
	"log"
	"web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/25

type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

type ResetPasswordRequestViewModelOp struct {
}

func (ResetPasswordRequestViewModelOp) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("重置密码")
	return v
}

// 检查该邮箱地址是否注册了用户
func CheckEmailExist(email string) bool {
	_, err := model.GetUserByEmail(email)
	if err != nil {
		log.Println("[CheckEmailExist] 无法找到邮件地址对应用户", err)
		return false
	}
	return true
}
