package vm

import (
	"log"
	"web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/24

type RegisterModelView struct {
	LoginViewModel
}

type RegisterModelViewOp struct {
}

func (RegisterModelViewOp) GetVM() RegisterModelView {
	v := RegisterModelView{}
	v.SetTitle("注册")
	return v
}

// 检查用户是否存在于数据库中
func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("无法找到用户:", username)
		return false
	}
	return true
}

func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}
